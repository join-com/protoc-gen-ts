package generator

import (
	"sort"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/join-com/protoc-gen-ts/internal/join_proto"
	"github.com/join-com/protoc-gen-ts/internal/utils"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/descriptorpb"
)

func (r *Runner) generateTypescriptMessageClasses(generatedFileStream *protogen.GeneratedFile, protoFile *protogen.File) {
	for _, messageSpec := range r.getTopologicallySortedMessages(protoFile.Proto.GetMessageType()) {
		r.generateTypescriptMessageClass(generatedFileStream, messageSpec)
	}
}

func (r *Runner) getTopologicallySortedMessages(messageSpecs []*descriptorpb.DescriptorProto) []*descriptorpb.DescriptorProto {
	if len(messageSpecs) <= 1 {
		return messageSpecs // No need to go to sort in this case
	}

	//                   referrer class               ->   referred classes
	refsMap := make(map[*descriptorpb.DescriptorProto]map[*descriptorpb.DescriptorProto]bool)

	// First step: generate the graph
	// We'll populate the implicit "graph" (represented by a double hashmap)
	for _, messageSpec := range messageSpecs {
		messageMap, ok := refsMap[messageSpec]
		if !ok || messageMap == nil {
			messageMap = make(map[*descriptorpb.DescriptorProto]bool)
		}

		for _, fieldSpec := range messageSpec.GetField() {
			if fieldSpec.GetType() != descriptorpb.FieldDescriptorProto_TYPE_MESSAGE {
				// We only care about "messages", because we already ensured that enums were generated before
				continue
			}

			fieldTypeName := fieldSpec.GetTypeName()
			fieldTypeNamespace := r.getNamespaceFromTypeName(fieldTypeName)

			if fieldTypeNamespace != r.currentNamespace {
				// We only care about messages declared in the same package/namespace, the rest have already been
				// imported, and therefore they are available, so we don't have to worry about topological order.
				continue
			}

			nestedMessageSpec, ok := r.messageSpecsByFQN[fieldTypeName]
			if !ok || nestedMessageSpec == nil {
				utils.LogError("Unable to retrieve message spec for " + fieldTypeName)
			}

			messageMap[nestedMessageSpec] = true
		}

		refsMap[messageSpec] = messageMap
	}

	// We need to presort by name to obtain a stable order, as topological order is not absolute
	orderedMessageSpecsByName := make([]*descriptorpb.DescriptorProto, len(messageSpecs))
	copy(orderedMessageSpecsByName, messageSpecs)
	sort.Slice(orderedMessageSpecsByName, func(i, j int) bool {
		return orderedMessageSpecsByName[i].GetName() < orderedMessageSpecsByName[j].GetName()
	})

	processedMessages := make(map[*descriptorpb.DescriptorProto]bool)

	// Second step, iteratively remove items from the graph and add them to our topologically sorted list
	orderedMessageSpecs := make([]*descriptorpb.DescriptorProto, 0, len(messageSpecs))
	for len(orderedMessageSpecs) < len(messageSpecs) {
		for _, referrer := range orderedMessageSpecsByName {
			referredCollection, ok := refsMap[referrer]
			if !ok || referredCollection == nil {
				utils.LogError("Unable to retrieve referred collection in topological sort for " + referrer.GetName())
			}

			if len(referredCollection) > 0 || processedMessages[referrer] {
				continue
			}
			orderedMessageSpecs = append(orderedMessageSpecs, referrer)
			processedMessages[referrer] = true

			// We can safely remote that item for the rest of items' dependencies
			for _referrer, _referredCollection := range refsMap {
				if _referrer == referrer {
					continue
				}

				_, present := _referredCollection[referrer]
				if present {
					delete(_referredCollection, referrer)
				}
			}
		}
	}

	return orderedMessageSpecs
}

func (r *Runner) getNamespaceFromTypeName(typeName string) string {
	typeParts := strings.Split(typeName, ".")
	lastIndex := len(typeParts) - 1

	protoPackageName := strings.Join(typeParts[0:lastIndex], ".")
	namespace := getNamespaceFromProtoPackage(protoPackageName)

	return namespace
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

	classDecoratorName := "@protobufjs.Type.d"
	if r.currentPackage == "google.protobuf" {
		classDecoratorName = "@registerCommonClass"
	}

	r.P(
		generatedFileStream,
		classDecoratorName+"('"+strings.Replace(r.currentPackage, ".", "_", -1)+"_"+className+"')",
		"export class "+className+" extends protobufjs.Message<"+className+"> implements "+implementedInterfaces+" {\n",
	)
	r.indentLevel += 2

	for _, fieldSpec := range messageSpec.GetField() {
		r.generateTypescriptClassField(generatedFileStream, fieldSpec, messageSpec, requiredFields)
	}

	r.generateTypescriptClassPatchedMethods(generatedFileStream, messageSpec, requiredFields, hasEnums)

	r.indentLevel -= 2
	r.P(generatedFileStream, "}\n")
}

