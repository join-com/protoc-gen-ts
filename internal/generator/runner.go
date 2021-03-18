package generator

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/iancoleman/strcase"
	"google.golang.org/protobuf/compiler/protogen"
)

type Runner struct {
	// To keep state
}

func NewRunner() Runner {
	return Runner{}
}

func (r *Runner) Run(plugin *protogen.Plugin) error {
	if len(plugin.Request.FileToGenerate) == 0 {
		return errors.New("there are no files to generate")
	}

	for _, file := range plugin.Files {
		outputDirectory := filepath.Dir(file.Desc.Path())
		outputFileName := strcase.ToCamel(strings.TrimSuffix(filepath.Base(file.Desc.Path()), filepath.Ext(file.Desc.Path()))) + ".ts"
		outputPath := fmt.Sprintf("%s/%s", outputDirectory, outputFileName)

		generatedFileStream := plugin.NewGeneratedFile(outputPath, "")
		r.generateTypescriptFile(file, generatedFileStream)
	}

	return nil
}

func (r *Runner) generateTypescriptFile(protoFile *protogen.File, generatedFileStream *protogen.GeneratedFile) {
	generatedFileStream.P("// Hello world")
}
