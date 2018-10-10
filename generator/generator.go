package generator

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode"

	"github.com/antonversal/protoc-gen-ts/base"
	"github.com/golang/protobuf/proto"
	google_protobuf "github.com/golang/protobuf/protoc-gen-go/descriptor"
	gen "github.com/golang/protobuf/protoc-gen-go/generator"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
)

type generator struct {
	*base.Generator
	protoFile *google_protobuf.FileDescriptorProto
}

func New() *generator {
	return &generator{Generator: base.New()}
}

func (g *generator) tsFileName(protoName *string) string {
	fileNames := strings.Split(*protoName, "/")
	fileBaseName := fileNames[len(fileNames)-1]
	return strings.Join(fileNames[:len(fileNames)-1], "/") + "/" + gen.CamelCase(g.ProtoFileBaseName(fileBaseName))
}

func (g *generator) tsFileNameWithExt(protoName *string) string {
	return g.tsFileName(protoName) + ".ts"
}

func (g *generator) generateGenericImports() {
	g.P(`
import * as grpc from 'grpc';
import * as grpcts from '@join-com/grpc-ts';
`)
}

func (g *generator) packageNameFromFullName(fullName string) string {
	parts := strings.Split(fullName, ".")
	partsWithotTypeName := parts[:len(parts)-1]
	return strings.Join(partsWithotTypeName, ".")
}

func (g *generator) generateImports(protoFile *google_protobuf.FileDescriptorProto, protoFiles []*google_protobuf.FileDescriptorProto) {
	for _, dependency := range protoFile.Dependency {
		var depProtoFile *google_protobuf.FileDescriptorProto
		for _, extProtoFile := range protoFiles {
			depProtoFileName := *extProtoFile.Name

			if depProtoFileName == dependency {
				depProtoFile = extProtoFile
			}
		}

		packageName := depProtoFile.GetPackage()
		namespaceName := g.namespaceName(packageName)
		fileNames := strings.Split(*protoFile.Name, "/")
		var path string
		if len(fileNames) > 0 {
			path = "./"
		} else {
			path = strings.Repeat("../", len(fileNames)-1)
		}
		path += g.tsFileName(depProtoFile.Name)
		importName := g.GetImportName(*protoFile.Name, *depProtoFile.Name)
		if importName == namespaceName {
			g.P(fmt.Sprintf("import { %s } from '%s';", namespaceName, path))
		} else {
			g.P(fmt.Sprintf("import { %s as %s} from '%s';", namespaceName, importName, path))
		}

	}
}

func (g *generator) validateParameters() {
	if _, ok := g.Param["proto_relative"]; !ok {
		g.Fail("parameter `proto_relative` is required (e.g. --ts_out=proto_relative=<proto relative to generated path>:<path to generated files>)")
	}
}

// copy
func (g *generator) namespaceName(packageName string) string {
	splits := strings.Split(packageName, ".")
	camelCaseName := ""
	for _, name := range splits {
		a := []string{camelCaseName, gen.CamelCase(name)}
		camelCaseName = strings.Join(a, "")
	}

	return camelCaseName
}

func (g *generator) generateNamespace(packageName string) {
	g.P(fmt.Sprintf("export namespace %s {", g.namespaceName(packageName)))
}

func (g *generator) messageName(message *google_protobuf.DescriptorProto) string {
	return gen.CamelCase(*message.Name)
}

func (g *generator) enumName(enum *google_protobuf.EnumDescriptorProto) string {
	return gen.CamelCase(*enum.Name)
}

