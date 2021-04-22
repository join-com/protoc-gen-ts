package generator

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/iancoleman/strcase"
)

func fromProtoPathToGeneratedPath(protoImportPath string, currentSourcePath string) string {

	importDirectory := filepath.Dir(protoImportPath)
	currentDirectory := filepath.Dir(currentSourcePath)

	if !strings.HasPrefix(importDirectory, ".") {
		importDirectory = "./" + importDirectory
	}
	if !strings.HasPrefix(currentDirectory, ".") {
		currentDirectory = "./" + currentDirectory
	}

	importDirectoryParts := strings.Split(importDirectory, "/")
	currentDirectoryParts := strings.Split(currentDirectory, "/")

	minPathLength := min(len(importDirectoryParts), len(currentDirectoryParts))

	for i := 0; i < minPathLength; i++ {
		if importDirectoryParts[0] == currentDirectoryParts[0] {
			importDirectoryParts = importDirectoryParts[1:]
			currentDirectoryParts = currentDirectoryParts[1:]
		} else {
			break
		}
	}

	var outputDirectory string

	numStepsBack := len(currentDirectoryParts)
	if numStepsBack == 0 {
		outputDirectory = "./" + strings.Join(importDirectoryParts, "/")
	} else {
		outputDirectory = strings.Repeat("../", numStepsBack) + strings.Join(importDirectoryParts, "/")
	}
	if !strings.HasSuffix(outputDirectory, "/") {
		outputDirectory += "/"
	}

	outputFileName := strcase.ToCamel(strings.TrimSuffix(filepath.Base(protoImportPath), filepath.Ext(protoImportPath)))
	generatedPath := fmt.Sprintf("%s%s", outputDirectory, outputFileName)

	return generatedPath
}

func getNamespaceFromProtoPackage(protoPackage string) string {
	return strcase.ToCamel(protoPackage)
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
