package generator

import (
	"github.com/iancoleman/strcase"
	"github.com/join-com/protoc-gen-ts/internal/join_proto"
	"github.com/join-com/protoc-gen-ts/internal/utils"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/descriptorpb"
)

func (r *Runner) generateTypescriptMessageClasses(generatedFileStream *protogen.GeneratedFile, protoFile *protogen.File) {
	for _, messageSpec := range protoFile.Proto.GetMessageType() {
		r.generateTypescriptMessageClass(generatedFileStream, messageSpec)
	}
}

func (r *Runner) generateTypescriptMessageClass(generatedFileStream *protogen.GeneratedFile, messageSpec *descriptorpb.DescriptorProto) {
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
	hasEnums := r.messageHasEnumsOrDates(messageSpec)
	implementedInterfaces := "ConvertibleTo<I" + className + ">"
	if !hasEnums {
		implementedInterfaces += ", I" + className
	}
	r.P(
		generatedFileStream,
		"@protobufjs.Type.d('"+className+"')",
		"export class "+className+" extends protobufjs.Message<"+className+"> implements "+implementedInterfaces+" {\n",
	)
	r.indentLevel += 2

	for _, fieldSpec := range messageSpec.GetField() {
		r.generateTypescriptClassField(generatedFileStream, fieldSpec, messageSpec, messageOptions, requiredFields)
	}

	r.generateTypescriptClassPatchedMethods(generatedFileStream, messageSpec, requiredFields, hasEnums)

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
	requiredField, foundRequired := join_proto.GetBooleanCustomFieldOption("typescript_required", fieldOptions, r.extensionTypes)
	optionalField, foundOptional := join_proto.GetBooleanCustomFieldOption("typescript_optional", fieldOptions, r.extensionTypes)
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

func (r *Runner) generateTypescriptClassPatchedMethods(generatedFileStream *protogen.GeneratedFile, messageSpec *descriptorpb.DescriptorProto, requiredFields bool, hasEnums bool) {
	r.generateAsInterfaceMethod(generatedFileStream, messageSpec, requiredFields, hasEnums)
	r.generateFromInterfaceMethod(generatedFileStream, messageSpec, requiredFields, hasEnums)

	r.generateDecodePatchedMethod(generatedFileStream, messageSpec, requiredFields, hasEnums)
	r.generateEncodePatchedMethod(generatedFileStream, messageSpec, requiredFields, hasEnums)
}

func (r *Runner) generateAsInterfaceMethod(generatedFileStream *protogen.GeneratedFile, messageSpec *descriptorpb.DescriptorProto, requiredFields bool, hasEnums bool) {
	className := strcase.ToCamel(messageSpec.GetName())
	r.P(generatedFileStream, "public asInterface(): I"+className+" {")
	r.indentLevel += 2

	if hasEnums {
		r.P(generatedFileStream, "return {")
		r.indentLevel += 2

		r.P(generatedFileStream, "...this,")
		for _, fieldSpec := range messageSpec.GetField() {
			r.generatePatchedInterfaceField(generatedFileStream, fieldSpec, messageSpec, requiredFields)
		}

		r.indentLevel -= 2
		r.P(generatedFileStream, "}")
	} else {
		r.P(generatedFileStream, "return this")
	}

	r.indentLevel -= 2
	r.P(generatedFileStream, "}\n")
}

func (r *Runner) generateFromInterfaceMethod(generatedFileStream *protogen.GeneratedFile, messageSpec *descriptorpb.DescriptorProto, requiredFields bool, hasEnums bool) {
	className := strcase.ToCamel(messageSpec.GetName())
	r.P(generatedFileStream, "public static fromInterface(value: I"+className+"): "+className+" {")
	r.indentLevel += 2

	if hasEnums {

		r.P(generatedFileStream, "const patchedValue = {")
		r.indentLevel += 2

		r.P(generatedFileStream, "...value,")
		for _, fieldSpec := range messageSpec.GetField() {
			r.generateUnpatchedInterfaceField(generatedFileStream, fieldSpec, messageSpec, requiredFields)
		}

		r.indentLevel -= 2
		r.P(generatedFileStream, "}\n")
		r.P(generatedFileStream, "return "+className+".fromObject(patchedValue)")
	} else {
		r.P(generatedFileStream, "return "+className+".fromObject(value)")
	}

	r.indentLevel -= 2
	r.P(generatedFileStream, "}\n")
}

func (r *Runner) generateDecodePatchedMethod(generatedFileStream *protogen.GeneratedFile, messageSpec *descriptorpb.DescriptorProto, requiredFields bool, hasEnums bool) {
	className := strcase.ToCamel(messageSpec.GetName())
	r.P(generatedFileStream, "public static decodePatched(reader: protobufjs.Reader | Uint8Array): I"+className+" {")
	r.indentLevel += 2

	if hasEnums {
		r.P(generatedFileStream, "return "+className+".decode(reader).asInterface()")
	} else {
		r.P(generatedFileStream, "return "+className+".decode(reader)")
	}

	r.indentLevel -= 2
	r.P(generatedFileStream, "}\n")
}

func (r *Runner) generateEncodePatchedMethod(generatedFileStream *protogen.GeneratedFile, messageSpec *descriptorpb.DescriptorProto, requiredFields bool, hasEnums bool) {
	className := strcase.ToCamel(messageSpec.GetName())

	// public static encode<T extends Message<T>>(this: Constructor<T>, message: (T|{ [k: string]: any }), writer?: Writer): Writer;
	r.P(generatedFileStream, "public static encodePatched(message: I"+className+", writer?: protobufjs.Writer): protobufjs.Writer {")
	r.indentLevel += 2

	if hasEnums {
		r.P(
			generatedFileStream,
			"const transformedMessage = "+className+".fromInterface(message)",
			"return "+className+".encode(transformedMessage, writer)",
		)
	} else {
		r.P(generatedFileStream, "return "+className+".encode(message, writer)")
	}

	r.indentLevel -= 2
	r.P(generatedFileStream, "}\n")
}

func (r *Runner) generatePatchedInterfaceField(
	generatedFileStream *protogen.GeneratedFile,
	fieldSpec *descriptorpb.FieldDescriptorProto,
	messageSpec *descriptorpb.DescriptorProto,
	requiredFields bool,
) {
	fieldType := fieldSpec.GetType()
	if fieldType != descriptorpb.FieldDescriptorProto_TYPE_ENUM && fieldType != descriptorpb.FieldDescriptorProto_TYPE_MESSAGE {
		return
	}

	fieldOptions := fieldSpec.GetOptions()
	requiredField, foundRequired := join_proto.GetBooleanCustomFieldOption("typescript_required", fieldOptions, r.extensionTypes)
	optionalField, foundOptional := join_proto.GetBooleanCustomFieldOption("typescript_optional", fieldOptions, r.extensionTypes)
	if foundRequired && requiredField && foundOptional && optionalField {
		utils.LogError("incompatible options for field " + fieldSpec.GetName() + " in " + messageSpec.GetName())
	}

	confirmRequired := requiredFields && !(foundOptional && optionalField) || foundRequired && requiredField
	isRepeated := fieldSpec.GetLabel() == descriptorpb.FieldDescriptorProto_LABEL_REPEATED

	value := "this." + fieldSpec.GetJsonName()
	if fieldType == descriptorpb.FieldDescriptorProto_TYPE_ENUM {
		unionTypeName := r.getEnumOrMessageTypeName(fieldSpec.GetTypeName(), false)
		enumTypeName := unionTypeName + "_Enum"
		if confirmRequired {
			if isRepeated {
				value += ".map((e) => " + enumTypeName + "[e]! as " + unionTypeName + "),"
			} else {
				value = enumTypeName + "[" + value + "]! as " + unionTypeName + ","
			}
		} else {
			if isRepeated {
				value += "?.map((e) => " + enumTypeName + "[e]! as " + unionTypeName + "),"
			} else {
				value = "((" + value + " != null) ? (" + enumTypeName + "[" + value + "]!) : undefined) as " + unionTypeName + " | undefined,"
			}
		}
	} else {
		fieldTypeName := fieldSpec.GetTypeName()
		nestedMessageSpec, ok := r.messageSpecsByFQN[fieldTypeName]
		if !ok || nestedMessageSpec == nil {
			utils.LogError("Unable to retrieve message spec for " + fieldTypeName)
		}
		if fieldTypeName != ".google.protobuf.Timestamp" && !r.messageHasEnumsOrDates(nestedMessageSpec) {
			return
		}

		if confirmRequired {
			if fieldTypeName == ".google.protobuf.Timestamp" {
				if isRepeated {
					value += ".map((ts) => new Date((ts.seconds ?? 0) * 1000 + (ts.nanos ?? 0) / 1000000)),"
				} else {
					value = "new Date((" + value + ".seconds ?? 0) * 1000 + (" + value + ".nanos ?? 0) / 1000000),"
				}
			} else {
				if isRepeated {
					value += ".map((o) => o.asInterface()),"
				} else {
					value += ".asInterface(),"
				}
			}
		} else {
			if fieldTypeName == ".google.protobuf.Timestamp" {
				if isRepeated {
					value += "?.map((ts) => new Date((ts.seconds ?? 0) * 1000 + (ts.nanos ?? 0) / 1000000)),"
				} else {
					value = value + " != null ? new Date((" + value + ".seconds ?? 0) * 1000 + (" + value + ".nanos ?? 0) / 1000000) : undefined,"
				}
			} else {
				if isRepeated {
					value += "?.map((o) => o.asInterface()),"
				} else {
					value += "?.asInterface(),"
				}
			}
		}
	}

	r.P(generatedFileStream, fieldSpec.GetJsonName()+": "+value)
}

func (r *Runner) generateUnpatchedInterfaceField(
	generatedFileStream *protogen.GeneratedFile,
	fieldSpec *descriptorpb.FieldDescriptorProto,
	messageSpec *descriptorpb.DescriptorProto,
	requiredFields bool,
) {
	fieldType := fieldSpec.GetType()
	if fieldType != descriptorpb.FieldDescriptorProto_TYPE_ENUM && fieldType != descriptorpb.FieldDescriptorProto_TYPE_MESSAGE {
		return
	}

	fieldOptions := fieldSpec.GetOptions()
	requiredField, foundRequired := join_proto.GetBooleanCustomFieldOption("typescript_required", fieldOptions, r.extensionTypes)
	optionalField, foundOptional := join_proto.GetBooleanCustomFieldOption("typescript_optional", fieldOptions, r.extensionTypes)
	if foundRequired && requiredField && foundOptional && optionalField {
		utils.LogError("incompatible options for field " + fieldSpec.GetName() + " in " + messageSpec.GetName())
	}

	confirmRequired := requiredFields && !(foundOptional && optionalField) || foundRequired && requiredField
	isRepeated := fieldSpec.GetLabel() == descriptorpb.FieldDescriptorProto_LABEL_REPEATED

	value := "value." + fieldSpec.GetJsonName()
	if fieldType == descriptorpb.FieldDescriptorProto_TYPE_ENUM {
		unionTypeName := r.getEnumOrMessageTypeName(fieldSpec.GetTypeName(), false)
		enumTypeName := unionTypeName + "_Enum"
		if confirmRequired {
			if isRepeated {
				value += ".map((e) => " + enumTypeName + "[e]!),"
			} else {
				value = enumTypeName + "[" + value + "]!,"
			}
		} else {
			if isRepeated {
				value += "?.map((e) => " + enumTypeName + "[e]!),"
			} else {
				value = "((" + value + " != null) ? (" + enumTypeName + "[" + value + "]!) : undefined) as " + enumTypeName + " | undefined,"
			}
		}
	} else {
		fieldTypeName := fieldSpec.GetTypeName()
		nestedMessageSpec, ok := r.messageSpecsByFQN[fieldTypeName]
		if !ok || nestedMessageSpec == nil {
			utils.LogError("Unable to retrieve message spec for " + fieldTypeName)
		}
		if fieldTypeName != ".google.protobuf.Timestamp" && !r.messageHasEnumsOrDates(nestedMessageSpec) {
			return
		}

		className := r.getEnumOrMessageTypeName(fieldTypeName, false)
		if confirmRequired {
			if fieldTypeName == ".google.protobuf.Timestamp" {
				if isRepeated {
					value += ".map((d) => " + className + ".fromInterface({ seconds: Math.floor(d.getTime() / 1000), nanos: d.getMilliseconds() * 1000000 })),"
				} else {
					value = className + ".fromInterface({ seconds: Math.floor(" + value + ".getTime() / 1000), nanos: " + value + ".getMilliseconds() * 1000000 }),"
				}
			} else {
				if isRepeated {
					value += ".map((o) => " + className + ".fromInterface(o)),"
				} else {
					value = className + ".fromInterface(" + value + "),"
				}
			}
		} else {
			if fieldTypeName == ".google.protobuf.Timestamp" {
				if isRepeated {
					value += "?.map((d) => " + className + ".fromInterface({ seconds: Math.floor(d.getTime() / 1000), nanos: d.getMilliseconds() * 1000000 })),"
				} else {
					value = "(" + value + " != null) ? " + className + ".fromInterface({ seconds: Math.floor(" + value + ".getTime() / 1000), nanos: " + value + ".getMilliseconds() * 1000000 }) : undefined,"
				}
			} else {
				if isRepeated {
					value += "?.map((o) => " + className + ".fromInterface(o)),"
				} else {
					value = "(" + value + " != null) ? " + className + ".fromInterface(" + value + ") : undefined,"
				}
			}
		}
	}

	r.P(generatedFileStream, fieldSpec.GetJsonName()+": "+value)
}

func (r *Runner) messageHasEnumsOrDates(messageSpec *descriptorpb.DescriptorProto) bool {
	for _, fieldSpec := range messageSpec.GetField() {
		switch t := fieldSpec.GetType(); t {
		case descriptorpb.FieldDescriptorProto_TYPE_ENUM:
			return true
		case descriptorpb.FieldDescriptorProto_TYPE_MESSAGE:
			fieldTypeName := fieldSpec.GetTypeName()
			if fieldTypeName == ".google.protobuf.Timestamp" {
				return true
			}

			nestedMessageSpec, ok := r.messageSpecsByFQN[fieldTypeName]
			if !ok || nestedMessageSpec == nil {
				utils.LogError("Unable to retrieve message spec for " + fieldTypeName)
			}

			if r.messageHasEnumsOrDates(nestedMessageSpec) {
				return true
			}
		}
	}

	return false
}