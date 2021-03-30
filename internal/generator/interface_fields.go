package generator

import (
	"strings"

	"google.golang.org/protobuf/types/descriptorpb"
)

func (r *Runner) getMessageFieldType(messageField *descriptorpb.FieldDescriptorProto) string {
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
		baseType = r.getEnumOrMessageTypeName(messageField.GetTypeName(), false)
	case descriptorpb.FieldDescriptorProto_TYPE_MESSAGE:
		baseType = r.getEnumOrMessageTypeName(messageField.GetTypeName(), true)
	}

	if messageField.GetLabel() == descriptorpb.FieldDescriptorProto_LABEL_REPEATED {
		baseType += "[]"
	}

	return baseType
}

func (r *Runner) getEnumOrMessageTypeName(typeName string, isInterface bool) string {
	names := strings.Split(typeName, ".")

	interfaceName := names[len(names)-1]
	if isInterface {
		interfaceName = "I" + interfaceName
	}

	importName := getNamespaceFromProtoPackage(strings.Join(names[0:len(names)-1], "."))

	if importName == "" || importName == r.currentNamespace {
		return interfaceName
	}
	return importName + "." + interfaceName
}
