package generator

/*
 * Runner's methods to collect data from proto files
 */

import (
	"path/filepath"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/join-com/protoc-gen-ts/internal/utils"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/descriptorpb"
)

func (r *Runner) collectData(protoFile *protogen.File) {
	packageName := protoFile.Proto.GetPackage()

	r.packagesByFile[r.currentProtoFilePath] = packageName

	r.collectExportedSymbols(packageName, protoFile.Proto)
	r.collectImportsMap(packageName, protoFile.Proto)
}

func (r *Runner) collectExportedSymbols(packageName string, proto *descriptorpb.FileDescriptorProto) {
	symbolsMap, filesMapExists := r.filesForExportedPackageSymbols[packageName]
	if !filesMapExists || symbolsMap == nil {
		symbolsMap = make(map[string]string)
	}

	// Collect Enums
	for _, enumSpec := range proto.GetEnumType() {
		symbolsMap[strcase.ToCamel(enumSpec.GetName())] = r.currentProtoFilePath
	}

	// Collect Messages
	for _, messageSpec := range proto.GetMessageType() {
		symbolsMap[strcase.ToCamel(messageSpec.GetName())] = r.currentProtoFilePath
	}

	// TODO?: Services

	r.filesForExportedPackageSymbols[packageName] = symbolsMap
}

func (r *Runner) collectImportsMap(packageName string, proto *descriptorpb.FileDescriptorProto) {
	alternativeImportNames, ok := r.alternativeImportNames[r.currentProtoFilePath]
	if !ok || alternativeImportNames == nil {
		alternativeImportNames = make(map[string]string)
	}

	packageRelatedFilesCounter := make(map[string]int)
	for _, importSourcePath := range proto.GetDependency() {
		packageName, validPkgName := r.packagesByFile[importSourcePath]
		if !validPkgName {
			utils.LogError("Unable to retrieve package name for " + importSourcePath)
		}
		packageRelatedFilesCounter[packageName] += 1
	}

	for _, importSourcePath := range proto.GetDependency() {
		// We don't need to validate its existence, as we do it in the previous loop
		packageName := r.packagesByFile[importSourcePath]

		var importName string
		if packageRelatedFilesCounter[packageName] > 1 {
			suffix := strcase.ToCamel(strings.TrimSuffix(filepath.Base(importSourcePath), filepath.Ext(importSourcePath)))
			importName = strcase.ToCamel(packageName) + "_" + suffix
		} else {
			importName = strcase.ToCamel(packageName)
		}
		alternativeImportNames[importSourcePath] = importName
	}

	r.alternativeImportNames[r.currentProtoFilePath] = alternativeImportNames
}