func (r *Runner) generateTypescriptClassField(
	generatedFileStream *protogen.GeneratedFile,
	fieldSpec *descriptorpb.FieldDescriptorProto,
	messageSpec *descriptorpb.DescriptorProto,
	requiredFields bool,
) {
	fieldOptions := fieldSpec.GetOptions()
	if fieldOptions != nil {
		if fieldOptions.GetDeprecated() {
			r.P(generatedFileStream, "/**\n  * @deprecated\n */")
		}
	}

	separator := "?: "
	requiredField, foundRequired := join_proto.GetBooleanCustomFieldOption("typescript_required", fieldOptions, r.extensionTypes)
	optionalField, foundOptional := join_proto.GetBooleanCustomFieldOption("typescript_optional", fieldOptions, r.extensionTypes)
	if foundRequired && requiredField && foundOptional && optionalField {
		utils.LogError("incompatible options for field " + fieldSpec.GetName() + " in " + messageSpec.GetName())
	}

	isRequiredField := requiredFields && !(foundOptional && optionalField) || foundRequired && requiredField
	if isRequiredField {
		separator = "!: "
	}

	r.P(
		generatedFileStream,
		r.getMessageFieldDecorator(fieldSpec, isRequiredField),
		"public "+fieldSpec.GetJsonName()+separator+r.getClassFieldType(fieldSpec)+"\n",
	)
}

func (r *Runner) generateTypescriptClassPatchedMethods(generatedFileStream *protogen.GeneratedFile, messageSpec *descriptorpb.DescriptorProto, requiredFields bool, hasEnums bool) {
	r.generateAsInterfaceMethod(generatedFileStream, messageSpec, requiredFields)
	r.generateFromInterfaceMethod(generatedFileStream, messageSpec, requiredFields, hasEnums)

	r.generateDecodePatchedMethod(generatedFileStream, messageSpec, requiredFields)
	r.generateEncodePatchedMethod(generatedFileStream, messageSpec, requiredFields, hasEnums)
}

func (r *Runner) generateAsInterfaceMethod(generatedFileStream *protogen.GeneratedFile, messageSpec *descriptorpb.DescriptorProto, requiredFields bool) {
	className := strcase.ToCamel(messageSpec.GetName())
	r.P(generatedFileStream, "public asInterface(): I"+className+" {")
	r.indentLevel += 2

	hasRepeated := false
	for _, fieldSpec := range messageSpec.GetField() {
		hasRepeated = hasRepeated || (fieldSpec.GetLabel() == descriptorpb.FieldDescriptorProto_LABEL_REPEATED)
	}

	r.P(generatedFileStream, "const message = {")
	r.indentLevel += 2
	r.P(generatedFileStream, "...this,")
	for _, fieldSpec := range messageSpec.GetField() {
		r.generatePatchedInterfaceField(generatedFileStream, fieldSpec, messageSpec, requiredFields)
	}
	r.indentLevel -= 2
	r.P(generatedFileStream, "}")

	var fieldAssignment string
	var loopInnerIfClause string
	if hasRepeated {
		fieldAssignment = "  const field = message[fieldName as keyof I" + className + "]"
		loopInnerIfClause = "field == null || Array.isArray(field) && field.length === 0"
	} else {
		fieldAssignment = ""
		loopInnerIfClause = "message[fieldName as keyof I" + className + "] == null"
	}

	r.P(
		generatedFileStream,
		"for (const fieldName of Object.keys(message)) {",
		fieldAssignment,
		"  if ("+loopInnerIfClause+") {",
		"    // We remove the key to avoid problems with code making too many assumptions",
		"    delete message[fieldName as keyof I"+className+"]",
		"  }",
		"}",
	)

	r.P(generatedFileStream, "return message")

	r.indentLevel -= 2
	r.P(generatedFileStream, "}\n")
}

