package generator

import (
	"fmt"

	"github.com/iancoleman/strcase"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/descriptorpb"
)

func (r *Runner) generateTypescriptServiceDefinitions(generatedFileStream *protogen.GeneratedFile, protoFile *protogen.File) {
	for _, serviceSpec := range protoFile.Proto.GetService() {
		r.generateTypescriptServiceImplementationInterface(generatedFileStream, serviceSpec)
		r.generateTypescriptServiceDefinition(generatedFileStream, serviceSpec)
	}
}

func (r *Runner) generateTypescriptServiceImplementationInterface(generatedFileStream *protogen.GeneratedFile, serviceSpec *descriptorpb.ServiceDescriptorProto) {
	r.P(generatedFileStream, "export interface I"+strcase.ToCamel(serviceSpec.GetName())+"ServiceImplementation extends grpc.UntypedServiceImplementation {")
	r.indentLevel += 2

	for _, methodSpec := range serviceSpec.GetMethod() {
		clientStream := methodSpec.GetClientStreaming()
		serverStream := methodSpec.GetServerStreaming()

		inputTypeName := methodSpec.GetInputType()
		inputInterface := r.getEnumOrMessageTypeName(inputTypeName, true)

		outputTypeName := methodSpec.GetOutputType()
		outputInterface := r.getEnumOrMessageTypeName(outputTypeName, true)

		if !clientStream && !serverStream {
			r.P(generatedFileStream, methodSpec.GetName()+": grpc.handleUnaryCall<"+inputInterface+", "+outputInterface+">")
		} else if clientStream {
			r.P(generatedFileStream, methodSpec.GetName()+": grpc.handleClientStreamingCall<"+inputInterface+", "+outputInterface+">")
		} else if serverStream {
			r.P(generatedFileStream, methodSpec.GetName()+": grpc.handleServerStreamingCall<"+inputInterface+", "+outputInterface+">")
		} else {
			r.P(generatedFileStream, methodSpec.GetName()+": grpc.handleBidiStreamingCall<"+inputInterface+", "+outputInterface+">")
		}
	}

	r.indentLevel -= 2
	r.P(generatedFileStream, "}\n")
}

func (r *Runner) generateTypescriptServiceDefinition(generatedFileStream *protogen.GeneratedFile, serviceSpec *descriptorpb.ServiceDescriptorProto) {
	r.P(generatedFileStream, "export const "+strcase.ToLowerCamel(serviceSpec.GetName())+"ServiceDefinition: grpc.ServiceDefinition<I"+strcase.ToCamel(serviceSpec.GetName())+"ServiceImplementation> = {")
	r.indentLevel += 2

	for _, methodSpec := range serviceSpec.GetMethod() {
		r.generateTypescriptServiceDefinitionMethod(generatedFileStream, serviceSpec, methodSpec)
	}

	r.indentLevel -= 2
	r.P(generatedFileStream, "}\n")
}

func (r *Runner) generateTypescriptServiceDefinitionMethod(generatedFileStream *protogen.GeneratedFile, serviceSpec *descriptorpb.ServiceDescriptorProto, methodSpec *descriptorpb.MethodDescriptorProto) {
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
		"requestSerialize: (request: "+inputInterface+") => "+inputClass+".encodePatched(request).finish() as Buffer,",
	)
	r.P(generatedFileStream, "requestDeserialize: "+inputClass+".decodePatched,")

	outputTypeName := methodSpec.GetOutputType()
	outputInterface := r.getEnumOrMessageTypeName(outputTypeName, true)
	outputClass := r.getEnumOrMessageTypeName(outputTypeName, false)
	r.P(generatedFileStream, "responseSerialize: (response: "+outputInterface+") => "+outputClass+".encodePatched(response).finish() as Buffer,")
	r.P(generatedFileStream, "responseDeserialize: "+outputClass+".decodePatched,")

	r.indentLevel -= 2
	r.P(generatedFileStream, "},")
}
