package generator

import (
	"fmt"
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
	baseName := gen.CamelCase(g.ProtoFileBaseName(fileBaseName))
	if len(fileNames) > 1 {
		return strings.Join(fileNames[:len(fileNames)-1], "/") + "/" + baseName
	}
	return baseName
}

func (g *generator) tsFileNameWithExt(protoName *string) string {
	return g.tsFileName(protoName) + ".ts"
}

func (g *generator) generateGenericImports() {
	g.P(`
import * as grpcts from '@join-com/grpc-ts';
import * as nodeTrace from '@join-com/node-trace';
`)
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
		if len(fileNames) <= 1 {
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

func (g *generator) isFieldOneOf(field *google_protobuf.FieldDescriptorProto) bool {
	return field.OneofIndex != nil
}

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

func (g *generator) getTsTypeFromMessage(typeName *string, isInterface bool) string {
	names := strings.Split(*typeName, ".")
	importName := g.GetImportNameForMessage(*g.protoFile.Name, *typeName)
	interfaceName := names[len(names)-1]
	if isInterface {
		interfaceName = "I" + interfaceName
	}
	if importName == "" {
		return interfaceName
	}
	return importName + "." + interfaceName
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
		return g.getTsTypeFromMessage(field.TypeName, *field.Type == google_protobuf.FieldDescriptorProto_TYPE_MESSAGE)
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

func (g *generator) generateMessageInterface(message *google_protobuf.DescriptorProto) {
	g.P()
	if g.isMessageDeprecated(message) {
		g.P("/**")
		g.P("* @deprecated")
		g.P("*/")
	}
	g.P(fmt.Sprintf("export interface I%s {", g.messageName(message)))
	for _, field := range message.Field {
		g.generateField(field, false)
	}
	g.P("}")
}

func (g *generator) generateConstructor(message *google_protobuf.DescriptorProto) {
	name := g.messageName(message)
	g.P(fmt.Sprintf("constructor(attrs?: I%s){", name))
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

func (g *generator) isType64Bit(field *google_protobuf.FieldDescriptorProto) bool {
	switch *field.Type {
	case google_protobuf.FieldDescriptorProto_TYPE_INT64:
		return true
	case google_protobuf.FieldDescriptorProto_TYPE_UINT64:
		return true
	case google_protobuf.FieldDescriptorProto_TYPE_FIXED64:
		return true
	case google_protobuf.FieldDescriptorProto_TYPE_SFIXED64:
		return true
	case google_protobuf.FieldDescriptorProto_TYPE_SINT64:
		return true
	default:
		return false
	}
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

func (g *generator) encodeEnumSwitch(field *google_protobuf.FieldDescriptorProto, name string, message *google_protobuf.DescriptorProto) {
	enum := g.GetEnumTypeByName(*field.TypeName)
	g.P(fmt.Sprintf("switch (val) {"))
	for _, value := range enum.Value {
		g.P(fmt.Sprintf("case '%s':", *value.Name))
		g.P(fmt.Sprintf("return %d;", *value.Number))
	}
	g.P("default:")
	g.P("return")
	g.P("};")
}

func (g *generator) generateEncode(message *google_protobuf.DescriptorProto) {
	g.P("public encode(writer: protobufjs.Writer = protobufjs.Writer.create()){")

	if g.isMessageDeprecated(message) {
		g.P(fmt.Sprintf("logger.warn('message %s is deprecated');", *message.Name))
	}
	for _, field := range message.Field {
		name := *field.JsonName
		g.P(fmt.Sprintf("if (this.%s != null) {", name))
		if *field.Type == google_protobuf.FieldDescriptorProto_TYPE_ENUM {
			if g.isFieldRepeated(field) {
				g.P(fmt.Sprintf("for (const value of this.%s) {", name))
				g.P(fmt.Sprintf("const %s = (val => {", name))
				g.encodeEnumSwitch(field, name, message)
				g.P("})(value);")
				g.P(fmt.Sprintf("if (%s != null) {", name))
				g.P(fmt.Sprintf("writer.uint32(%d).int32(%s);", g.getFieldIndex(field), name))
				g.P("}")
				g.P("}")
			} else {
				g.P(fmt.Sprintf("const %s = (val => {", name))
				g.encodeEnumSwitch(field, name, message)
				g.P(fmt.Sprintf("})(this.%s);", name))
				g.P(fmt.Sprintf("if (%s != null) {", name))
				g.P(fmt.Sprintf("writer.uint32(%d).int32(%s);", g.getFieldIndex(field), name))
				g.P("}")
			}
		} else if *field.Type == google_protobuf.FieldDescriptorProto_TYPE_MESSAGE {
			if g.isFieldRepeated(field) {
				g.P(fmt.Sprintf("for (const value of this.%s) {", name))
				g.P("if (!value) { continue; }")
				if field.GetTypeName() == ".google.protobuf.Timestamp" {
					g.P(fmt.Sprintf("const msg = new %s({seconds: Math.floor(value.getTime() / 1000) , nanos: value.getMilliseconds() * 1000000});", g.getTsTypeFromMessage(field.TypeName, false)))
				} else {
					g.P(fmt.Sprintf("const msg = new %s(value);", g.getTsTypeFromMessage(field.TypeName, false)))
				}
				g.P(fmt.Sprintf("msg.encode(writer.uint32(%d).fork()).ldelim();", g.getFieldIndex(field)))
				g.P("}")
			} else {
				if field.GetTypeName() == ".google.protobuf.Timestamp" {
					g.P(fmt.Sprintf("const msg = new %s({seconds: Math.floor(this.%s.getTime() / 1000) , nanos: this.%s.getMilliseconds() * 1000000});", g.getTsTypeFromMessage(field.TypeName, false), name, name))
				} else {
					g.P(fmt.Sprintf("const msg = new %s(this.%s);", g.getTsTypeFromMessage(field.TypeName, false), name))
				}
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

func (g *generator) decodeEnumSwitch(field *google_protobuf.FieldDescriptorProto, name string, message *google_protobuf.DescriptorProto) {
	enum := g.GetEnumTypeByName(*field.TypeName)

	g.P(fmt.Sprintf("switch (val) {"))
	for _, value := range enum.Value {
		g.P(fmt.Sprintf("case %d:", *value.Number))
		g.P(fmt.Sprintf("return '%s';", *value.Name))
	}
	g.P("default:")
	g.P("return")
	g.P("};")
}

func (g *generator) generateDecode(message *google_protobuf.DescriptorProto) {
	g.P("public static decode(inReader: Uint8Array | protobufjs.Reader, length?: number){")
	if g.isMessageDeprecated(message) {
		g.P(fmt.Sprintf("logger.warn('message %s is deprecated');", *message.Name))
	}
	g.P("const reader = !(inReader instanceof protobufjs.Reader)")
	g.P("? protobufjs.Reader.create(inReader)")
	g.P(": inReader")
	g.P("const end = length === undefined ? reader.len : reader.pos + length;")
	g.P(fmt.Sprintf("const message = new %s();", g.getTsTypeFromMessage(message.Name, false)))
	g.P("while (reader.pos < end) {")
	g.P("const tag = reader.uint32()")
	g.P("switch (tag >>> 3) {")
	for _, field := range message.Field {
		name := *field.JsonName
		g.P(fmt.Sprintf("case %d:", field.GetNumber()))
		if g.isFieldDeprecated(field) {
			g.P(fmt.Sprintf("logger.warn('field %s is deprecated in %s');", name, *message.Name))
		}
		if *field.Type == google_protobuf.FieldDescriptorProto_TYPE_ENUM {
			if g.isFieldRepeated(field) {
				assign := func() {
					g.P(fmt.Sprintf("const %s = (((val) => {", name))
					g.decodeEnumSwitch(field, name, message)
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
				g.decodeEnumSwitch(field, name, message)
				g.P("})(reader.int32());")
			}
		} else if *field.Type == google_protobuf.FieldDescriptorProto_TYPE_MESSAGE {
			if g.isFieldRepeated(field) {
				g.P(fmt.Sprintf("if (!(message.%s && message.%s.length)) {", name, name))
				g.P(fmt.Sprintf("message.%s = [];", name))
				g.P("}")
				if field.GetTypeName() == ".google.protobuf.Timestamp" {
					g.P(fmt.Sprintf("const %s = %s.decode(reader, reader.uint32());", name, g.getTsTypeFromMessage(field.TypeName, false)))
					g.P(fmt.Sprintf("message.%s.push(new Date(((%s.seconds || 0) * 1000) + ((%s.nanos || 0) / 1000000)));", name, name, name))
				} else {
					g.P(fmt.Sprintf("message.%s.push(%s.decode(reader, reader.uint32()));", name, g.getTsTypeFromMessage(field.TypeName, false)))
				}
			} else {
				if field.GetTypeName() == ".google.protobuf.Timestamp" {
					g.P(fmt.Sprintf("const %s = %s.decode(reader, reader.uint32());", name, g.getTsTypeFromMessage(field.TypeName, false)))
					g.P(fmt.Sprintf("message.%s = new Date(((%s.seconds || 0) * 1000) + ((%s.nanos || 0) / 1000000));", name, name, name))
				} else {
					g.P(fmt.Sprintf("message.%s = %s.decode(reader, reader.uint32());", name, g.getTsTypeFromMessage(field.TypeName, false)))
				}
			}
		} else {
			if g.isFieldRepeated(field) {
				assign := func() {
					if g.isType64Bit(field) {
						g.P(fmt.Sprintf("const %s = reader.%s();", name, g.getWriteFunction(field)))
						g.P(fmt.Sprintf("message.%s.push(new protobufjs.util.LongBits(%s.low >>> 0, %s.high >>> 0).toNumber());", name, name, name))
					} else if *field.Type == google_protobuf.FieldDescriptorProto_TYPE_BYTES {
						g.P(fmt.Sprintf("message.%s.push(new Uint8Array(reader.%s()));", name, g.getWriteFunction(field)))
					} else {
						g.P(fmt.Sprintf("message.%s.push(reader.%s());", name, g.getWriteFunction(field)))
					}
				}
				g.P(fmt.Sprintf("if (!(message.%s && message.%s.length)) {", name, name))
				g.P(fmt.Sprintf("message.%s = [];", name))
				g.P("}")

				if (*field.Type != google_protobuf.FieldDescriptorProto_TYPE_STRING) && (*field.Type != google_protobuf.FieldDescriptorProto_TYPE_BYTES) {
					g.P("if ((tag & 7) === 2) {")
					g.P("const end2 = reader.uint32() + reader.pos;")
					g.P("while (reader.pos < end2) {")
					assign()
					g.P("}")
					g.P("} else {")
				}
				assign()
				if (*field.Type != google_protobuf.FieldDescriptorProto_TYPE_STRING) && (*field.Type != google_protobuf.FieldDescriptorProto_TYPE_BYTES) {
					g.P("}")
				}
			} else {
				if g.isType64Bit(field) {
					g.P(fmt.Sprintf("const %s = reader.%s();", name, g.getWriteFunction(field)))
					g.P(fmt.Sprintf("message.%s = new protobufjs.util.LongBits(%s.low >>> 0, %s.high >>> 0).toNumber();", name, name, name))
				} else if *field.Type == google_protobuf.FieldDescriptorProto_TYPE_BYTES {
					g.P(fmt.Sprintf("message.%s = new Uint8Array(reader.%s());", name, g.getWriteFunction(field)))
				} else {
					g.P(fmt.Sprintf("message.%s = reader.%s();", name, g.getWriteFunction(field)))
				}
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

func (g *generator) generateMessageClass(message *google_protobuf.DescriptorProto) {
	g.P()
	if g.isMessageDeprecated(message) {
		g.P("/**")
		g.P("* @deprecated")
		g.P("*/")
	}
	name := g.messageName(message)
	g.P(fmt.Sprintf("export class %s implements I%s{", name, name))
	g.generateDecode(message)
	for _, field := range message.Field {
		g.generateField(field, true)
	}
	g.generateConstructor(message)
	g.generateEncode(message)
	g.P("}")
}

func (g *generator) generateEnum(enum *google_protobuf.EnumDescriptorProto) {
	g.P()

	var s []string
	for _, value := range enum.Value {
		s = append(s, "'"+*value.Name+"'")
	}
	g.P(fmt.Sprintf("export type %s = %s", g.enumName(enum), strings.Join(s, " | ")))
}

func (g *generator) methodDeprecated(method *google_protobuf.MethodDescriptorProto) {
	if g.isMethodDeprecated(method) {
		g.P("/**")
		g.P("* @deprecated")
		g.P("*/")
	}
}

func (g *generator) methodDeprecatedLog(method *google_protobuf.MethodDescriptorProto) {
	if g.isMethodDeprecated(method) {
		g.P(fmt.Sprintf("logger.warn('method %s is deprecated');", *method.Name))
	}
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
		requestType := g.getTsTypeFromMessage(method.InputType, false)
		iRequestType := g.getTsTypeFromMessage(method.InputType, true)
		g.P(fmt.Sprintf("requestType: %s,", requestType))
		responseType := g.getTsTypeFromMessage(method.OutputType, false)
		iResponseType := g.getTsTypeFromMessage(method.OutputType, true)
		g.P(fmt.Sprintf("responseType: %s,", responseType))
		g.P(fmt.Sprintf("requestSerialize: (args: %s) => new %s(args).encode().finish() as Buffer,", iRequestType, requestType))
		g.P(fmt.Sprintf("requestDeserialize: (argBuf: Buffer) => %s.decode(argBuf),", requestType))
		g.P(fmt.Sprintf("responseSerialize: (args: %s) => new %s(args).encode().finish() as Buffer,", iResponseType, responseType))
		g.P(fmt.Sprintf("responseDeserialize: (argBuf: Buffer) => %s.decode(argBuf),", responseType))
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
	g.P(fmt.Sprintf("export interface I%sImplementation extends grpcts.Implementations {", gen.CamelCase(*service.Name)))

	for _, method := range service.Method {
		g.methodDeprecated(method)
		inputTypeName := g.getTsTypeFromMessage(method.InputType, true)
		outputTypeName := g.getTsTypeFromMessage(method.OutputType, true)
		if method.ServerStreaming != nil && *method.ServerStreaming && method.ClientStreaming != nil && *method.ClientStreaming {
			g.P(fmt.Sprintf("%s(call: grpcts.grpc.ServerDuplexStream<%s, %s>): void;", g.toLowerFirst(*method.Name), inputTypeName, outputTypeName))
		} else if method.ServerStreaming != nil && *method.ServerStreaming {
			// TODO why there is no type for write stream?
			g.P(fmt.Sprintf("%s(call: grpcts.grpc.ServerWriteableStream<%s>): void;", g.toLowerFirst(*method.Name), inputTypeName))
		} else if method.ClientStreaming != nil && *method.ClientStreaming {
			g.P(fmt.Sprintf("%s(call: grpcts.grpc.ServerReadableStream<%s>): Promise<%s>;", g.toLowerFirst(*method.Name), inputTypeName, outputTypeName))
			g.P(fmt.Sprintf("%s(call: grpcts.grpc.ServerReadableStream<%s>, callback: grpcts.grpc.sendUnaryData<%s>): void;", g.toLowerFirst(*method.Name), inputTypeName, outputTypeName))
		} else {
			g.P(fmt.Sprintf("%s(call: grpcts.grpc.ServerUnaryCall<%s>): Promise<%s>;", g.toLowerFirst(*method.Name), inputTypeName, outputTypeName))
			g.P(fmt.Sprintf("%s(call: grpcts.grpc.ServerUnaryCall<%s>, callback: grpcts.grpc.sendUnaryData<%s>): void;", g.toLowerFirst(*method.Name), inputTypeName, outputTypeName))
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
	g.P("constructor(address: string, credentials?: grpcts.grpc.ChannelCredentials, trace: grpcts.ClientTrace = nodeTrace, options?: object){")
	g.P(fmt.Sprintf("super(%sServiceDefinition, address, credentials, trace, options);", g.toLowerFirst(*service.Name)))
	g.P("}")
	for _, method := range service.Method {
		inputTypeName := g.getTsTypeFromMessage(method.InputType, true)
		g.methodDeprecated(method)
		if method.ServerStreaming != nil && *method.ServerStreaming && method.ClientStreaming != nil && *method.ClientStreaming {
			outputTypeName := g.getTsTypeFromMessage(method.OutputType, true)
			g.P(fmt.Sprintf("public %s(metadata?: grpcts.Metadata) {", g.toLowerFirst(*method.Name)))
			g.methodDeprecatedLog(method)
			g.P(fmt.Sprintf("return super.makeBidiStreamRequest<%s, %s>('%s', metadata);", inputTypeName, outputTypeName, g.toLowerFirst(*method.Name)))
			g.P(fmt.Sprint("}"))
		} else if method.ServerStreaming != nil && *method.ServerStreaming {
			outputTypeName := g.getTsTypeFromMessage(method.OutputType, true)
			g.P(fmt.Sprintf("public %s(req: %s, metadata?: grpcts.Metadata) {", g.toLowerFirst(*method.Name), inputTypeName))
			g.methodDeprecatedLog(method)
			g.P(fmt.Sprintf("return super.makeServerStreamRequest<%s, %s>('%s', req, metadata);", inputTypeName, outputTypeName, g.toLowerFirst(*method.Name)))
			g.P(fmt.Sprint("}"))
		} else if method.ClientStreaming != nil && *method.ClientStreaming {
			outputTypeName := g.getTsTypeFromMessage(method.OutputType, true)
			g.P(fmt.Sprintf("public %s(metadata?: grpcts.Metadata) {", g.toLowerFirst(*method.Name)))
			g.methodDeprecatedLog(method)
			g.P(fmt.Sprintf("return super.makeClientStreamRequest<%s, %s>('%s', metadata);", inputTypeName, outputTypeName, g.toLowerFirst(*method.Name)))
			g.P(fmt.Sprint("};"))
		} else {
			outputTypeName := g.getTsTypeFromMessage(method.OutputType, true)
			g.P(fmt.Sprintf("public %s(req: %s, metadata?: grpcts.Metadata) {", g.toLowerFirst(*method.Name), inputTypeName))
			g.methodDeprecatedLog(method)
			g.P(fmt.Sprintf("return super.makeUnaryRequest<%s, %s>('%s', req, metadata);", inputTypeName, outputTypeName, g.toLowerFirst(*method.Name)))
			g.P(fmt.Sprint("};"))
		}
	}

	g.P(fmt.Sprint("}"))

}

func (g *generator) Make(protoFile *google_protobuf.FileDescriptorProto, protoFiles []*google_protobuf.FileDescriptorProto) (*plugin.CodeGeneratorResponse_File, error) {
	g.protoFile = protoFile

	g.P("// GENERATED CODE -- DO NOT EDIT!")

	g.generateImports(protoFile, protoFiles)
	g.P("import * as protobufjs from 'protobufjs/minimal';")
	g.P("// @ts-ignore ignored as it's generated and it's difficult to predict if logger is needed")
	g.P("import { logger } from '@join-com/gcloud-logger-trace';")
	if len(protoFile.Service) > 0 {
		g.generateGenericImports()
	}
	packageName := protoFile.GetPackage()
	g.generateNamespace(packageName)

	for _, enum := range protoFile.EnumType {
		g.generateEnum(enum)
	}

	for _, message := range protoFile.MessageType {
		g.generateMessageInterface(message)
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
