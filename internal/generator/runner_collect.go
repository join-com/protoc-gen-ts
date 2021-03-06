package generator

/*
 * Runner's methods to collect data from proto files
 */

import (
	"path/filepath"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/join-com/protoc-gen-ts/internal/join_proto"
	"github.com/join-com/protoc-gen-ts/internal/utils"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/descriptorpb"
)

func (r *Runner) collectData(protoFile *protogen.File) {
	r.collectCodeGenerationOptions(protoFile)
	if !r.generateCodeOptions[r.currentProtoFilePath] {
		return
	}

	packageName := protoFile.Proto.GetPackage()

	r.packagesByFile[r.currentProtoFilePath] = packageName

	r.collectExportedSymbols(packageName, protoFile.Proto)
	r.collectImportsMap(packageName, protoFile.Proto)
}

func (r *Runner) collectCodeGenerationOptions(protoFile *protogen.File) {
	generateCode, ok := join_proto.GetBooleanCustomFileOption("typescript_generate_code", protoFile.Proto.Options, r.extensionTypes)
	if !ok {
		generateCode = true
	}
	generateImports, ok := join_proto.GetBooleanCustomFileOption("typescript_generate_imports", protoFile.Proto.Options, r.extensionTypes)
	if !ok {
		generateImports = true
	}
	r.generateCodeOptions[r.currentProtoFilePath] = generateCode && protoFile.Proto.GetSyntax() == "proto3"
	r.importCodeOptions[r.currentProtoFilePath] = generateImports && protoFile.Proto.GetSyntax() == "proto3"
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
	currentPackageName := "." + r.packagesByFile[r.currentProtoFilePath]
	for _, messageSpec := range proto.GetMessageType() {
		symbolsMap[strcase.ToCamel(messageSpec.GetName())] = r.currentProtoFilePath
		r.messageSpecsByFQN[currentPackageName + "." + messageSpec.GetName()] = messageSpec
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
		if !r.importCodeOptions[importSourcePath] {
			continue
		}

		packageName, validPkgName := r.packagesByFile[importSourcePath]
		if !validPkgName {
			utils.LogError("Unable to retrieve package name for " + importSourcePath)
		}
		packageRelatedFilesCounter[packageName] += 1
	}

	for _, importSourcePath := range proto.GetDependency() {
		if !r.importCodeOptions[importSourcePath] {
			continue
		}

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