func (g *generator) getTsFieldTypeForScalar(typeID google_protobuf.FieldDescriptorProto_Type) string {
	m := make(map[google_protobuf.FieldDescriptorProto_Type]string)
	m[google_protobuf.FieldDescriptorProto_TYPE_DOUBLE] = "number"    // TYPE_DOUBLE
	m[google_protobuf.FieldDescriptorProto_TYPE_FLOAT] = "number"     // TYPE_FLOAT
	m[google_protobuf.FieldDescriptorProto_TYPE_INT64] = "number"     // TYPE_INT64
	m[google_protobuf.FieldDescriptorProto_TYPE_UINT64] = "number"    // TYPE_UINT64
	m[google_protobuf.FieldDescriptorProto_TYPE_INT32] = "number"     // TYPE_INT32
	m[google_protobuf.FieldDescriptorProto_TYPE_FIXED64] = "number"   // TYPE_FIXED64
	m[google_protobuf.FieldDescriptorProto_TYPE_FIXED32] = "number"   // TYPE_FIXED32
	m[google_protobuf.FieldDescriptorProto_TYPE_BOOL] = "boolean"     // TYPE_BOOL
	m[google_protobuf.FieldDescriptorProto_TYPE_STRING] = "string"    // TYPE_STRING
	m[google_protobuf.FieldDescriptorProto_TYPE_GROUP] = "Object"     // TYPE_GROUP
	m[google_protobuf.FieldDescriptorProto_TYPE_MESSAGE] = "Object"   // TYPE_MESSAGE - Length-delimited aggregate.
	m[google_protobuf.FieldDescriptorProto_TYPE_BYTES] = "Uint8Array" // TYPE_BYTES
	m[google_protobuf.FieldDescriptorProto_TYPE_UINT32] = "number"    // TYPE_UINT32
	m[google_protobuf.FieldDescriptorProto_TYPE_ENUM] = "number"      // TYPE_ENUM
	m[google_protobuf.FieldDescriptorProto_TYPE_SFIXED32] = "number"  // TYPE_SFIXED32
	m[google_protobuf.FieldDescriptorProto_TYPE_SFIXED64] = "number"  // TYPE_SFIXED64
	m[google_protobuf.FieldDescriptorProto_TYPE_SINT32] = "number"    // TYPE_SINT32 - Uses ZigZag encoding.
	m[google_protobuf.FieldDescriptorProto_TYPE_SINT64] = "number"    // TYPE_SINT64 - Uses ZigZag encoding.
	return m[typeID]
}

var wellKnownTypes = map[string]string{
	// "Any":       "any", // TODO: Check protobuf js
	// "Duration":  true,
	"Empty": "{}",
	// "Struct":    true,
	"Timestamp": "Date",
	// "Value":       true,
	// "ListValue":   true,
	"DoubleValue": "number",
	"FloatValue":  "number",
	"Int64Value":  "number",
	"UInt64Value": "number",
	"Int32Value":  "number",
	"UInt32Value": "number",
	"BoolValue":   "boolean",
	"StringValue": "string",
	"BytesValue":  "Uint8Array",
	// support https://github.com/google/protobuf/blob/fd046f6263fb17383cafdbb25c361e3451c31105/src/google/protobuf/struct.proto
}

func (g *generator) isFieldOptional(field *google_protobuf.FieldDescriptorProto) bool {
	return *field.Type == google_protobuf.FieldDescriptorProto_TYPE_MESSAGE || *field.Type == google_protobuf.FieldDescriptorProto_TYPE_ENUM
}

func (g *generator) isFieldOneOf(field *google_protobuf.FieldDescriptorProto) bool {
	return field.OneofIndex != nil
}

// TODO: understand how to accept more than one type
func (g *generator) isFieldDeprecated(field *google_protobuf.FieldDescriptorProto) bool {
	if field.Options == nil || field.Options.Deprecated == nil {
		return false
	}
	return *field.Options.Deprecated
}

func (g *generator) isMessageDeprecated(field *google_protobuf.DescriptorProto) bool {
	if field.Options == nil || field.Options.Deprecated == nil {
		return false
	}
	return *field.Options.Deprecated
}

func (g *generator) isMethodDeprecated(field *google_protobuf.MethodDescriptorProto) bool {
	if field.Options == nil || field.Options.Deprecated == nil {
		return false
	}
	return *field.Options.Deprecated
}

func (g *generator) isServiceDeprecated(field *google_protobuf.ServiceDescriptorProto) bool {
	if field.Options == nil || field.Options.Deprecated == nil {
		return false
	}
	return *field.Options.Deprecated
}

func (g *generator) getTsTypeFromMessage(typeName *string) string {
	names := strings.Split(*typeName, ".")
	importName := g.GetImportNameForMessage(*g.protoFile.Name, *typeName)
	if importName == "" {
		return names[len(names)-1]
	}
	return importName + "." + names[len(names)-1]
}

func (g *generator) getTsFieldType(field *google_protobuf.FieldDescriptorProto) string {
	if field.Type == nil {
		return ""
	}

	if field.TypeName != nil && strings.Contains(strings.ToLower(*field.TypeName), strings.ToLower(*field.Name+"Entry")) {
		g.Fail("proto map type is not supported")
	}

	if *field.Type == google_protobuf.FieldDescriptorProto_TYPE_MESSAGE ||
		*field.Type == google_protobuf.FieldDescriptorProto_TYPE_ENUM {
		return g.getTsTypeFromMessage(field.TypeName)
	}

	return g.getTsFieldTypeForScalar(*field.Type)
}

