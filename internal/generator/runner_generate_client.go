package generator

import (
	"sort"

	"github.com/iancoleman/strcase"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/descriptorpb"
)

func (r *Runner) generateTypescriptClient(generatedFileStream *protogen.GeneratedFile, protoFile *protogen.File) {
	for _, serviceSpec := range protoFile.Proto.GetService() {
		r.generateTypescriptClientInterface(generatedFileStream, serviceSpec)
		r.generateTypescriptClientClass(generatedFileStream, serviceSpec)
	}
}

func (r *Runner) generateTypescriptClientInterface(generatedFileStream *protogen.GeneratedFile, serviceSpec *descriptorpb.ServiceDescriptorProto) {
	serviceOptions := serviceSpec.GetOptions()
	if serviceOptions != nil {
		if serviceOptions.GetDeprecated() {
			r.P(generatedFileStream, "/**\n  * @deprecated\n */")
		}
	}
	r.P(
		generatedFileStream,
		"export interface I"+strcase.ToCamel(serviceSpec.GetName())+"Client",
		"extends joinGRPC.IExtendedClient<I"+strcase.ToCamel(serviceSpec.GetName())+"ServiceImplementation, '"+r.currentPackage+"."+serviceSpec.GetName()+"'> {",
	)
	r.indentLevel += 2

	methods := serviceSpec.GetMethod()
	sort.Slice(methods, func(i, j int) bool {
		return methods[i].GetName() < methods[j].GetName()
	})

	for _, methodSpec := range methods {
		r.generateTypescriptClientInterfaceMethod(generatedFileStream, serviceSpec, methodSpec)
	}

	r.indentLevel -= 2
	r.P(generatedFileStream, "}\n")
}

func (r *Runner) generateTypescriptClientInterfaceMethod(generatedFileStream *protogen.GeneratedFile, serviceSpec *descriptorpb.ServiceDescriptorProto, methodSpec *descriptorpb.MethodDescriptorProto) {
	methodOptions := methodSpec.GetOptions()
	if methodOptions != nil {
		if methodOptions.GetDeprecated() {
			r.P(generatedFileStream, "/**\n  * @deprecated\n */")
		}
	}

	// Function's Signature
	methodName := strcase.ToLowerCamel(methodSpec.GetName())
	r.P(generatedFileStream, methodName+"(")
	r.indentLevel += 2

	inputTypeName := methodSpec.GetInputType()
	inputInterface := r.getEnumOrMessageTypeName(inputTypeName, true)
	outputTypeName := methodSpec.GetOutputType()
	outputInterface := r.getEnumOrMessageTypeName(outputTypeName, true)

	clientStream := methodSpec.GetClientStreaming()
	serverStrean := methodSpec.GetServerStreaming()

	if !clientStream {
		r.P(generatedFileStream, "request: "+inputInterface+",")
	}

	r.P(generatedFileStream,
		"metadata?: Record<string, string>,",
		"options?: grpc.CallOptions,",
	)

	var returnType string
	if clientStream && serverStrean {
		returnType = "joinGRPC.IBidiStreamRequest<" + inputInterface + ", " + outputInterface + ">"
	} else if !clientStream && !serverStrean {
		returnType = "joinGRPC.IUnaryRequest<" + outputInterface + ">"
	} else if clientStream {
		returnType = "joinGRPC.IClientStreamRequest<" + inputInterface + ", " + outputInterface + ">"
	} else { // if serverStream
		returnType = "joinGRPC.IServerStreamRequest<" + outputInterface + ">"
	}

	r.indentLevel -= 2
	r.P(generatedFileStream, "): "+returnType+"\n")
}

