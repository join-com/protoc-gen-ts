package generator

/*
 * Runner's methods to to generate typescript files
 */

import (
	"fmt"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/join-com/protoc-gen-ts/internal/join_proto"
	"github.com/join-com/protoc-gen-ts/internal/utils"
	"google.golang.org/protobuf/compiler/protogen"
)

func (r *Runner) generateTypescriptFile(protoFile *protogen.File, generatedFileStream *protogen.GeneratedFile) {
	// TODO: Generate comment with version, in order to improve traceability & debugging experience
	generatedFileStream.P("// GENERATED CODE -- DO NOT EDIT!\n")

	r.generateTypescriptImports(protoFile.Desc.Path(), protoFile.Proto.GetDependency(), generatedFileStream)
	r.generateTypescriptNamespace(protoFile, generatedFileStream)
}

func (r *Runner) generateTypescriptImports(currentSourcePath string, importSourcePaths []string, generatedFileStream *protogen.GeneratedFile) {
	// Generic imports
	generatedFileStream.P("// import * as joinGRPC from '@join-com/grpc'")        // TODO: Remove comment when import is used
	generatedFileStream.P("// import * as nodeTrace from '@join-com/node-trace'") // TODO: Remove comment when import is used
	generatedFileStream.P("")

	// Custom imports
	for _, importSourcePath := range importSourcePaths {
		if !r.importCodeOptions[importSourcePath] {
			continue
		}

		generatedFileStream.P(fmt.Sprintf(
			"import { %s } from './%s'",
			r.generateImportName(currentSourcePath, importSourcePath),
			fromProtoPathToGeneratedPath(importSourcePath),
		))
	}
	generatedFileStream.P("")
}

func (r *Runner) generateImportName(currentSourcePath string, importSourcePath string) string {
	packageName, validPkgName := r.packagesByFile[importSourcePath]
	if !validPkgName {
		utils.LogError("Unable to retrieve package name for " + importSourcePath)
	}

	pkgNamespace := strcase.ToCamel(packageName)

	alternativeImportNames, validImportsMap := r.alternativeImportNames[currentSourcePath]
	if !validImportsMap || alternativeImportNames == nil {
		utils.LogError("Unable to retrieve alternative imports map for " + currentSourcePath)
	}

	alternativeImportName, validAlternativeName := alternativeImportNames[importSourcePath]
	if !validAlternativeName {
		utils.LogError("Unable to retrieve alternative import name for " + importSourcePath + " on " + currentSourcePath)
	}

	if pkgNamespace == alternativeImportName {
		return pkgNamespace
	} else {
		return pkgNamespace + " as " + alternativeImportName
	}
}

func (r *Runner) generateTypescriptNamespace(protoFile *protogen.File, generatedFileStream *protogen.GeneratedFile) {
	namespace := getNamespaceFromProtoPackage(protoFile.Proto.GetPackage())
	r.P(generatedFileStream, fmt.Sprintf("export namespace %s {\n", namespace))
	r.indentLevel += 2
	r.currentNamespace = namespace

	r.generateTypescriptEnums(protoFile, generatedFileStream)
	r.generateTypescriptMessageInterfaces(protoFile, generatedFileStream)

	r.currentNamespace = ""
	r.indentLevel -= 2
	r.P(generatedFileStream, "\n}")
}

func (r *Runner) generateTypescriptEnums(protoFile *protogen.File, generatedFileStream *protogen.GeneratedFile) {
	for _, enumSpec := range protoFile.Proto.GetEnumType() {
		options := enumSpec.GetOptions()
		if options != nil {
			if options.Deprecated != nil && *options.Deprecated {
				r.P(generatedFileStream, "/**\n  * @deprecated\n */")
			}
		}

		var values []string
		for _, enumValue := range enumSpec.GetValue() {
			values = append(values, "'"+enumValue.GetName()+"'")
		}

		enumName := strcase.ToCamel(enumSpec.GetName())
		enumValues := strings.Join(values, " | ")

		r.P(generatedFileStream, "export type "+enumName+" = "+enumValues)
	}
	r.P(generatedFileStream, "")
}

func (r *Runner) generateTypescriptMessageInterfaces(protoFile *protogen.File, generatedFileStream *protogen.GeneratedFile) {
	for _, messageSpec := range protoFile.Proto.GetMessageType() {
		messageOptions := messageSpec.GetOptions()
		if messageOptions != nil {
			if messageOptions.GetDeprecated() {
				r.P(generatedFileStream, "/**\n  * @deprecated\n */")
			}
		}

		interfaceName := "export interface I" + strcase.ToCamel(messageSpec.GetName()) + " {"

		r.P(generatedFileStream, interfaceName)
		r.indentLevel += 2

		for _, fieldSpec := range messageSpec.GetField() {
			// TODO: Add support for required fields (via proto options)
			fieldOptions := fieldSpec.GetOptions()
			if fieldOptions != nil {
				if messageOptions.GetDeprecated() {
					r.P(generatedFileStream, "/**\n  * @deprecated\n */")
				}
			}

			separator := "?: "
			required, found := join_proto.GetBooleanCustomFieldOption("typescript_required", fieldSpec.GetOptions(), r.extensionTypes)
			if found && required {
				separator = ": "
			}

			r.P(generatedFileStream, fieldSpec.GetJsonName()+separator+r.getMessageFieldType(fieldSpec))
		}

		r.indentLevel -= 2
		r.P(generatedFileStream, "}\n")
	}
}
