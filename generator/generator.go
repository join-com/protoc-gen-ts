package generator

import (
	"fmt"
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
import * as grpc from 'grpc';
import * as grpcTs from 'grpc-ts';
`))
}

func (g *generator) generateImports(protoFile *google_protobuf.FileDescriptorProto, protoFiles []*google_protobuf.FileDescriptorProto) {
	for _, dependency := range protoFile.Dependency {
		if strings.HasPrefix(dependency, "google/protobuf") {
			continue
		}
		var depProtoFile *google_protobuf.FileDescriptorProto
		for _, extProtoFile := range protoFiles {
			depProtoFileName := *extProtoFile.Name

			if depProtoFileName == dependency {
				depProtoFile = extProtoFile
			}
		}
		packageName := depProtoFile.GetPackage()
		namespaceName := g.namespaceName(packageName)
		fileNames := strings.Split(*depProtoFile.Name, "/")
		path := strings.Repeat("../", len(fileNames)-1)
		path += g.tsFileName(depProtoFile.Name)
		g.P(fmt.Sprintf("import { %s } from '%s';", namespaceName, path))
	}
}

func (g *generator) validateParameters() {
	if _, ok := g.Param["pb_pkg_path"]; !ok {
		g.Fail("parameter `pb_pkg_path` is required (e.g. --ts_out=pb_pkg_path=<pb package path>:<output path>)")
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
	if !strings.HasPrefix(*typeName, ".google.protobuf.") {
		names := strings.Split(*typeName, ".")
		packageName := strings.Join(names[:len(names)-1], ".")
		messageName := g.namespaceName(packageName) + "." + names[len(names)-1]
		return messageName
	}

	names := strings.Split(*typeName, ".")
	name := names[len(names)-1]
	return wellKnownTypes[name]
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

	if g.isFieldOptional(field) || g.isFieldDeprecated(field) || g.isFieldOneOf(field) {
		s += "?"
	}
	s += fmt.Sprintf(": %s", g.getTsFieldType(field))
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
	outputType := g.GetTypeByNamed(*method.OutputType)
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

	outputTypeName := g.getTsFieldType(resultField)
	return outputTypeName
}

func (g *generator) methodDeprecated(method *google_protobuf.MethodDescriptorProto) {
	if g.isMethodDeprecated(method) {
		g.P("/**")
		g.P("* @deprecated")
		g.P("*/")
	}
}

func (g *generator) generateService(service *google_protobuf.ServiceDescriptorProto, packageName string) {
	g.P()
	g.In()
	if g.isServiceDeprecated(service) {
		g.P("/**")
		g.P("* @deprecated")
		g.P("*/")
	}
	g.P(fmt.Sprintf("export class %sService extends grpcTs.Service<%sImplementation> {", gen.CamelCase(*service.Name), gen.CamelCase(*service.Name)))
	g.In()
	g.P(fmt.Sprintf("constructor(protoPath: string, implementations: %sImplementation) {", gen.CamelCase(*service.Name)))
	g.In()
	g.P(fmt.Sprintf("super(protoPath, '%s', '%s', implementations);", packageName, *service.Name))
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

func (g *generator) generateClient(service *google_protobuf.ServiceDescriptorProto, packageName string) {
	g.P()
	g.In()
	if g.isServiceDeprecated(service) {
		g.P("/**")
		g.P("* @deprecated")
		g.P("*/")
	}

	g.P(fmt.Sprintf("export class %sClient extends grpcTs.Client {", gen.CamelCase(*service.Name)))
	g.In()
	g.P("constructor(protoPath: string, host: string) {")
	g.In()
	g.P(fmt.Sprintf("super(protoPath, '%s', '%s', host);", packageName, *service.Name))
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
	g.validateParameters()
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
	}

	for _, service := range protoFile.Service {
		g.generateImplementation(service)
		g.generateService(service, *protoFile.Package)
	}

	for _, service := range protoFile.Service {
		g.generateClient(service, *protoFile.Package)
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
