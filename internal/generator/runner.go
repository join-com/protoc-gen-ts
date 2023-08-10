package generator

/*
 * Code generation orchestrator
 */

import (
	"errors"

	"github.com/join-com/protoc-gen-ts/internal/join_proto"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/descriptorpb"
)

// Sort of a global state for the generator's runner
type Runner struct {
	extensionTypes                 *protoregistry.Types // used to parse custom options
	currentProtoFilePath           string
	currentNamespace               string
	currentPackage                 string
	indentLevel                    int
	packagesByFile                 map[string]string                        // source file path -> package name
	filesForExportedPackageSymbols map[string]map[string]string             // pkg_name -> symbol -> source file path
	alternativeImportNames         map[string]map[string]string             // "current" source file path -> imported source file path -> alternative import name
	generateCodeOptions            map[string]bool                          // source file path -> (should we generate .ts file for it)
	importCodeOptions              map[string]bool                          // source file path -> (should we generate imports on .ts files when imported in .proto files)
	messageSpecsByFQN              map[string]*descriptorpb.DescriptorProto // message fully qualified name -> message "spec"
}

func NewRunner() Runner {
	return Runner{
		extensionTypes:                 new(protoregistry.Types),
		currentProtoFilePath:           "",
		currentNamespace:               "",
		currentPackage:                 "",
		indentLevel:                    0,
		packagesByFile:                 make(map[string]string),
		filesForExportedPackageSymbols: make(map[string]map[string]string),
		alternativeImportNames:         make(map[string]map[string]string),
		generateCodeOptions:            make(map[string]bool),
		importCodeOptions:              make(map[string]bool),
		messageSpecsByFQN:              make(map[string]*descriptorpb.DescriptorProto),
	}
}

func (r *Runner) Run(plugin *protogen.Plugin) error {
	if len(plugin.Request.FileToGenerate) == 0 {
		return errors.New("there are no files to generate")
	}

	// Register all "extensions", needed to support custom options
	for _, file := range plugin.Files {
		if err := join_proto.RegisterAllExtensions(r.extensionTypes, file.Desc); err != nil {
			panic(err)
		}
	}

	// Data collection step (files are listed in topological order)
	for _, file := range plugin.Files {
		r.currentProtoFilePath = file.Desc.Path()
		r.collectData(file)
	}

	// Generation step (files are listed in topological order)
	for _, file := range plugin.Files {
		r.currentProtoFilePath = file.Desc.Path()

		if !r.generateCodeOptions[r.currentProtoFilePath] {
			continue
		}

		outputPath := fromProtoPathToGeneratedPath(file.Desc.Path(), ".") + ".ts"
		generatedFileStream := plugin.NewGeneratedFile(outputPath, "")
		r.generateTypescriptFile(file, generatedFileStream)
	}

	// Generate file with shared protobuf Root instance
	rootFileStream := plugin.NewGeneratedFile("root.ts", "")
	r.generateRoot(rootFileStream)

	return nil
}