func (g *generator) isFieldRepeated(field *google_protobuf.FieldDescriptorProto) bool {
	return field.Label != nil && *field.Label == google_protobuf.FieldDescriptorProto_LABEL_REPEATED
}

func (g *generator) generateField(field *google_protobuf.FieldDescriptorProto, setPublic bool) {
	if g.isFieldDeprecated(field) {
		g.P("/** @deprecated */")
	}
	s := ""
	if setPublic {
		s += "public "
	}
	s += *field.JsonName
	s += "?"
	if field.GetTypeName() == ".google.protobuf.Timestamp" {
		s += ": Date"
	} else {
		s += fmt.Sprintf(": %s", g.getTsFieldType(field))
	}

	if g.isFieldRepeated(field) {
		s += "[]"
	}
	s += ";"
	g.P(s)
}

func (g *generator) generateMessage(message *google_protobuf.DescriptorProto) {
	g.P()
	if g.isMessageDeprecated(message) {
		g.P("/**")
		g.P("* @deprecated")
		g.P("*/")
	}
	g.P(fmt.Sprintf("export interface %s {", g.messageName(message)))
	for _, field := range message.Field {
		g.generateField(field, false)
	}
	g.P("}")
}

func (g *generator) generateConstructor(message *google_protobuf.DescriptorProto) {
	name := g.messageName(message)
	g.P(fmt.Sprintf("constructor(attrs?: %s){", name))
	g.P("Object.assign(this, attrs)")
	g.P("}")
}

func (g *generator) getFieldIndex(field *google_protobuf.FieldDescriptorProto) uint32 {
	wireType := g.getWireType(field)
	index := uint32(field.GetNumber())
	return ((index << 3) | wireType) >> 0
}

func (g *generator) getWireType(field *google_protobuf.FieldDescriptorProto) uint32 {
	switch *field.Type {
	case google_protobuf.FieldDescriptorProto_TYPE_DOUBLE:
		return proto.WireFixed64
	case google_protobuf.FieldDescriptorProto_TYPE_FLOAT:
		return proto.WireFixed32
	case google_protobuf.FieldDescriptorProto_TYPE_INT64,
		google_protobuf.FieldDescriptorProto_TYPE_UINT64:
		return proto.WireVarint
	case google_protobuf.FieldDescriptorProto_TYPE_INT32,
		google_protobuf.FieldDescriptorProto_TYPE_UINT32,
		google_protobuf.FieldDescriptorProto_TYPE_ENUM:
		return proto.WireVarint
	case google_protobuf.FieldDescriptorProto_TYPE_FIXED64,
		google_protobuf.FieldDescriptorProto_TYPE_SFIXED64:
		return proto.WireFixed64
	case google_protobuf.FieldDescriptorProto_TYPE_FIXED32,
		google_protobuf.FieldDescriptorProto_TYPE_SFIXED32:
		return proto.WireFixed32
	case google_protobuf.FieldDescriptorProto_TYPE_BOOL:
		return proto.WireVarint
	case google_protobuf.FieldDescriptorProto_TYPE_STRING:
		return proto.WireBytes
	case google_protobuf.FieldDescriptorProto_TYPE_GROUP:
		return proto.WireStartGroup
	case google_protobuf.FieldDescriptorProto_TYPE_MESSAGE:
		return proto.WireBytes
	case google_protobuf.FieldDescriptorProto_TYPE_BYTES:
		return proto.WireBytes
	case google_protobuf.FieldDescriptorProto_TYPE_SINT32:
		return proto.WireVarint
	case google_protobuf.FieldDescriptorProto_TYPE_SINT64:
		return proto.WireVarint
	default:
		g.Fail("undefined field type", field.Type.String())
	}
	return 2
}

