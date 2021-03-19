package generator

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/iancoleman/strcase"
)

func fromProtoPathToGeneratedPath(protoPath string) string {
	outputDirectory := filepath.Dir(protoPath)
	outputFileName := strcase.ToCamel(strings.TrimSuffix(filepath.Base(protoPath), filepath.Ext(protoPath)))
	generatedPath := fmt.Sprintf("%s/%s", outputDirectory, outputFileName)

	return generatedPath
}
