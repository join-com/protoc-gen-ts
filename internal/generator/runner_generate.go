package generator

/*
 * Runner's methods to to generate typescript files
 */

import (
	"fmt"

	"github.com/iancoleman/strcase"
	"github.com/join-com/protoc-gen-ts/internal/utils"
	"google.golang.org/protobuf/compiler/protogen"
)

func (r *Runner) generateTypescriptFile(protoFile *protogen.File, generatedFileStream *protogen.GeneratedFile) {
	// TODO: Generate comment with version, in order to improve traceability & debugging experience
	generatedFileStream.P("// GENERATED CODE -- DO NOT EDIT!\n")

	r.generateTypescriptImports(protoFile.Proto.Dependency, generatedFileStream)
	r.generateTypescriptNamespace(protoFile, generatedFileStream)
}

func (r *Runner) generateTypescriptImports(dependencies []string, generatedFileStream *protogen.GeneratedFile) {
	// Generic imports
	generatedFileStream.P("import * as joinGRPC from '@join-com/grpc'")
	generatedFileStream.P("import * as nodeTrace from '@join-com/node-trace'")
	generatedFileStream.P("")

	// Custom imports
	for _, dependency := range dependencies {
		packageName, validPkgName := r.packagesByFile[dependency]
		if !validPkgName {
			utils.LogError("Unable to retrieve package name for " + dependency)
		}
		generatedFileStream.P(fmt.Sprintf(
			"import { %s } from './%s'",
			strcase.ToCamel(packageName),
			fromProtoPathToGeneratedPath(dependency),
		))
	}
	generatedFileStream.P("")
}

func (r *Runner) generateTypescriptNamespace(protoFile *protogen.File, generatedFileStream *protogen.GeneratedFile) {
	namespaceName := strcase.ToCamel(protoFile.Proto.GetPackage())
	generatedFileStream.P(fmt.Sprintf("namespace %s {\n", namespaceName))
	r.indentLevel += 2

	generatedFileStream.P("\n}\n")
	r.indentLevel -= 2
}