func (g *generator) getWriteFunction(field *google_protobuf.FieldDescriptorProto) string {
	switch *field.Type {
	case google_protobuf.FieldDescriptorProto_TYPE_DOUBLE:
		return "double"
	case google_protobuf.FieldDescriptorProto_TYPE_FLOAT:
		return "float"
	case google_protobuf.FieldDescriptorProto_TYPE_INT64:
		return "int64"
	case google_protobuf.FieldDescriptorProto_TYPE_UINT64:
		return "uint64"
	case google_protobuf.FieldDescriptorProto_TYPE_INT32:
		return "int32"
	case google_protobuf.FieldDescriptorProto_TYPE_UINT32:
		return "uint32"
	case google_protobuf.FieldDescriptorProto_TYPE_ENUM:
		return "int32"
	case google_protobuf.FieldDescriptorProto_TYPE_FIXED64:
		return "fixed64"
	case google_protobuf.FieldDescriptorProto_TYPE_SFIXED64:
		return "sfixed64"
	case google_protobuf.FieldDescriptorProto_TYPE_FIXED32:
		return "fixed32"
	case google_protobuf.FieldDescriptorProto_TYPE_SFIXED32:
		return "sfixed32"
	case google_protobuf.FieldDescriptorProto_TYPE_BOOL:
		return "bool"
	case google_protobuf.FieldDescriptorProto_TYPE_STRING:
		return "string"
	case google_protobuf.FieldDescriptorProto_TYPE_MESSAGE:
		return "MESSAGE"
	case google_protobuf.FieldDescriptorProto_TYPE_BYTES:
		return "bytes"
	case google_protobuf.FieldDescriptorProto_TYPE_SINT32:
		return "sint32"
	case google_protobuf.FieldDescriptorProto_TYPE_SINT64:
		return "sint64"
	default:
		g.Fail("undefined field type", field.Type.String())
	}
	return "int32"
}

func (g *generator) generateEncode(message *google_protobuf.DescriptorProto) {
	g.P("public encode(writer: protobufjs.Writer = protobufjs.Writer.create()){")
	for _, field := range message.Field {
		name := *field.JsonName
		g.P(fmt.Sprintf("if (this.%s != null) {", name))
		if *field.Type == google_protobuf.FieldDescriptorProto_TYPE_MESSAGE {
			if g.isFieldRepeated(field) {
				g.P(fmt.Sprintf("for (const value of this.%s) {", name))
				g.P(fmt.Sprintf("const msg = new %sMsg(value);", g.getTsTypeFromMessage(field.TypeName)))
				g.P(fmt.Sprintf("msg.encode(writer.uint32(%d).fork()).ldelim();", g.getFieldIndex(field)))
				g.P("}")
			} else {
				g.P(fmt.Sprintf("const msg = new %sMsg(this.%s);", g.getTsTypeFromMessage(field.TypeName), name))
				g.P(fmt.Sprintf("msg.encode(writer.uint32(%d).fork()).ldelim();", g.getFieldIndex(field)))
			}
		} else {
			if g.isFieldRepeated(field) {
				g.P(fmt.Sprintf("for (const value of this.%s) {", name))
				g.P(fmt.Sprintf("writer.uint32(%d).%s(value);", g.getFieldIndex(field), g.getWriteFunction(field)))
				g.P("}")
			} else {
				g.P(fmt.Sprintf("writer.uint32(%d).%s(this.%s);", g.getFieldIndex(field), g.getWriteFunction(field), name))
			}
		}
		g.P("}")
	}
	g.P("return writer")
	g.P("}")
}

func (g *generator) enumSwitch(field *google_protobuf.FieldDescriptorProto, name string, message *google_protobuf.DescriptorProto) {
	enum := g.GetEnumTypeByName(*field.TypeName)
	g.P(fmt.Sprintf("switch (val) {"))
	typeName := g.getTsFieldType(field)
	for _, value := range enum.Value {
		g.P(fmt.Sprintf("case %d:", *value.Number))
		g.P(fmt.Sprintf("return %s.%s;", typeName, *value.Name))
	}
	g.P("default:")
	// TODO: raise error? g.P(fmt.Sprintf("throw new Error('Undefined value of enum %s for field %s of message %s');", typeName, field.GetJsonName(), message.GetName()))
	g.P("return")
	g.P("};")
}

