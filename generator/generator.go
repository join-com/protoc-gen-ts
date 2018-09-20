package generator

import (
	"fmt"
	"log"
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
	g.P(fmt.Sprint(`
import * as protobufjs from 'protobufjs/minimal';
import * as grpc from 'grpc';
import * as grpcTs from '@join-com/grpc-ts';
import * as path from 'path';
`))
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
		path := strings.Repeat("../", len(fileNames)-1)
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

func (g *generator) generateField(field *google_protobuf.FieldDescriptorProto) {
	if g.isFieldDeprecated(field) {
		g.P("/** @deprecated */")
	}
	s := *field.JsonName
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
	g.In()

	if g.isMessageDeprecated(message) {
		g.P("/**")
		g.P("* @deprecated")
		g.P("*/")
	}
	g.P(fmt.Sprintf("export interface %s {", g.messageName(message)))
	g.In()
	for _, field := range message.Field {
		g.generateField(field)
	}
	g.Out()
	g.P("}")
	g.Out()
}

func (g *generator) generatePrivateField(field *google_protobuf.FieldDescriptorProto) {
	if g.isFieldDeprecated(field) {
		g.P("/** @deprecated */")
	}
	s := "private _" + *field.JsonName
	s += "?"
	if *field.Type == google_protobuf.FieldDescriptorProto_TYPE_ENUM {
		s += ": number"
	} else {
		s += fmt.Sprintf(": %s", g.getTsFieldType(field))
	}
	if g.isFieldRepeated(field) {
		s += "[]"
	}
	s += ";"
	g.P(s)
}

func (g *generator) generateConstructor(message *google_protobuf.DescriptorProto) {
	name := g.messageName(message)
	g.P(fmt.Sprintf("constructor(attrs?: %s){", name))
	g.In()
	g.P("Object.assign(this, attrs)")
	g.Out()
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
	g.In()
	for _, field := range message.Field {
		name := *field.JsonName
		g.P(fmt.Sprintf("if (this._%s != null) {", name))
		g.In()

		if *field.Type == google_protobuf.FieldDescriptorProto_TYPE_MESSAGE {
			if g.isFieldRepeated(field) {
				g.P(fmt.Sprintf("for (const value of this._%s) {", name))
				g.In()
				g.P(fmt.Sprintf("const msg = new %sMsg(value)", g.getTsTypeFromMessage(field.TypeName)))
				g.P(fmt.Sprintf("msg.encode(writer.uint32(%d).fork()).ldelim();", g.getFieldIndex(field)))
				g.Out()
				g.P("}")
			} else {
				g.P(fmt.Sprintf("const msg = new %sMsg(this._%s)", g.getTsTypeFromMessage(field.TypeName), name))
				g.P(fmt.Sprintf("msg.encode(writer.uint32(%d).fork()).ldelim();", g.getFieldIndex(field)))
			}
		} else {
			if g.isFieldRepeated(field) {
				g.P(fmt.Sprintf("for (const value of this._%s) {", name))
				g.In()
				g.P(fmt.Sprintf("writer.uint32(%d).%s(value);", g.getFieldIndex(field), g.getWriteFunction(field)))
				g.Out()
				g.P("}")
			} else {
				g.P(fmt.Sprintf("writer.uint32(%d).%s(this._%s);", g.getFieldIndex(field), g.getWriteFunction(field), name))
			}

		}
		g.Out()
		g.P("}")
	}
	g.P("return writer")
	g.Out()
	g.P("}")
}

func (g *generator) generateDecode(message *google_protobuf.DescriptorProto) {
	g.P("public decode(inReader: Uint8Array | protobufjs.Reader, length?: number){")
	g.In()
	g.P("const reader = !(inReader instanceof protobufjs.Reader)")
	g.In()
	g.P("? protobufjs.Reader.create(inReader)")
	g.P(": inReader")
	g.Out()
	g.P("const end = length === undefined ? reader.len : reader.pos + length")
	g.P("const message = new Request()")
	g.P("while (reader.pos < end) {")
	g.In()
	g.P("const tag = reader.uint32()")
	g.P("switch (tag >>> 3) {")
	g.In()

	g.Out()
	g.P("}")
	g.Out()
	g.P("}")
	g.Out()
	g.P("}")
}

func (g *generator) generateGettersSetters(message *google_protobuf.DescriptorProto) {
	for _, field := range message.Field {
		name := *field.JsonName
		g.P(fmt.Sprintf("public get %s() {", name))
		g.In()
		if g.isFieldDeprecated(field) {
			g.P(fmt.Sprintf("console.warn('%s is deprecated')", name))
		}
		if *field.Type == google_protobuf.FieldDescriptorProto_TYPE_ENUM {
			enum := g.GetEnumTypeByName(*field.TypeName)
			g.P(fmt.Sprintf("if (!this._%s) {", name))
			g.In()
			g.P("return")
			g.Out()
			g.P("}")
			if g.isFieldRepeated(field) {
				g.P(fmt.Sprintf("return this._%s.map((val) => {", name))
				g.In()
				g.P("switch (val) {")
				typeName := g.getTsFieldType(field)
				for _, value := range enum.Value {
					g.P(fmt.Sprintf("case %d:", *value.Number))
					g.In()
					g.P(fmt.Sprintf("return %s.%s", typeName, *value.Name))
					g.Out()
				}
				g.P("default:")
				g.In()
				g.P(fmt.Sprintf("throw new Error('Undefined value of enum %s for field %s of message %s');", typeName, field.GetJsonName(), message.GetName()))
				g.Out()
				g.P("}")
				g.Out()
				g.P("})")
			} else {
				g.P(fmt.Sprintf("switch (this._%s) {", name))
				typeName := g.getTsFieldType(field)
				for _, value := range enum.Value {
					g.P(fmt.Sprintf("case %d:", *value.Number))
					g.In()
					g.P(fmt.Sprintf("return %s.%s", typeName, *value.Name))
					g.Out()
				}
				g.P("default:")
				g.In()
				g.P("return")
				g.Out()
				g.P("}")
			}
		} else if field.GetTypeName() == ".google.protobuf.Timestamp" {
			if g.isFieldRepeated(field) {
				g.P(fmt.Sprintf("return this._%s && this._%s.map(v => new Date(((v.seconds || 0) * 1000) + ((v.nanos || 0) / 1000000)))", name, name))
			} else {
				g.P(fmt.Sprintf("return this._%s && new Date(((this._%s.seconds || 0) * 1000) + ((this._%s.nanos || 0) / 1000000))", name, name, name))
			}
		} else {
			g.P(fmt.Sprintf("return this._%s", name))
		}

		g.Out()
		g.P("}")

		g.P(fmt.Sprintf("public set %s(val) {", name))
		g.In()
		if g.isFieldDeprecated(field) {
			g.P(fmt.Sprintf("console.warn('%s is deprecated')", name))
		}

		if *field.Type == google_protobuf.FieldDescriptorProto_TYPE_MESSAGE {
			if field.GetTypeName() == ".google.protobuf.Timestamp" {
				if g.isFieldRepeated(field) {
					g.P(fmt.Sprintf("this._%s = val && val.map(v => ({ seconds: Math.floor(v.getTime() / 1000), nanos: v.getMilliseconds() * 1000000 }))", name))
				} else {
					g.P(fmt.Sprintf("this._%s = val && { seconds: Math.floor(val.getTime() / 1000), nanos: val.getMilliseconds() * 1000000 }", name))
				}
			} else {
				if g.isFieldRepeated(field) {
					g.P(fmt.Sprintf("this._%s = val && val.map(v => new %sMsg(v))", name, g.getTsTypeFromMessage(field.TypeName)))
				} else {
					g.P(fmt.Sprintf("this._%s = new %sMsg(val)", name, g.getTsTypeFromMessage(field.TypeName)))
				}
			}
		} else if *field.Type == google_protobuf.FieldDescriptorProto_TYPE_ENUM {
			enum := g.GetEnumTypeByName(*field.TypeName)
			if g.isFieldRepeated(field) {
				g.P("if (!val) {")
				g.In()
				g.P("return")
				g.Out()
				g.P("}")
				g.P(fmt.Sprintf("this._%s = val.map((value) => {", name))
				g.In()
				g.P("switch (value) {")
				typeName := g.getTsFieldType(field)
				for _, value := range enum.Value {
					g.P(fmt.Sprintf("case %s.%s:", typeName, *value.Name))
					g.In()
					g.P(fmt.Sprintf("return %d", *value.Number))
					g.Out()
				}
				g.P("default:")
				g.In()
				g.P(fmt.Sprintf("throw new Error('Undefined value of enum %s for field %s of message %s');", typeName, field.GetJsonName(), message.GetName()))
				g.Out()
				g.P("}")
				g.Out()
				g.P("})")
			} else {
				g.P("switch (val) {")
				typeName := g.getTsFieldType(field)
				for _, value := range enum.Value {
					g.P(fmt.Sprintf("case %s.%s:", typeName, *value.Name))
					g.In()
					g.P(fmt.Sprintf("this._%s = %d", name, *value.Number))
					g.P("break;")
					g.Out()
				}
				g.P("default:")
				g.In()
				g.P(fmt.Sprintf("this._%s = undefined", name))
				g.Out()
				g.P("}")
			}
		} else {
			g.P(fmt.Sprintf("this._%s = val", name))
		}

		g.Out()
		g.P("}")
	}
}

func (g *generator) generateMessageClass(message *google_protobuf.DescriptorProto) {
	g.P()
	g.In()

	if g.isMessageDeprecated(message) {
		g.P("/**")
		g.P("* @deprecated")
		g.P("*/")
	}
	name := g.messageName(message)
	g.P(fmt.Sprintf("export class %sMsg implements %s{", name, name))
	g.In()
	for _, field := range message.Field {
		g.generatePrivateField(field)
	}
	g.generateConstructor(message)
	g.generateGettersSetters(message)
	g.generateEncode(message)
	g.Out()
	g.P("}")
	g.Out()
}

func (g *generator) generateEnum(enum *google_protobuf.EnumDescriptorProto) {
	g.P()
	g.In()
	g.P(fmt.Sprintf("export enum %s {", g.enumName(enum)))
	g.In()
	for _, value := range enum.Value {
		g.P(fmt.Sprintf("%s = '%s',", *value.Name, *value.Name))
	}
	g.Out()
	g.P(fmt.Sprint("}"))
	g.Out()
}

func (g *generator) methodOutputType(method *google_protobuf.MethodDescriptorProto) string {
	outputType := g.GetMessageTypeByName(*method.OutputType)
	var resultField *google_protobuf.FieldDescriptorProto
	var errorField *google_protobuf.FieldDescriptorProto
	for _, field := range outputType.Field {
		if *field.Name == "result" {
			resultField = field
		}
		if *field.Name == "error" {
			errorField = field
		}
	}
	if resultField == nil {
		g.Fail(fmt.Sprintf("Response massage %s must have result field", *method.OutputType))
	}

	if errorField == nil {
		g.Fail(fmt.Sprintf("Response massage %s must have error field", *method.OutputType))
	}

	var outputTypeName string

	outputTypeName = g.getTsFieldType(resultField)
	if g.isFieldRepeated(resultField) {
		outputTypeName = outputTypeName + "[]"
	}
	return outputTypeName
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

func (g *generator) generateService(service *google_protobuf.ServiceDescriptorProto, packageName string, protoFileName string) {
	g.P()
	g.In()
	if g.isServiceDeprecated(service) {
		g.P("/**")
		g.P("* @deprecated")
		g.P("*/")
	}
	g.P(fmt.Sprintf("export class %sService extends grpcTs.Service<%sImplementation> {", gen.CamelCase(*service.Name), gen.CamelCase(*service.Name)))
	g.In()
	g.P(fmt.Sprintf("constructor(implementations: %sImplementation, errorHandler?: grpcTs.ErrorHandler ) {", gen.CamelCase(*service.Name)))
	g.In()
	g.P(fmt.Sprintf("const protoPath = '%s';", protoFileName))
	g.P(fmt.Sprintf("const includeDirs = [path.join(__dirname, '%s')];", g.protoPathRelative(protoFileName)))
	g.P(fmt.Sprintf("super(protoPath, includeDirs, '%s', '%s', implementations, errorHandler);", packageName, *service.Name))
	g.Out()
	g.P(fmt.Sprint("}"))
	g.Out()
	g.P(fmt.Sprint("}"))
	g.Out()
}

func (g *generator) generateImplementation(service *google_protobuf.ServiceDescriptorProto) {
	g.P()
	g.In()
	if g.isServiceDeprecated(service) {
		g.P("/**")
		g.P("* @deprecated")
		g.P("*/")
	}
	g.P(fmt.Sprintf("export interface %sImplementation {", gen.CamelCase(*service.Name)))
	g.In()
	for _, method := range service.Method {
		g.methodDeprecated(method)
		inputTypeName := g.getTsTypeFromMessage(method.InputType)
		if method.ServerStreaming != nil && *method.ServerStreaming && method.ClientStreaming != nil && *method.ClientStreaming {
			outputTypeName := g.getTsTypeFromMessage(method.OutputType)
			g.P(fmt.Sprintf("%s(duplexStream: grpc.ServerDuplexStream<%s, %s>): void;", g.toLowerFirst(*method.Name), inputTypeName, outputTypeName))
		} else if method.ServerStreaming != nil && *method.ServerStreaming {
			outputTypeName := g.getTsTypeFromMessage(method.OutputType)
			g.P(fmt.Sprintf("%s(req: %s, stream: grpc.ServerWriteableStream<%s>): void;", g.toLowerFirst(*method.Name), inputTypeName, outputTypeName))
		} else if method.ClientStreaming != nil && *method.ClientStreaming {
			outputTypeName := g.methodOutputType(method)
			g.P(fmt.Sprintf("%s(stream: grpc.ServerReadableStream<%s>): Promise<%s>;", g.toLowerFirst(*method.Name), inputTypeName, outputTypeName))
		} else {
			outputTypeName := g.methodOutputType(method)
			g.P(fmt.Sprintf("%s(req: %s): Promise<%s>;", g.toLowerFirst(*method.Name), inputTypeName, outputTypeName))
		}
	}
	g.Out()
	g.P(fmt.Sprint("}"))
	g.Out()
}

func (g *generator) generateClient(service *google_protobuf.ServiceDescriptorProto, packageName string, protoFileName string) {
	g.P()
	g.In()
	if g.isServiceDeprecated(service) {
		g.P("/**")
		g.P("* @deprecated")
		g.P("*/")
	}

	g.P(fmt.Sprintf("export class %sClient extends grpcTs.Client {", gen.CamelCase(*service.Name)))
	g.In()
	g.P("constructor(host: string, credentials: grpc.ChannelCredentials) {")
	g.In()
	g.P(fmt.Sprintf("const protoPath = '%s';", protoFileName))
	g.P(fmt.Sprintf("const includeDirs = [path.join(__dirname, '%s')];", g.protoPathRelative(protoFileName)))
	g.P(fmt.Sprintf("super(protoPath, includeDirs, '%s', '%s', host, credentials);", packageName, *service.Name))
	g.Out()
	g.P("}")
	for _, method := range service.Method {
		inputTypeName := g.getTsTypeFromMessage(method.InputType)
		g.methodDeprecated(method)
		if method.ServerStreaming != nil && *method.ServerStreaming && method.ClientStreaming != nil && *method.ClientStreaming {
			outputTypeName := g.getTsTypeFromMessage(method.OutputType)
			g.P(fmt.Sprintf("public %s(): grpc.ClientDuplexStream<%s, %s> {", g.toLowerFirst(*method.Name), inputTypeName, outputTypeName))
			g.In()
			g.P(fmt.Sprintf("return super.makeBidiStreamRequest('%s');", g.toLowerFirst(*method.Name)))
			g.Out()
			g.P(fmt.Sprint("}"))
		} else if method.ServerStreaming != nil && *method.ServerStreaming {
			outputTypeName := g.getTsTypeFromMessage(method.OutputType)
			g.P(fmt.Sprintf("public %s(req: %s): grpc.ClientReadableStream<%s> {", g.toLowerFirst(*method.Name), inputTypeName, outputTypeName))
			g.In()
			g.P(fmt.Sprintf("return super.makeServerStreamRequest('%s', req);", g.toLowerFirst(*method.Name)))
			g.Out()
			g.P(fmt.Sprint("}"))
		} else if method.ClientStreaming != nil && *method.ClientStreaming {
			outputTypeName := g.getTsTypeFromMessage(method.OutputType)
			g.P(fmt.Sprintf("public %s(callback: grpcTs.Callback<%s>): grpc.ClientWritableStream<%s> {", g.toLowerFirst(*method.Name), outputTypeName, inputTypeName))
			g.In()
			g.P(fmt.Sprintf("return super.makeClientStreamRequest('%s', callback);", g.toLowerFirst(*method.Name)))
			g.Out()
			g.P(fmt.Sprint("}"))
		} else {
			outputTypeName := g.methodOutputType(method)
			g.P(fmt.Sprintf("public %s(req: %s): Promise<%s> {", g.toLowerFirst(*method.Name), inputTypeName, outputTypeName))
			g.In()
			g.P(fmt.Sprintf("return super.makeUnaryRequest('%s', req);", g.toLowerFirst(*method.Name)))
			g.Out()
			g.P(fmt.Sprint("}"))
		}
	}
	g.Out()
	g.P(fmt.Sprint("}"))
	g.Out()
}

func (g *generator) Make(protoFile *google_protobuf.FileDescriptorProto, protoFiles []*google_protobuf.FileDescriptorProto) (*plugin.CodeGeneratorResponse_File, error) {
	g.protoFile = protoFile
	g.validateParameters()
	log.Print(*protoFile.Name)

	g.P("// GENERATED CODE -- DO NOT EDIT!")

	g.generateImports(protoFile, protoFiles)
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
		g.generateImplementation(service)
		g.generateService(service, *protoFile.Package, *protoFile.Name)
	}

	for _, service := range protoFile.Service {
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
