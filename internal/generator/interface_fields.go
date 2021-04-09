package generator

import (
	"strconv"
	"strings"

	"github.com/join-com/protoc-gen-ts/internal/utils"
	"google.golang.org/protobuf/types/descriptorpb"
)

func (r *Runner) getEnumOrMessageTypeName(typeName string, isInterface bool) string {
	typeParts := strings.Split(typeName, ".")
	lastIndex := len(typeParts) - 1

	interfaceName := typeParts[lastIndex]
	if isInterface {
		interfaceName = "I" + interfaceName
	}

	protoPackageName := strings.Join(typeParts[0:lastIndex], ".")
	importName := getNamespaceFromProtoPackage(protoPackageName)

	if importName == "" || importName == r.currentNamespace {
		return interfaceName
	}

	symbolsMapKey := strings.TrimPrefix(protoPackageName, ".")
	symbolsMap, ok := r.filesForExportedPackageSymbols[symbolsMapKey]
	if !ok || symbolsMap == nil {
		utils.LogError("Unable to retrieve symbols map for " + protoPackageName)
	}

	symbolSourcePath, ok := symbolsMap[typeParts[lastIndex]]
	if !ok {
		utils.LogError("Unable to retrieve source path for symbol " + typeParts[lastIndex] + " in " + protoPackageName)
	}

	alternativeImportNames, ok := r.alternativeImportNames[r.currentProtoFilePath]
	if !ok || alternativeImportNames == nil {
		utils.LogError("Unable to retrieve alternative import names for " + r.currentProtoFilePath)
	}

	alternativeImportName, ok := alternativeImportNames[symbolSourcePath]
	if !ok {
		utils.LogError("Unable to retrieve alternative import name for " + symbolSourcePath + " in " + r.currentProtoFilePath)
	}

	return alternativeImportName + "." + interfaceName
}

func (r *Runner) getMessageFieldDecorator(fieldSpec *descriptorpb.FieldDescriptorProto) string {
	decorator := "@protobufjs.Field.d(" + strconv.FormatInt(int64(*fieldSpec.Number), 10) + ", " + r.getProtobufMessageFieldType(fieldSpec)

	if fieldSpec.GetLabel() == descriptorpb.FieldDescriptorProto_LABEL_REPEATED {
		decorator += ", 'repeated'"
	}

	decorator += ")"

	return decorator
}

func (r *Runner) getInterfaceFieldType(fieldSpec *descriptorpb.FieldDescriptorProto) string {
	baseType := "unknown"

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
	case descriptorpb.FieldDescriptorProto_TYPE_ENUM:
		baseType = r.getEnumOrMessageTypeName(fieldSpec.GetTypeName(), false)
	case descriptorpb.FieldDescriptorProto_TYPE_MESSAGE:
		baseType = r.getEnumOrMessageTypeName(fieldSpec.GetTypeName(), true)
	}

	if fieldSpec.GetLabel() == descriptorpb.FieldDescriptorProto_LABEL_REPEATED {
		baseType += "[]"
	}

	return baseType
}

func (r *Runner) getClassFieldType(fieldSpec *descriptorpb.FieldDescriptorProto) string {
	baseType := "unknown"

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
	case descriptorpb.FieldDescriptorProto_TYPE_ENUM:
		baseType = r.getEnumOrMessageTypeName(fieldSpec.GetTypeName(), false) + "_Enum"
	case descriptorpb.FieldDescriptorProto_TYPE_MESSAGE:
		baseType = r.getEnumOrMessageTypeName(fieldSpec.GetTypeName(), false)
	}

	if fieldSpec.GetLabel() == descriptorpb.FieldDescriptorProto_LABEL_REPEATED {
		baseType += "[]"
	}

	return baseType
}

func (r *Runner) getProtobufMessageFieldType(fieldSpec *descriptorpb.FieldDescriptorProto) string {
	baseType := "unknown"

	switch fieldSpec.GetType() {
	case descriptorpb.FieldDescriptorProto_TYPE_BOOL:
		baseType = "'bool'"
	case descriptorpb.FieldDescriptorProto_TYPE_INT32:
		baseType = "'int32'"
	case descriptorpb.FieldDescriptorProto_TYPE_UINT32:
		baseType = "'uint32'"
	case descriptorpb.FieldDescriptorProto_TYPE_SINT32:
		baseType = "'sint32'"
	case descriptorpb.FieldDescriptorProto_TYPE_INT64:
		baseType = "'int64'"
	case descriptorpb.FieldDescriptorProto_TYPE_UINT64:
		baseType = "'uint64'"
	case descriptorpb.FieldDescriptorProto_TYPE_SINT64:
		baseType = "'sint64'"
	case descriptorpb.FieldDescriptorProto_TYPE_FLOAT:
		baseType = "'float'"
	case descriptorpb.FieldDescriptorProto_TYPE_DOUBLE:
		baseType = "'double'"
	case descriptorpb.FieldDescriptorProto_TYPE_FIXED32:
		baseType = "'fixed32'"
	case descriptorpb.FieldDescriptorProto_TYPE_FIXED64:
		baseType = "'fixed64'"
	case descriptorpb.FieldDescriptorProto_TYPE_SFIXED32:
		baseType = "'sfixed32'"
	case descriptorpb.FieldDescriptorProto_TYPE_SFIXED64:
		baseType = "'sfixed64'"
	case descriptorpb.FieldDescriptorProto_TYPE_STRING:
		baseType = "'string'"
	case descriptorpb.FieldDescriptorProto_TYPE_BYTES:
		baseType = "'bytes'"
	case descriptorpb.FieldDescriptorProto_TYPE_ENUM:
		baseType = r.getEnumOrMessageTypeName(fieldSpec.GetTypeName(), false) + "_Enum"
	case descriptorpb.FieldDescriptorProto_TYPE_MESSAGE:
		baseType = r.getEnumOrMessageTypeName(fieldSpec.GetTypeName(), false)
	}

	return baseType
}