func (g *generator) generateDecode(message *google_protobuf.DescriptorProto) {
	g.P("public static decode(inReader: Uint8Array | protobufjs.Reader, length?: number){")
	g.P("const reader = !(inReader instanceof protobufjs.Reader)")
	g.P("? protobufjs.Reader.create(inReader)")
	g.P(": inReader")
	g.P("const end = length === undefined ? reader.len : reader.pos + length;")
	g.P(fmt.Sprintf("const message = new %sMsg();", g.getTsTypeFromMessage(message.Name)))
	g.P("while (reader.pos < end) {")
	g.P("const tag = reader.uint32()")
	g.P("switch (tag >>> 3) {")
	for _, field := range message.Field {
		name := *field.JsonName
		if g.isFieldDeprecated(field) {
			g.P(fmt.Sprintf("console.warn('%s is deprecated');", name))
		}
		g.P(fmt.Sprintf("case %d:", field.GetNumber()))
		if *field.Type == google_protobuf.FieldDescriptorProto_TYPE_ENUM {
			if g.isFieldRepeated(field) {
				assign := func() {
					g.P(fmt.Sprintf("const %s = (((val) => {", name))
					g.enumSwitch(field, name, message)
					g.P("})(reader.int32()));")
					g.P(fmt.Sprintf("if (%s) {", name))
					g.P(fmt.Sprintf("message.%s.push(%s);", name, name))
					g.P("}")
				}
				g.P(fmt.Sprintf("if (!(message.%s && message.%s.length)) { message.%s = []; }", name, name, name))
				g.P("if ((tag & 7) === 2) {")
				g.P("const end2 = reader.uint32() + reader.pos;")
				g.P("while (reader.pos < end2) {")
				assign()
				g.P("}")
				g.P("} else {")
				assign()
				g.P("}")
			} else {
				g.P(fmt.Sprintf("message.%s = ((val) => {", name))
				g.enumSwitch(field, name, message)
				g.P("})(reader.int32());")
			}
		} else if *field.Type == google_protobuf.FieldDescriptorProto_TYPE_MESSAGE {
			if g.isFieldRepeated(field) {
				g.P(fmt.Sprintf("if (!(message.%s && message.%s.length)) {", name, name))
				g.P(fmt.Sprintf("message.%s = [];", name))
				g.P("}")
				if field.GetTypeName() == ".google.protobuf.Timestamp" {
					g.P(fmt.Sprintf("const %s = %sMsg.decode(reader, reader.uint32());", name, g.getTsTypeFromMessage(field.TypeName)))
					g.P(fmt.Sprintf("message.%s.push(new Date(((%s.seconds || 0) * 1000) + ((%s.nanos || 0) / 1000000)));", name, name, name))
				} else {
					g.P(fmt.Sprintf("message.%s.push(%sMsg.decode(reader, reader.uint32()));", name, g.getTsTypeFromMessage(field.TypeName)))
				}
			} else {
				if field.GetTypeName() == ".google.protobuf.Timestamp" {
					g.P(fmt.Sprintf("const %s = %sMsg.decode(reader, reader.uint32());", name, g.getTsTypeFromMessage(field.TypeName)))
					g.P(fmt.Sprintf("message.%s = new Date(((%s.seconds || 0) * 1000) + ((%s.nanos || 0) / 1000000));", name, name, name))
				} else {
					g.P(fmt.Sprintf("message.%s = %sMsg.decode(reader, reader.uint32());", name, g.getTsTypeFromMessage(field.TypeName)))
				}
			}
		} else {
			if g.isFieldRepeated(field) {
				g.P(fmt.Sprintf("if (!(message.%s && message.%s.length)) {", name, name))
				g.P(fmt.Sprintf("message.%s = [];", name))
				g.P("}")
				g.P(fmt.Sprintf("message.%s.push(reader.%s());", name, g.getWriteFunction(field)))
			} else {
				g.P(fmt.Sprintf("message.%s = reader.%s();", name, g.getWriteFunction(field)))
			}
		}
		g.P("break;")
	}
	g.P("default:")
	g.P("reader.skipType(tag & 7);")
	g.P("break;")

	g.P("}")

	g.P("}")
	g.P("return message;")

	g.P("}")
}

