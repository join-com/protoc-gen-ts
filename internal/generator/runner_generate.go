package generator

/*
 * Runner's methods to to generate typescript files
 */

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/join-com/protoc-gen-ts/internal/join_proto"
	"github.com/join-com/protoc-gen-ts/internal/utils"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/descriptorpb"
)

func (r *Runner) generateTypescriptFile(protoFile *protogen.File, generatedFileStream *protogen.GeneratedFile) {
	// TODO: Generate comment with version, in order to improve traceability & debugging experience
	generatedFileStream.P("// GENERATED CODE -- DO NOT EDIT!\n")

	r.generateTypescriptImports(protoFile.Desc.Path(), protoFile.Proto.GetDependency(), generatedFileStream)
	r.generateTypescriptNamespace(generatedFileStream, protoFile)
}

func (r *Runner) generateTypescriptImports(currentSourcePath string, importSourcePaths []string, generatedFileStream *protogen.GeneratedFile) {
	// Generic imports
	generatedFileStream.P("// import * as joinGRPC from '@join-com/grpc'")        // TODO: Remove comment when import is used
	generatedFileStream.P("// import * as nodeTrace from '@join-com/node-trace'") // TODO: Remove comment when import is used
	generatedFileStream.P("import * as protobufjs from 'protobufjs/light'")
	generatedFileStream.P("")

	// Custom imports
	for _, importSourcePath := range importSourcePaths {
		if !r.importCodeOptions[importSourcePath] {
			continue
		}

		generatedFileStream.P(fmt.Sprintf(
			"import { %s } from './%s'",
			r.generateImportName(currentSourcePath, importSourcePath),
			fromProtoPathToGeneratedPath(importSourcePath),
		))
	}
	generatedFileStream.P("")
}

func (r *Runner) generateImportName(currentSourcePath string, importSourcePath string) string {
	packageName, validPkgName := r.packagesByFile[importSourcePath]
	if !validPkgName {
		utils.LogError("Unable to retrieve package name for " + importSourcePath)
	}

	pkgNamespace := strcase.ToCamel(packageName)

	alternativeImportNames, validImportsMap := r.alternativeImportNames[currentSourcePath]
	if !validImportsMap || alternativeImportNames == nil {
		utils.LogError("Unable to retrieve alternative imports map for " + currentSourcePath)
	}

	alternativeImportName, validAlternativeName := alternativeImportNames[importSourcePath]
	if !validAlternativeName {
		utils.LogError("Unable to retrieve alternative import name for " + importSourcePath + " on " + currentSourcePath)
	}

	if pkgNamespace == alternativeImportName {
		return pkgNamespace
	} else {
		return pkgNamespace + " as " + alternativeImportName
	}
}

func (r *Runner) generateTypescriptNamespace(generatedFileStream *protogen.GeneratedFile, protoFile *protogen.File) {
	namespace := getNamespaceFromProtoPackage(protoFile.Proto.GetPackage())
	r.P(generatedFileStream, fmt.Sprintf("export namespace %s {\n", namespace))
	r.indentLevel += 2
	r.currentNamespace = namespace

	r.generateTypescriptEnums(generatedFileStream, protoFile)
	r.generateTypescriptMessageInterfaces(generatedFileStream, protoFile)
	r.generateTypescriptMessageClasses(protoFile, generatedFileStream)

	r.currentNamespace = ""
	r.indentLevel -= 2
	r.P(generatedFileStream, "\n}")
}

func (r *Runner) generateTypescriptEnums(generatedFileStream *protogen.GeneratedFile, protoFile *protogen.File) {
	for _, enumSpec := range protoFile.Proto.GetEnumType() {
		options := enumSpec.GetOptions()
		if options != nil {
			if options.Deprecated != nil && *options.Deprecated {
				r.P(generatedFileStream, "/**\n  * @deprecated\n */")
			}
		}

		var values []string
		enumValueSpecs := enumSpec.GetValue()
		for _, enumValue := range enumValueSpecs {
			values = append(values, "'"+enumValue.GetName()+"'")
		}

		enumName := strcase.ToCamel(enumSpec.GetName())
		enumValues := strings.Join(values, " | ")

		r.P(generatedFileStream, "export type "+enumName+" = "+enumValues)
		r.P(generatedFileStream, "export enum "+enumName+"_Enum {")
		r.indentLevel += 2
		for valueIndex, enumValue := range enumValueSpecs {
			valueLine := enumValue.GetName() + " = " + strconv.FormatInt(int64(enumValue.GetNumber()), 10)
			if valueIndex < len(enumValueSpecs)-1 {
				valueLine += ","
			}
			r.P(generatedFileStream, valueLine)
		}
		r.indentLevel -= 2
		r.P(generatedFileStream, "}\n")
	}
	r.P(generatedFileStream, "")
}