func (r *Runner) generateFromInterfaceMethod(generatedFileStream *protogen.GeneratedFile, messageSpec *descriptorpb.DescriptorProto, requiredFields bool, hasEnums bool) {
	className := strcase.ToCamel(messageSpec.GetName())
	r.P(generatedFileStream, "public static fromInterface(this: void, value: I"+className+"): "+className+" {")
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

func (r *Runner) generateDecodePatchedMethod(generatedFileStream *protogen.GeneratedFile, messageSpec *descriptorpb.DescriptorProto, requiredFields bool) {
	className := strcase.ToCamel(messageSpec.GetName())
	r.P(generatedFileStream, "public static decodePatched(this: void, reader: protobufjs.Reader | Uint8Array): I"+className+" {")
	r.indentLevel += 2

	messageRequiredFields, hasRepeated := r.getMessageRequiredFields(messageSpec, requiredFields)
	if len(messageRequiredFields) > 0 {
		var fieldAssignment string
		var loopInerIfClause string
		if hasRepeated {
			fieldAssignment = "  const field = message[fieldName]"
			loopInerIfClause = "field == null || Array.isArray(field) && field.length === 0"
		} else {
			fieldAssignment = ""
			loopInerIfClause = "message[fieldName] == null"
		}

		r.P(generatedFileStream, "const message = "+className+".decode(reader).asInterface()")
		r.P(
			generatedFileStream,
			"for (const fieldName of ["+strings.Join(messageRequiredFields, ", ")+"] as (keyof I"+className+")[]) {",
			fieldAssignment,
			"  if ("+loopInerIfClause+") {",
			"    throw new Error(`Required field ${fieldName} in "+className+" is null or undefined`)",
			"  }",
			"}",
			"return message",
		)
	} else {
		r.P(generatedFileStream, "return "+className+".decode(reader).asInterface()")
	}

	r.indentLevel -= 2
	r.P(generatedFileStream, "}\n")
}

func (r *Runner) getMessageRequiredFields(messageSpec *descriptorpb.DescriptorProto, requiredFields bool) ([]string, bool) {
	hasRepeated := false
	messageRequiredFields := make([]string, 0)

	for _, fieldSpec := range messageSpec.GetField() {
		fieldOptions := fieldSpec.GetOptions()
		requiredField, foundRequired := join_proto.GetBooleanCustomFieldOption("typescript_required", fieldOptions, r.extensionTypes)
		optionalField, foundOptional := join_proto.GetBooleanCustomFieldOption("typescript_optional", fieldOptions, r.extensionTypes)
		if foundRequired && requiredField && foundOptional && optionalField {
			utils.LogError("incompatible options for field " + fieldSpec.GetName() + " in " + messageSpec.GetName())
		}
		confirmRequired := requiredFields && !(foundOptional && optionalField) || foundRequired && requiredField

		if confirmRequired {
			messageRequiredFields = append(messageRequiredFields, "'"+fieldSpec.GetJsonName()+"'") // We add the quotation for convenience
		}
		hasRepeated = hasRepeated || (fieldSpec.GetLabel() == descriptorpb.FieldDescriptorProto_LABEL_REPEATED)
	}

	return messageRequiredFields, hasRepeated
}

func (r *Runner) generateEncodePatchedMethod(generatedFileStream *protogen.GeneratedFile, messageSpec *descriptorpb.DescriptorProto, requiredFields bool, hasEnums bool) {
	className := strcase.ToCamel(messageSpec.GetName())

	r.P(generatedFileStream, "public static encodePatched(this: void, message: I"+className+", writer?: protobufjs.Writer): protobufjs.Writer {")
	r.indentLevel += 2

	messageRequiredFields, hasRepeated := r.getMessageRequiredFields(messageSpec, requiredFields)
	if len(messageRequiredFields) > 0 {
		var fieldAssignment string
		var loopInerIfClause string
		if hasRepeated {
			fieldAssignment = "  const field = message[fieldName]"
			loopInerIfClause = "field == null || Array.isArray(field) && field.length === 0"
		} else {
			fieldAssignment = ""
			loopInerIfClause = "message[fieldName] == null"
		}

		r.P(
			generatedFileStream,
			"for (const fieldName of ["+strings.Join(messageRequiredFields, ", ")+"] as (keyof I"+className+")[]) {",
			fieldAssignment,
			"  if ("+loopInerIfClause+") {",
			"    throw new Error(`Required field ${fieldName} in "+className+" is null or undefined`)",
			"  }",
			"}",
		)
	}
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