func (g *generator) generateGettersSetters(message *google_protobuf.DescriptorProto) {
	for _, field := range message.Field {
		name := *field.JsonName
		g.P(fmt.Sprintf("public get %s() {", name))

		// if g.isFieldDeprecated(field) {
		// 	g.P(fmt.Sprintf("console.warn('%s is deprecated');", name))
		// }
		if *field.Type == google_protobuf.FieldDescriptorProto_TYPE_ENUM {
			enum := g.GetEnumTypeByName(*field.TypeName)
			g.P(fmt.Sprintf("if (!this.%s) {", name))

			g.P("return")

			g.P("}")
			if g.isFieldRepeated(field) {
				g.P(fmt.Sprintf("return this.%s.map((val) => {", name))

				g.P("switch (val) {")
				typeName := g.getTsFieldType(field)
				for _, value := range enum.Value {
					g.P(fmt.Sprintf("case %d:", *value.Number))

					g.P(fmt.Sprintf("return %s.%s;", typeName, *value.Name))

				}
				g.P("default:")

				g.P(fmt.Sprintf("throw new Error('Undefined value of enum %s for field %s of message %s');", typeName, field.GetJsonName(), message.GetName()))

				g.P("}")

				g.P("})")
			} else {
				g.P(fmt.Sprintf("switch (this.%s) {", name))
				typeName := g.getTsFieldType(field)
				for _, value := range enum.Value {
					g.P(fmt.Sprintf("case %d:", *value.Number))

					g.P(fmt.Sprintf("return %s.%s", typeName, *value.Name))

				}
				g.P("default:")

				g.P("return")

				g.P("}")
			}
		} else if field.GetTypeName() == ".google.protobuf.Timestamp" {
			if g.isFieldRepeated(field) {
				g.P(fmt.Sprintf("return this.%s && this.%s.map(v => new Date(((v.seconds || 0) * 1000) + ((v.nanos || 0) / 1000000)));", name, name))
			} else {
				g.P(fmt.Sprintf("return this.%s && new Date(((this.%s.seconds || 0) * 1000) + ((this.%s.nanos || 0) / 1000000));", name, name, name))
			}
		} else {
			g.P(fmt.Sprintf("return this.%s;", name))
		}

		g.P("}")

		g.P(fmt.Sprintf("public set %s(val) {", name))

		if g.isFieldDeprecated(field) {
			g.P(fmt.Sprintf("console.warn('%s is deprecated');", name))
		}

		if *field.Type == google_protobuf.FieldDescriptorProto_TYPE_MESSAGE {
			if field.GetTypeName() == ".google.protobuf.Timestamp" {
				if g.isFieldRepeated(field) {
					g.P(fmt.Sprintf("this.%s = val && val.map(v => ({ seconds: Math.floor(v.getTime() / 1000), nanos: v.getMilliseconds() * 1000000 }));", name))
				} else {
					g.P(fmt.Sprintf("this.%s = val && { seconds: Math.floor(val.getTime() / 1000), nanos: val.getMilliseconds() * 1000000 };", name))
				}
			} else {
				if g.isFieldRepeated(field) {
					g.P(fmt.Sprintf("this.%s = val && val.map(v => new %sMsg(v));", name, g.getTsTypeFromMessage(field.TypeName)))
				} else {
					g.P(fmt.Sprintf("this.%s = new %sMsg(val);", name, g.getTsTypeFromMessage(field.TypeName)))
				}
			}
		} else if *field.Type == google_protobuf.FieldDescriptorProto_TYPE_ENUM {
			enum := g.GetEnumTypeByName(*field.TypeName)
			if g.isFieldRepeated(field) {
				g.P("if (!val) {")

				g.P("return;")

				g.P("}")
				g.P(fmt.Sprintf("this.%s = val.map((value) => {", name))

				g.P("switch (value) {")
				typeName := g.getTsFieldType(field)
				for _, value := range enum.Value {
					g.P(fmt.Sprintf("case %s.%s:", typeName, *value.Name))

					g.P(fmt.Sprintf("return %d;", *value.Number))

				}
				g.P("default:")

				g.P(fmt.Sprintf("throw new Error('Undefined value of enum %s for field %s of message %s');", typeName, field.GetJsonName(), message.GetName()))

				g.P("}")

				g.P("})")
			} else {
				g.P("switch (val) {")
				typeName := g.getTsFieldType(field)
				for _, value := range enum.Value {
					g.P(fmt.Sprintf("case %s.%s:", typeName, *value.Name))

					g.P(fmt.Sprintf("this.%s = %d;", name, *value.Number))
					g.P("break;")

				}
				g.P("default:")

				g.P(fmt.Sprintf("this.%s = undefined;", name))

				g.P("}")
			}
		} else {
			g.P(fmt.Sprintf("this.%s = val;", name))
		}

		g.P("}")
	}
}

func (g *generator) generateMessageClass(message *google_protobuf.DescriptorProto) {
	g.P()

	if g.isMessageDeprecated(message) {
		g.P("/**")
		g.P("* @deprecated")
		g.P("*/")
	}
	name := g.messageName(message)
	g.P(fmt.Sprintf("export class %sMsg implements %s{", name, name))

	g.generateDecode(message)
	for _, field := range message.Field {
		g.generateField(field, true)
	}
	g.generateConstructor(message)
	g.generateEncode(message)
	// g.generateGettersSetters(message)

	g.P("}")

}

func (g *generator) generateEnum(enum *google_protobuf.EnumDescriptorProto) {
	g.P()

	g.P(fmt.Sprintf("export enum %s {", g.enumName(enum)))

	for _, value := range enum.Value {
		g.P(fmt.Sprintf("%s = '%s',", *value.Name, *value.Name))
	}

	g.P(fmt.Sprint("}"))

}

func (g *generator) methodDeprecated(method *google_protobuf.MethodDescriptorProto) {
	if g.isMethodDeprecated(method) {
		g.P("/**")
		g.P("* @deprecated")
		g.P("*/")
	}
}