func (r *Runner) generateTypescriptMessageInterfaces(generatedFileStream *protogen.GeneratedFile, protoFile *protogen.File) {
	for _, messageSpec := range protoFile.Proto.GetMessageType() {
		requiredFields := false
		messageOptions := messageSpec.GetOptions()
		if messageOptions != nil {
			if messageOptions.GetDeprecated() {
				r.P(generatedFileStream, "/**\n  * @deprecated\n */")
			}

			_requiredFields, found := join_proto.GetBooleanCustomMessageOption("typescript_required_fields", messageOptions, r.extensionTypes)
			if found {
				requiredFields = _requiredFields
			}
		}

		interfaceOpening := "export interface I" + strcase.ToCamel(messageSpec.GetName()) + " {"
		r.P(generatedFileStream, interfaceOpening)
		r.indentLevel += 2

		for _, fieldSpec := range messageSpec.GetField() {
			fieldOptions := fieldSpec.GetOptions()
			if fieldOptions != nil {
				if messageOptions.GetDeprecated() {
					r.P(generatedFileStream, "/**\n  * @deprecated\n */")
				}
			}

			separator := "?: "
			requiredField, foundRequired := join_proto.GetBooleanCustomFieldOption("typescript_required", fieldSpec.GetOptions(), r.extensionTypes)
			optionalField, foundOptional := join_proto.GetBooleanCustomFieldOption("typescript_optional", fieldSpec.GetOptions(), r.extensionTypes)
			if foundRequired && requiredField && foundOptional && optionalField {
				utils.LogError("incompatible options for field " + fieldSpec.GetName() + " in " + messageSpec.GetName())
			}
			if requiredFields && !(foundOptional && optionalField) || foundRequired && requiredField {
				separator = ": "
			}

			r.P(generatedFileStream, fieldSpec.GetJsonName()+separator+r.getInterfaceFieldType(fieldSpec))
		}

		r.indentLevel -= 2
		r.P(generatedFileStream, "}\n")
	}
}

func (r *Runner) generateTypescriptMessageClasses(protoFile *protogen.File, generatedFileStream *protogen.GeneratedFile) {
	for _, messageSpec := range protoFile.Proto.GetMessageType() {
		r.generateTypescriptMessageClass(messageSpec, generatedFileStream)
	}
}

func (r *Runner) generateTypescriptMessageClass(messageSpec *descriptorpb.DescriptorProto, generatedFileStream *protogen.GeneratedFile) {
	requiredFields := false
	messageOptions := messageSpec.GetOptions()
	if messageOptions != nil {
		if messageOptions.GetDeprecated() {
			r.P(generatedFileStream, "/**\n  * @deprecated\n */")
		}

		_requiredFields, found := join_proto.GetBooleanCustomMessageOption("typescript_required_fields", messageOptions, r.extensionTypes)
		if found {
			requiredFields = _requiredFields
		}
	}

	className := strcase.ToCamel(messageSpec.GetName())
	r.P(
		generatedFileStream,
		"@protobufjs.Type.d('"+className+"')",
		"export class "+className+" extends protobufjs.Message<"+className+"> {\n",
	)
	r.indentLevel += 2

	for _, fieldSpec := range messageSpec.GetField() {
		r.generateTypescriptClassField(generatedFileStream, fieldSpec, messageSpec, messageOptions, requiredFields)
	}

	r.generateTypescriptClassPatchedMethods()

	r.indentLevel -= 2
	r.P(generatedFileStream, "}\n")
}

func (r *Runner) generateTypescriptClassField(
	generatedFileStream *protogen.GeneratedFile,
	fieldSpec *descriptorpb.FieldDescriptorProto,
	messageSpec *descriptorpb.DescriptorProto,
	messageOptions *descriptorpb.MessageOptions,
	requiredFields bool,
) {
	fieldOptions := fieldSpec.GetOptions()
	if fieldOptions != nil {
		if messageOptions.GetDeprecated() {
			r.P(generatedFileStream, "/**\n  * @deprecated\n */")
		}
	}

	separator := "?: "
	requiredField, foundRequired := join_proto.GetBooleanCustomFieldOption("typescript_required", fieldSpec.GetOptions(), r.extensionTypes)
	optionalField, foundOptional := join_proto.GetBooleanCustomFieldOption("typescript_optional", fieldSpec.GetOptions(), r.extensionTypes)
	if foundRequired && requiredField && foundOptional && optionalField {
		utils.LogError("incompatible options for field " + fieldSpec.GetName() + " in " + messageSpec.GetName())
	}
	if requiredFields && !(foundOptional && optionalField) || foundRequired && requiredField {
		separator = "!: "
	}

	r.P(
		generatedFileStream,
		r.getMessageFieldDecorator(fieldSpec),
		"public "+fieldSpec.GetJsonName()+separator+r.getClassFieldType(fieldSpec)+"\n",
	)
}

func (r *Runner) generateTypescriptClassPatchedMethods() {
	// Decode method

	// Encode method

	// Create method
}
