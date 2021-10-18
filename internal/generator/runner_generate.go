package generator

/*
 * Runner's methods to to generate typescript files
 */

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/iancoleman/strcase"
	"github.com/join-com/protoc-gen-ts/internal/join_proto"
	"github.com/join-com/protoc-gen-ts/internal/utils"
	"github.com/join-com/protoc-gen-ts/version"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/descriptorpb"
)

func (r *Runner) generateTypescriptFile(protoFile *protogen.File, generatedFileStream *protogen.GeneratedFile) {
	// TODO: Generate comment with version, in order to improve traceability & debugging experience
	generatedFileStream.P(
		"// GENERATED CODE -- DO NOT EDIT!\n",
		"// GENERATOR VERSION: "+version.MajorVersion+"."+version.MinorVersion+"."+version.PatchVersion+"."+version.BuildCommit+"."+version.BuildTime+"\n",
		"/* eslint-disable @typescript-eslint/no-non-null-assertion */\n",
	)

	r.generateTypescriptImports(protoFile, generatedFileStream)
	r.generateTypescriptNamespace(generatedFileStream, protoFile)
}

func (r *Runner) generateTypescriptImports(protoFile *protogen.File, generatedFileStream *protogen.GeneratedFile) {
	currentSourcePath := protoFile.Desc.Path()
	importSourcePaths := protoFile.Proto.GetDependency()

	// Generic imports
	if len(protoFile.Proto.GetService()) > 0 {
		generatedFileStream.P("import * as joinGRPC from '@join-com/grpc'")
	}
	generatedFileStream.P("import * as protobufjs from 'protobufjs/light'\n")

	// Custom imports
	importLines := make([]string, 0, len(importSourcePaths))
	for _, importSourcePath := range importSourcePaths {
		if !r.importCodeOptions[importSourcePath] {
			continue
		}

		importLines = append(importLines, fmt.Sprintf(
			"import { %s } from '%s'",
			r.generateImportName(currentSourcePath, importSourcePath),
			fromProtoPathToGeneratedPath(importSourcePath, currentSourcePath),
		))
	}

	if r.fileHasFlavors(generatedFileStream, protoFile) {
		importLines = append(importLines, "import { WithFlavor } from '@coderspirit/nominal'")
	}

	sort.Slice(importLines, func(i, j int) bool {
		return importLines[i] < importLines[j]
	})
	for _, importLine := range importLines {
		generatedFileStream.P(importLine)
	}
	generatedFileStream.P("")

	if len(protoFile.Proto.GetService()) > 0 {
		generatedFileStream.P("import { grpc } from '@join-com/grpc'\n")
	}
}

func (r *Runner) fileHasFlavors(generatedFileStream *protogen.GeneratedFile, protoFile *protogen.File) bool {
	for _, messageSpec := range protoFile.Proto.GetMessageType() {
		for _, fieldSpec := range messageSpec.GetField() {
			fieldOptions := fieldSpec.GetOptions()
			if fieldOptions == nil {
				continue
			}

			flavorName, found := join_proto.GetStringCustomFieldOption("typescript_flavor", fieldOptions, r.extensionTypes)
			if found && flavorName != "" {
				return true
			}
		}
	}

	return false
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
	r.currentPackage = protoFile.Proto.GetPackage()
	r.currentNamespace = getNamespaceFromProtoPackage(r.currentPackage)
	r.P(
		generatedFileStream,
		"// eslint-disable-next-line @typescript-eslint/no-namespace",
		"export namespace "+r.currentNamespace+" {\n",
	)
	r.indentLevel += 2

	r.generateTypescriptClassDecoratorDefinition(generatedFileStream, protoFile)

	// This interface is namespace-private, as it's being replicated for every generated file
	r.P(
		generatedFileStream,
		"interface ConvertibleTo<T> {",
		"  asInterface(): T",
		"}\n",
	)

	r.generateTypescriptFlavors(generatedFileStream, protoFile)
	r.generateTypescriptEnums(generatedFileStream, protoFile)
	r.generateTypescriptMessageInterfaces(generatedFileStream, protoFile)
	r.generateTypescriptMessageClasses(generatedFileStream, protoFile)
	r.generateTypescriptServiceDefinitions(generatedFileStream, protoFile)
	r.generateTypescriptClient(generatedFileStream, protoFile)

	r.currentNamespace = ""
	r.currentPackage = ""
	r.indentLevel -= 2
	r.P(generatedFileStream, "\n}")
}