func (g *generator) protoPathRelative(protoFileName string) string {
	splittedName := strings.Split(protoFileName, "/")
	upPathArray := make([]string, len(splittedName)-1)
	for i := 0; i < len(splittedName)-1; i++ {
		upPathArray[i] = ".."
	}
	splittedRelative := strings.Split(g.Param["proto_relative"], "/")
	relativePathArray := make([]string, len(splittedRelative))
	for i := 0; i < len(splittedRelative); i++ {
		relativePathArray[i] = splittedRelative[i]
	}
	fullPathArray := append(upPathArray, relativePathArray...)
	return strings.Join(fullPathArray, `', '`)
}

func (g *generator) generateDefinition(service *google_protobuf.ServiceDescriptorProto) {
	g.P()

	g.P(fmt.Sprintf("export const %sServiceDefinition = {", g.toLowerFirst(*service.Name)))

	for _, method := range service.Method {
		g.P(fmt.Sprintf("%s: {", g.toLowerFirst(*method.Name)))

		g.P(fmt.Sprintf("path: '/%s/%s',", *service.Name, *method.Name))
		var clientStreaming bool
		if method.ClientStreaming == nil {
			clientStreaming = false
		} else {
			clientStreaming = *method.ClientStreaming
		}
		g.P(fmt.Sprintf("requestStream: %s,", strconv.FormatBool(clientStreaming)))
		var serverStreaming bool
		if method.ServerStreaming == nil {
			serverStreaming = false
		} else {
			serverStreaming = *method.ServerStreaming
		}
		g.P(fmt.Sprintf("responseStream: %s,", strconv.FormatBool(serverStreaming)))
		requestType := g.getTsTypeFromMessage(method.InputType)
		g.P(fmt.Sprintf("requestType: %sMsg,", requestType))
		responseType := g.getTsTypeFromMessage(method.OutputType)
		g.P(fmt.Sprintf("responseType: %sMsg,", responseType))
		g.P(fmt.Sprintf("requestSerialize: (args: %s) => new %sMsg(args).encode().finish() as Buffer,", requestType, requestType))
		g.P(fmt.Sprintf("requestDeserialize: (argBuf: Buffer) => %sMsg.decode(argBuf),", requestType))
		g.P(fmt.Sprintf("responseSerialize: (args: %s) => new %sMsg(args).encode().finish() as Buffer,", responseType, responseType))
		g.P(fmt.Sprintf("responseDeserialize: (argBuf: Buffer) => %sMsg.decode(argBuf),", responseType))

		g.P("},")
	}

	g.P("}")

}

func (g *generator) generateImplementation(service *google_protobuf.ServiceDescriptorProto) {
	g.P()

	if g.isServiceDeprecated(service) {
		g.P("/**")
		g.P("* @deprecated")
		g.P("*/")
	}
	g.P(fmt.Sprintf("export interface %sImplementation extends grpcts.Implementations {", gen.CamelCase(*service.Name)))

	for _, method := range service.Method {
		g.methodDeprecated(method)
		inputTypeName := g.getTsTypeFromMessage(method.InputType)
		outputTypeName := g.getTsTypeFromMessage(method.OutputType)
		if method.ServerStreaming != nil && *method.ServerStreaming && method.ClientStreaming != nil && *method.ClientStreaming {
			g.P(fmt.Sprintf("%s(call: grpc.ServerDuplexStream<%s, %s>): void;", g.toLowerFirst(*method.Name), inputTypeName, outputTypeName))
		} else if method.ServerStreaming != nil && *method.ServerStreaming {
			// TODO why there is no type for write stream?
			g.P(fmt.Sprintf("%s(call: grpc.ServerWriteableStream<%s>): void;", g.toLowerFirst(*method.Name), inputTypeName))
		} else if method.ClientStreaming != nil && *method.ClientStreaming {
			g.P(fmt.Sprintf("%s(call: grpc.ServerReadableStream<%s>): Promise<%s>;", g.toLowerFirst(*method.Name), inputTypeName, outputTypeName))
			g.P(fmt.Sprintf("%s(call: grpc.ServerReadableStream<%s>, callback: grpc.sendUnaryData<%s>): void;", g.toLowerFirst(*method.Name), inputTypeName, outputTypeName))
		} else {
			g.P(fmt.Sprintf("%s(call: grpc.ServerUnaryCall<%s>): Promise<%s>;", g.toLowerFirst(*method.Name), inputTypeName, outputTypeName))
			g.P(fmt.Sprintf("%s(call: grpc.ServerUnaryCall<%s>, callback: grpc.sendUnaryData<%s>): void;", g.toLowerFirst(*method.Name), inputTypeName, outputTypeName))
		}
	}

	g.P(fmt.Sprint("}"))

}

