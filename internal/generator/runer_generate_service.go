package generator

import (
	"fmt"

	"github.com/iancoleman/strcase"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/descriptorpb"
)

func (r *Runner) generateTypescriptServiceDefinitions(generatedFileStream *protogen.GeneratedFile, protoFile *protogen.File) {
	for _, serviceSpec := range protoFile.Proto.GetService() {
		r.generateTypescriptServiceDefinition(generatedFileStream, serviceSpec)
	}
}

func (r *Runner) generateTypescriptServiceDefinition(generatedFileStream *protogen.GeneratedFile, serviceSpec *descriptorpb.ServiceDescriptorProto) {
	r.P(generatedFileStream, "export const "+strcase.ToLowerCamel(serviceSpec.GetName())+"ServiceDefinition = {")
	r.indentLevel += 2

	for _, methodSpec := range serviceSpec.GetMethod() {
		r.generateTypescriptServiceMethod(generatedFileStream, serviceSpec, methodSpec)
	}

	r.indentLevel -= 2
	r.P(generatedFileStream, "}\n")
}

func (r *Runner) generateTypescriptServiceMethod(generatedFileStream *protogen.GeneratedFile, serviceSpec *descriptorpb.ServiceDescriptorProto, methodSpec *descriptorpb.MethodDescriptorProto) {
	methodName := methodSpec.GetName()

	r.P(generatedFileStream, methodName+": {")
	r.indentLevel += 2

	r.P(generatedFileStream, "path: '/"+r.currentPackage+"."+serviceSpec.GetName()+"/"+methodName+"',")

	r.P(generatedFileStream, "requestStream: "+fmt.Sprintf("%t", methodSpec.GetClientStreaming())+",")
	r.P(generatedFileStream, "responseStream: "+fmt.Sprintf("%t", methodSpec.GetServerStreaming())+",")

	inputTypeName := methodSpec.GetInputType()
	inputInterface := r.getEnumOrMessageTypeName(inputTypeName, true)
	inputClass := r.getEnumOrMessageTypeName(inputTypeName, false)
	r.P(
		generatedFileStream,
		"requestSerialize: (request: "+inputInterface+") => "+inputClass+".encodePatched(request).finish(),",
	)
	r.P(generatedFileStream, "requestDeserialize: "+inputClass+".decodePatched,")

	outputTypeName := methodSpec.GetOutputType()
	outputInterface := r.getEnumOrMessageTypeName(outputTypeName, true)
	outputClass := r.getEnumOrMessageTypeName(outputTypeName, false)
	r.P(generatedFileStream, "responseSerialize: (response: "+outputInterface+") => "+outputClass+".encodePatched(response).finish(),")
	r.P(generatedFileStream, "responseDeserialize: "+outputClass+".decodePatched,")

	r.indentLevel -= 2
	r.P(generatedFileStream, "},")
}