func (r *Runner) generateTypescriptClassDecoratorDefinition(generatedFileStream *protogen.GeneratedFile, protoFile *protogen.File) {
	r.P(
		generatedFileStream,
		"const registerGrpcClass = <T extends protobufjs.Message<T>>(",
		"  typeName: string",
		"): protobufjs.TypeDecorator<T> => {",
		"  if (protobufjs.util.decorateRoot.get(typeName) != null) {",
		"    // eslint-disable-next-line @typescript-eslint/ban-types",
		"    return (",
		"      // eslint-disable-next-line @typescript-eslint/no-unused-vars",
		"      _: protobufjs.Constructor<T>",
		"    ): void => {",
		"      // Do nothing",
		"    }",
		"  }",
		"  return protobufjs.Type.d(typeName)",
		"}",
	)
}

func (r *Runner) generateTypescriptFlavors(generatedFileStream *protogen.GeneratedFile, protoFile *protogen.File) {
	flavorBaseTypesMap := make(map[string]string)
	flavorDeclarations := make([]string, 0)

	for _, messageSpec := range protoFile.Proto.GetMessageType() {
		for _, fieldSpec := range messageSpec.GetField() {
			fieldOptions := fieldSpec.GetOptions()
			if fieldOptions == nil {
				continue
			}

			flavorName, found := join_proto.GetStringCustomFieldOption("typescript_flavor", fieldOptions, r.extensionTypes)
			if !found || flavorName == "" {
				continue
			}

			var baseType string
			switch fieldSpec.GetType() {
			case descriptorpb.FieldDescriptorProto_TYPE_BOOL:
				baseType = "boolean"
			case descriptorpb.FieldDescriptorProto_TYPE_INT32:
				baseType = "number"
			case descriptorpb.FieldDescriptorProto_TYPE_UINT32:
				baseType = "number"
			case descriptorpb.FieldDescriptorProto_TYPE_SINT32:
				baseType = "number"
			case descriptorpb.FieldDescriptorProto_TYPE_INT64:
				baseType = "number"
			case descriptorpb.FieldDescriptorProto_TYPE_UINT64:
				baseType = "number"
			case descriptorpb.FieldDescriptorProto_TYPE_SINT64:
				baseType = "number"
			case descriptorpb.FieldDescriptorProto_TYPE_FLOAT:
				baseType = "number"
			case descriptorpb.FieldDescriptorProto_TYPE_DOUBLE:
				baseType = "number"
			case descriptorpb.FieldDescriptorProto_TYPE_FIXED32:
				baseType = "number"
			case descriptorpb.FieldDescriptorProto_TYPE_FIXED64:
				baseType = "number"
			case descriptorpb.FieldDescriptorProto_TYPE_SFIXED32:
				baseType = "number"
			case descriptorpb.FieldDescriptorProto_TYPE_SFIXED64:
				baseType = "number"
			case descriptorpb.FieldDescriptorProto_TYPE_STRING:
				baseType = "string"
			case descriptorpb.FieldDescriptorProto_TYPE_BYTES:
				baseType = "Uint8Array"
			default:
				utils.LogError("Only primitive types can be flavored (" + fieldSpec.GetJsonName() + ")")
			}

			previousBaseType, ok := flavorBaseTypesMap[flavorName]
			if ok {
				if previousBaseType != baseType {
					utils.LogError("Declared flavor " + flavorName + " multiple times with different base types (" + baseType + ", " + previousBaseType + ")")
				}
			} else {
				flavorBaseTypesMap[flavorName] = baseType
				flavorDeclarations = append(flavorDeclarations, "export type "+flavorName+" = WithFlavor<"+baseType+", '"+flavorName+"'>")
			}
		}
	}

	sort.Slice(flavorDeclarations, func(i, j int) bool {
		return flavorDeclarations[i] < flavorDeclarations[j]
	})
	for _, flavorDeclaration := range flavorDeclarations {
		generatedFileStream.P(flavorDeclaration)
	}
	generatedFileStream.P("")
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
	if fieldOptions != nil && fieldOptions.GetDeprecated() {
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
