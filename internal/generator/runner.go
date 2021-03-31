package generator

/*
 * Code generation orchestrator
 */

import (
	"errors"

	"google.golang.org/protobuf/compiler/protogen"
)

// Sort of a global state for the generator's runner
type Runner struct {
	currentProtoFilePath           string
	currentNamespace               string
	indentLevel                    int
	packagesByFile                 map[string]string
	filesForExportedPackageSymbols map[string]map[string]string // pkg_name -> symbol -> source file path
	alternativeImportNames         map[string]map[string]string // "current" source file path -> imported source file path -> alternative import name
}

func NewRunner() Runner {
	return Runner{
		currentProtoFilePath:           "",
		currentNamespace:               "",
		indentLevel:                    0,
		packagesByFile:                 make(map[string]string),
		filesForExportedPackageSymbols: make(map[string]map[string]string),
		alternativeImportNames:         make(map[string]map[string]string),
	}
}

func (r *Runner) Run(plugin *protogen.Plugin) error {
	if len(plugin.Request.FileToGenerate) == 0 {
		return errors.New("there are no files to generate")
	}

	// Data collection step (files are listed in topological order)
	for _, file := range plugin.Files {
		r.currentProtoFilePath = file.Desc.Path()
		r.collectData(file)
	}

	// Generation step (files are listed in topological order)
	for _, file := range plugin.Files {
		r.currentProtoFilePath = file.Desc.Path()

		outputPath := fromProtoPathToGeneratedPath(file.Desc.Path()) + ".ts"
		generatedFileStream := plugin.NewGeneratedFile(outputPath, "")
		r.generateTypescriptFile(file, generatedFileStream)
	}

	return nil
}
