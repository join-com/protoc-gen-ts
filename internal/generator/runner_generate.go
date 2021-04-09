package generator

/*
 * Runner's methods to to generate typescript files
 */

import (
	"fmt"
	"strconv"

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
	r.P(generatedFileStream, "export namespace "+namespace+" {\n")
	r.indentLevel += 2
	r.currentNamespace = namespace

	// This interface is namespace-private, as it's being replicated for every generated file
	r.P(
		generatedFileStream,
		"interface ConvertibleTo<T> {",
		"  asInterface(): T",
		"}\n",
	)

	r.generateTypescriptEnums(generatedFileStream, protoFile)
	r.generateTypescriptMessageInterfaces(generatedFileStream, protoFile)
	r.generateTypescriptMessageClasses(generatedFileStream, protoFile)

	r.currentNamespace = ""
	r.indentLevel -= 2
	r.P(generatedFileStream, "\n}")
}

func (r *Runner) generateTypescriptEnums(generatedFileStream *protogen.GeneratedFile, protoFile *protogen.File) {
	for _, enumSpec := range protoFile.Proto.GetEnumType() {
		options := enumSpec.GetOptions()
		isDeprecated := false
		if options != nil {
			if options.Deprecated != nil && *options.Deprecated {
				isDeprecated = true
			}
		}

		enumName := strcase.ToCamel(enumSpec.GetName())

		if isDeprecated {
			r.P(generatedFileStream, "/**\n  * @deprecated\n */")
		}
		r.P(generatedFileStream, "export enum "+enumName+"_Enum {")
		r.indentLevel += 2
		enumValueSpecs := enumSpec.GetValue()
		for valueIndex, enumValue := range enumValueSpecs {
			valueLine := enumValue.GetName() + " = " + strconv.FormatInt(int64(enumValue.GetNumber()), 10)
			if valueIndex < len(enumValueSpecs)-1 {
				valueLine += ","
			}
			r.P(generatedFileStream, valueLine)
		}
		r.indentLevel -= 2
		r.P(generatedFileStream, "}")

		if isDeprecated {
			r.P(generatedFileStream, "/**\n  * @deprecated\n */")
		}
		r.P(generatedFileStream, "export type "+enumName+" = keyof typeof "+enumName+"_Enum\n")
	}
	r.P(generatedFileStream, "")
}

func (r *Runner) generateTypescriptMessageInterfaces(generatedFileStream *protogen.GeneratedFile, protoFile *protogen.File) {
	for _, messageSpec := range protoFile.Proto.GetMessageType() {
		r.generateTypescriptMessageInterface(generatedFileStream, messageSpec)
	}
}

func (r *Runner) generateTypescriptMessageInterface(generatedFileStream *protogen.GeneratedFile, messageSpec *descriptorpb.DescriptorProto) {
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
		r.generateTypescriptInterfaceField(
			generatedFileStream,
			fieldSpec,
			messageSpec,
			requiredFields,
		)
	}

	r.indentLevel -= 2
	r.P(generatedFileStream, "}\n")
}

func (r *Runner) generateTypescriptInterfaceField(
	generatedFileStream *protogen.GeneratedFile,
	fieldSpec *descriptorpb.FieldDescriptorProto,
	messageSpec *descriptorpb.DescriptorProto,
	requiredFields bool,
) {
	fieldOptions := fieldSpec.GetOptions()
	messageOptions := messageSpec.GetOptions()
	if fieldOptions != nil && messageOptions != nil && messageOptions.GetDeprecated() {
		r.P(generatedFileStream, "/**\n  * @deprecated\n */")
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

func messageHasEnums(messageSpec *descriptorpb.DescriptorProto) bool {
	for _, fieldSpec := range messageSpec.GetField() {
		switch t := fieldSpec.GetType(); t {
		case descriptorpb.FieldDescriptorProto_TYPE_ENUM:
			return true
		}
	}

	for _, nestedMessageSpec := range messageSpec.GetNestedType() {
		hasEnums := messageHasEnums(nestedMessageSpec)
		if hasEnums {
			return true
		}
	}

	return false
}