func (g *generator) generateClient(service *google_protobuf.ServiceDescriptorProto, packageName string, protoFileName string) {
	g.P()
	if g.isServiceDeprecated(service) {
		g.P("/**")
		g.P("* @deprecated")
		g.P("*/")
	}
	g.P(fmt.Sprintf("export class %sClient extends grpcts.Client {", gen.CamelCase(*service.Name)))
	for _, method := range service.Method {
		inputTypeName := g.getTsTypeFromMessage(method.InputType)
		g.methodDeprecated(method)
		if method.ServerStreaming != nil && *method.ServerStreaming && method.ClientStreaming != nil && *method.ClientStreaming {
			outputTypeName := g.getTsTypeFromMessage(method.OutputType)
			g.P(fmt.Sprintf("public %s(metadata?: grpcts.Metadata) {", g.toLowerFirst(*method.Name)))
			g.P(fmt.Sprintf("return super.makeBidiStreamRequest<%s, %s>('%s', metadata);", inputTypeName, outputTypeName, g.toLowerFirst(*method.Name)))
			g.P(fmt.Sprint("}"))
		} else if method.ServerStreaming != nil && *method.ServerStreaming {
			outputTypeName := g.getTsTypeFromMessage(method.OutputType)
			g.P(fmt.Sprintf("public %s(req: %s, metadata?: grpcts.Metadata) {", g.toLowerFirst(*method.Name), inputTypeName))

			g.P(fmt.Sprintf("return super.makeServerStreamRequest<%s, %s>('%s', req, metadata);", inputTypeName, outputTypeName, g.toLowerFirst(*method.Name)))

			g.P(fmt.Sprint("}"))
		} else if method.ClientStreaming != nil && *method.ClientStreaming {
			outputTypeName := g.getTsTypeFromMessage(method.OutputType)
			g.P(fmt.Sprintf("public %s(metadata?: grpcts.Metadata) {", g.toLowerFirst(*method.Name)))

			g.P(fmt.Sprintf("return super.makeClientStreamRequest<%s, %s>('%s', metadata);", inputTypeName, outputTypeName, g.toLowerFirst(*method.Name)))

			g.P(fmt.Sprint("};"))
		} else {
			outputTypeName := g.getTsTypeFromMessage(method.OutputType)
			g.P(fmt.Sprintf("public %s(req: %s, metadata?: grpcts.Metadata) {", g.toLowerFirst(*method.Name), inputTypeName))
			g.P(fmt.Sprintf("return super.makeUnaryRequest<%s, %s>('%s', req, metadata);", inputTypeName, outputTypeName, g.toLowerFirst(*method.Name)))
			g.P(fmt.Sprint("};"))
		}
	}

	g.P(fmt.Sprint("}"))

}

func (g *generator) Make(protoFile *google_protobuf.FileDescriptorProto, protoFiles []*google_protobuf.FileDescriptorProto) (*plugin.CodeGeneratorResponse_File, error) {
	g.protoFile = protoFile
	g.validateParameters()
	log.Print(*protoFile.Name)

	g.P("// GENERATED CODE -- DO NOT EDIT!")

	g.generateImports(protoFile, protoFiles)
	g.P("import * as protobufjs from 'protobufjs/minimal';")
	if len(protoFile.Service) > 0 {
		g.generateGenericImports()
	}
	packageName := protoFile.GetPackage()
	g.generateNamespace(packageName)

	for _, enum := range protoFile.EnumType {
		g.generateEnum(enum)
	}

	for _, message := range protoFile.MessageType {
		g.generateMessage(message)
		g.generateMessageClass(message)
	}

	for _, service := range protoFile.Service {
		g.generateDefinition(service)
		g.generateImplementation(service)
		g.generateClient(service, *protoFile.Package, *protoFile.Name)
	}

	g.P("}")

	file := &plugin.CodeGeneratorResponse_File{
		Name:    proto.String(g.tsFileNameWithExt(protoFile.Name)),
		Content: proto.String(g.String()),
	}
	g.Reset()
	return file, nil
}

func (g *generator) toLowerFirst(str string) string {
	a := []rune(str)
	a[0] = unicode.ToLower(a[0])
	return string(a)
}

func (g *generator) Generate() {
	g.Generator.Generate(g)
}