func (r *Runner) generateTypescriptClientClass(generatedFileStream *protogen.GeneratedFile, serviceSpec *descriptorpb.ServiceDescriptorProto) {
	serviceName := serviceSpec.GetName()

	serviceOptions := serviceSpec.GetOptions()
	if serviceOptions != nil {
		if serviceOptions.GetDeprecated() {
			r.P(generatedFileStream, "/**\n  * @deprecated\n */")
		}
	}
	r.P(
		generatedFileStream,
		"export class "+strcase.ToCamel(serviceName)+"Client",
		"extends joinGRPC.Client<I"+strcase.ToCamel(serviceName)+"ServiceImplementation, '"+r.currentPackage+"."+serviceName+"'> ",
		"implements I"+strcase.ToCamel(serviceName)+"Client {",
	)
	r.indentLevel += 2

	r.P(
		generatedFileStream,
		"constructor(",
		"  config: joinGRPC.ISimplifiedClientConfig<I"+strcase.ToCamel(serviceName)+"ServiceImplementation>",
		") {",
	)
	r.indentLevel += 2

	r.P(
		generatedFileStream,
		"super({",
		"  ...config,",
		"  serviceDefinition: "+strcase.ToLowerCamel(serviceName)+"ServiceDefinition,",
		"  credentials: config?.credentials ?? grpc.credentials.createInsecure()",
		"},",
		"'"+r.currentPackage+"."+serviceName+"')",
	)

	r.indentLevel -= 2
	r.P(generatedFileStream, "}\n")

	methods := serviceSpec.GetMethod()
	sort.Slice(methods, func(i, j int) bool {
		return methods[i].GetName() < methods[j].GetName()
	})

	for _, methodSpec := range methods {
		r.generateTypescriptClientMethod(generatedFileStream, serviceSpec, methodSpec)
	}

	r.indentLevel -= 2
	r.P(generatedFileStream, "}\n")
}

func (r *Runner) generateTypescriptClientMethod(generatedFileStream *protogen.GeneratedFile, serviceSpec *descriptorpb.ServiceDescriptorProto, methodSpec *descriptorpb.MethodDescriptorProto) {
	methodOptions := methodSpec.GetOptions()
	isDeprecated := methodOptions != nil && methodOptions.GetDeprecated()

	if isDeprecated {
		r.P(generatedFileStream, "/**\n  * @deprecated\n */")
	}

	// Function's Signature
	methodName := strcase.ToLowerCamel(methodSpec.GetName())
	r.P(generatedFileStream, "public "+methodName+"(")
	r.indentLevel += 2

	inputTypeName := methodSpec.GetInputType()
	inputInterface := r.getEnumOrMessageTypeName(inputTypeName, true)
	outputTypeName := methodSpec.GetOutputType()
	outputInterface := r.getEnumOrMessageTypeName(outputTypeName, true)

	clientStream := methodSpec.GetClientStreaming()
	serverStrean := methodSpec.GetServerStreaming()

	if !clientStream {
		r.P(generatedFileStream, "request: "+inputInterface+",")
	}

	r.P(generatedFileStream,
		"metadata?: Record<string, string>,",
		"options?: grpc.CallOptions,",
	)

	var returnType string
	if clientStream && serverStrean {
		returnType = "joinGRPC.IBidiStreamRequest<" + inputInterface + ", " + outputInterface + ">"
	} else if !clientStream && !serverStrean {
		returnType = "joinGRPC.IUnaryRequest<" + outputInterface + ">"
	} else if clientStream {
		returnType = "joinGRPC.IClientStreamRequest<" + inputInterface + ", " + outputInterface + ">"
	} else { // if serverStream
		returnType = "joinGRPC.IServerStreamRequest<" + outputInterface + ">"
	}

	r.indentLevel -= 2
	r.P(generatedFileStream, "): "+returnType+" {")
	r.indentLevel += 2

	if isDeprecated {
		r.P(generatedFileStream, "this.logger?.warn('using deprecated service method \\'"+strcase.ToCamel(serviceSpec.GetName())+"Client."+methodName+"\\'')")
	}

	methodName = strcase.ToCamel(methodSpec.GetName())
	if clientStream && serverStrean {
		r.P(generatedFileStream, "return this.makeBidiStreamRequest('"+methodName+"', metadata, options)")
	} else if !clientStream && !serverStrean {
		r.P(generatedFileStream, "return this.makeUnaryRequest('"+methodName+"', request, metadata, options)")
	} else if clientStream {
		r.P(generatedFileStream, "return this.makeClientStreamRequest('"+methodName+"', metadata, options)")
	} else { // if serverStream
		r.P(generatedFileStream, "return this.makeServerStreamRequest('"+methodName+"', request, metadata, options)")
	}

	r.indentLevel -= 2
	r.P(generatedFileStream, "}\n")
}
