package generator

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/join-com/protoc-gen-ts/internal/utils"
	"google.golang.org/protobuf/compiler/protogen"
)

type Runner struct {
	packagesByFile map[string]string
}

func NewRunner() Runner {
	return Runner{
		packagesByFile: make(map[string]string),
	}
}

func (r *Runner) Run(plugin *protogen.Plugin) error {
	if len(plugin.Request.FileToGenerate) == 0 {
		return errors.New("there are no files to generate")
	}

	// Data collection step
	for _, file := range plugin.Files {
		r.collectData(file)
	}

	// Generation step
	for _, file := range plugin.Files {
		outputPath := fromProtoPathToGeneratedPath(file.Desc.Path()) + ".ts"

		generatedFileStream := plugin.NewGeneratedFile(outputPath, "")
		r.generateTypescriptFile(file, generatedFileStream)
	}

	return nil
}

func (r *Runner) collectData(protoFile *protogen.File) {
	r.packagesByFile[protoFile.Desc.Path()] = protoFile.Proto.GetPackage()
}

func (r *Runner) generateTypescriptFile(protoFile *protogen.File, generatedFileStream *protogen.GeneratedFile) {
	// TODO: Generate comment with version, in order to improve traceability & debugging experience
	generatedFileStream.P("// GENERATED CODE -- DO NOT EDIT!\n")

	r.generateTypescriptImports(protoFile.Proto.Dependency, generatedFileStream)
}

func (r *Runner) generateTypescriptImports(dependencies []string, generatedFileStream *protogen.GeneratedFile) {
	// "Static imports"
	generatedFileStream.P("import * as joinGRPC from '@join-com/grpc'")
	generatedFileStream.P("import * as nodeTrace from '@join-com/node-trace'")
	generatedFileStream.P("")

	// "Dynamic" imports
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

func fromProtoPathToGeneratedPath(protoPath string) string {
	outputDirectory := filepath.Dir(protoPath)
	outputFileName := strcase.ToCamel(strings.TrimSuffix(filepath.Base(protoPath), filepath.Ext(protoPath)))
	generatedPath := fmt.Sprintf("%s/%s", outputDirectory, outputFileName)

	return generatedPath
}
