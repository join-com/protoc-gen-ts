package generator

import (
	"errors"

	"google.golang.org/protobuf/compiler/protogen"
)

type Runner struct {
	indentLevel    int
	packagesByFile map[string]string
}

func NewRunner() Runner {
	return Runner{
		indentLevel:    0,
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
