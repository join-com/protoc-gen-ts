package generator

/*
 * Runner's methods to to generate typescript files
 */

import (
	"fmt"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/join-com/protoc-gen-ts/internal/utils"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/descriptorpb"
)

func (r *Runner) generateTypescriptFile(protoFile *protogen.File, generatedFileStream *protogen.GeneratedFile) {
	// TODO: Generate comment with version, in order to improve traceability & debugging experience
	generatedFileStream.P("// GENERATED CODE -- DO NOT EDIT!\n")

	r.generateTypescriptImports(protoFile.Proto.Dependency, generatedFileStream)
	r.generateTypescriptNamespace(protoFile, generatedFileStream)
}

func (r *Runner) generateTypescriptImports(dependencies []string, generatedFileStream *protogen.GeneratedFile) {
	// Generic imports
	generatedFileStream.P("import * as joinGRPC from '@join-com/grpc'")
	generatedFileStream.P("import * as nodeTrace from '@join-com/node-trace'")
	generatedFileStream.P("")

	// Custom imports
	for _, dependency := range dependencies {
		packageName, validPkgName := r.packagesByFile[dependency]
		if !validPkgName {
			utils.LogError("Unable to retrieve package name for " + dependency)
		}
		generatedFileStream.P(fmt.Sprintf(
			"import { %s } from './%s'",
			strcase.ToCamel(packageName),
			fromProtoPathToGeneratedPath(dependency),
		))
	}
	generatedFileStream.P("")
}

func (r *Runner) generateTypescriptNamespace(protoFile *protogen.File, generatedFileStream *protogen.GeneratedFile) {
	namespaceName := strcase.ToCamel(protoFile.Proto.GetPackage())
	r.P(generatedFileStream, fmt.Sprintf("export namespace %s {\n", namespaceName))
	r.indentLevel += 2

	r.generateTypescriptEnums(protoFile, generatedFileStream)
	r.generateTypescriptMessageInterfaces(protoFile, generatedFileStream)

	r.P(generatedFileStream, "\n}")
	r.indentLevel -= 2
}

func (r *Runner) generateTypescriptEnums(protoFile *protogen.File, generatedFileStream *protogen.GeneratedFile) {
	for _, enumDescriptor := range protoFile.Proto.GetEnumType() {
		var values []string
		for _, enumValue := range enumDescriptor.GetValue() {
			values = append(values, fmt.Sprintf("'%s'", enumValue.GetName()))
		}

		enumName := strcase.ToCamel(enumDescriptor.GetName())
		enumValues := strings.Join(values, " | ")

		r.P(generatedFileStream, fmt.Sprintf("export type %s = %s", enumName, enumValues))
	}
	r.P(generatedFileStream, "")
}

func (r *Runner) generateTypescriptMessageInterfaces(protoFile *protogen.File, generatedFileStream *protogen.GeneratedFile) {
	for _, messageSpec := range protoFile.Proto.GetMessageType() {
		interfaceName := fmt.Sprintf("export interface I%s {", strcase.ToCamel(messageSpec.GetName()))

		r.P(generatedFileStream, interfaceName)
		r.indentLevel += 2

		for _, field := range messageSpec.GetField() {
			// TODO: Add support for required fields (via proto options)
			r.P(generatedFileStream, fmt.Sprintf("%s?: %s", field.GetJsonName(), getMessageFieldType(field)))
		}

		r.indentLevel -= 2
		r.P(generatedFileStream, "}\n")
	}
}

func getMessageFieldType(messageField *descriptorpb.FieldDescriptorProto) string {
	baseType := "unknown"

	switch messageField.GetType() {
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
	case descriptorpb.FieldDescriptorProto_TYPE_ENUM:
		baseType = getEnumOrMessageTypeName(messageField.GetTypeName(), false)
	case descriptorpb.FieldDescriptorProto_TYPE_MESSAGE:
		baseType = getEnumOrMessageTypeName(messageField.GetTypeName(), true)
	}

	if messageField.GetLabel() == descriptorpb.FieldDescriptorProto_LABEL_REPEATED {
		baseType = fmt.Sprintf("%s[]", baseType)
	}

	return baseType
}

func getEnumOrMessageTypeName(typeName string, isInterface bool) string {
	names := strings.Split(typeName, ".")
	interfaceName := names[len(names)-1]
	if isInterface {
		interfaceName = "I" + interfaceName
	}
	return interfaceName
}
