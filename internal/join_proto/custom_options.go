package join_proto

import (
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/descriptorpb"
)

type IOptions interface {
	ProtoReflect() protoreflect.Message
	Reset()
}

// Returns: (result, foundOption)
func GetBooleanCustomFileOption(optionName string, options *descriptorpb.FileOptions, extensionTypes *protoregistry.Types) (bool, bool) {
	// It would be ideal to merge this method with GetBooleanCustomFieldOption, and use IOptions instead of *descriptorpb.FileOptions,
	// but in the case of "field options", there's a problematic segmentation fault due to an unknown reason.

	if options == nil {
		return false, false
	}

	buffer, err := proto.Marshal(options)
	if err != nil {
		panic(err)
	}

	options.Reset()
	err = proto.UnmarshalOptions{Resolver: extensionTypes}.Unmarshal(buffer, options)
	if err != nil {
		panic(err)
	}

	result, found := false, false
	options.ProtoReflect().Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
		if !fd.IsExtension() {
			return true
		}
		if fd.Name() == protoreflect.Name(optionName) {
			found = true
			result = v.Bool()
			return false
		}
		return true
	})

	return result, found
}

// Returns: (result, foundOption)
func GetBooleanCustomFieldOption(optionName string, options *descriptorpb.FieldOptions, extensionTypes *protoregistry.Types) (bool, bool) {
	// It would be ideal to merge this method with GetBooleanCustomFileOption, and use IOptions instead of *descriptorpb.FileOptions,
	// but in the case of "field options", there's a problematic segmentation fault due to an unknown reason.

	if options == nil {
		return false, false
	}

	buffer, err := proto.Marshal(options)
	if err != nil {
		panic(err)
	}

	options.Reset()
	err = proto.UnmarshalOptions{Resolver: extensionTypes}.Unmarshal(buffer, options)
	if err != nil {
		panic(err)
	}

	result, found := false, false
	options.ProtoReflect().Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
		if !fd.IsExtension() {
			return true
		}
		if fd.Name() == protoreflect.Name(optionName) {
			found = true
			result = v.Bool()
			return false
		}
		return true
	})

	return result, found
}
