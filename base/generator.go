package base

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"regexp"
	"strings"

	"github.com/golang/protobuf/proto"
	google_protobuf "github.com/golang/protobuf/protoc-gen-go/descriptor"
	gen "github.com/golang/protobuf/protoc-gen-go/generator"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
)

var camel = regexp.MustCompile("(^[^A-Z0-9]*|[A-Z0-9]*)([A-Z0-9][^A-Z]+|$)")

type Dependency struct {
	protoFileName string
	depFileName   string
}

type fileMaker interface {
	Make(*google_protobuf.FileDescriptorProto, []*google_protobuf.FileDescriptorProto) (*plugin.CodeGeneratorResponse_File, error)
}

type Generator struct {
	*gen.Generator
	indent                  string
	enumNameToObject        map[string]*google_protobuf.EnumDescriptorProto
	reader                  io.Reader
	writer                  io.Writer
	dependencyNameImportMap map[Dependency]string
	messageToFileMap        map[string]string
}

// New creates a new base generator
func New() *Generator {
	return &Generator{
		Generator: gen.New(),
		reader:    os.Stdin,
		writer:    os.Stdout,
	}
}

// P prints the arguments to the generated output.  It handles strings and int32s, plus
// handling indirections because they may be *string, etc.
func (g *Generator) P(str ...interface{}) {
	g.WriteString(g.indent)
	for _, v := range str {
		switch s := v.(type) {
		case string:
			g.WriteString(s)
		case *string:
			g.WriteString(*s)
		case bool:
			fmt.Fprintf(g, "%t", s)
		case *bool:
			fmt.Fprintf(g, "%t", *s)
		case int:
			fmt.Fprintf(g, "%d", s)
		case *int32:
			fmt.Fprintf(g, "%d", *s)
		case *int64:
			fmt.Fprintf(g, "%d", *s)
		case float64:
			fmt.Fprintf(g, "%g", s)
		case *float64:
			fmt.Fprintf(g, "%g", *s)
		default:
			g.Fail(fmt.Sprintf("unknown type in printer: %T", v))
		}
	}
	g.WriteByte('\n')
}

// In Indents the output one tab stop.
func (g *Generator) In() { g.indent += "  " }

// Out unindents the output one tab stop.
func (g *Generator) Out() {
	if len(g.indent) > 0 {
		g.indent = g.indent[2:]
	}
}

// Error reports a problem, including an error, and exits the program.
func (g *Generator) Error(err error, msgs ...string) {
	s := strings.Join(msgs, " ") + ":" + err.Error()
	log.Print("protoc-gen-ts: error:", s)
	os.Exit(1)
}

// Fail reports a problem and exits the program.
func (g *Generator) Fail(msgs ...string) {
	s := strings.Join(msgs, " ")
	log.Print("protoc-gen-ts: error:", s)
	os.Exit(1)
}

// sideEffect calls some methods of the embedded generator from protoc-gen-go
// to make it possible to get object name by type name (via TypeName).
func (g *Generator) sideEffect() {
	g.CommandLineParameters(g.Request.GetParameter())
	g.BuildEnumNameMap(g.Request)
	g.BuildImportsMap(g.Request)
	g.BuildMessageOrEnumToFileMap(g.Request)
	log.Printf("#%v", g.messageToFileMap)
	g.Reset()
}

func (g *Generator) ProtoFileBaseName(name string) string {
	if ext := path.Ext(name); ext == ".proto" || ext == ".protodevel" {
		name = name[:len(name)-len(ext)]
	}
	return name
}

func (g *Generator) generate(maker fileMaker, request *plugin.CodeGeneratorRequest) (*plugin.CodeGeneratorResponse, error) {
	response := new(plugin.CodeGeneratorResponse)
	for _, protoFile := range request.ProtoFile {
		file, err := maker.Make(protoFile, request.ProtoFile)
		if err != nil {
			return response, err
		}
		response.File = append(response.File, file)
	}
	return response, nil
}

func (g *Generator) Generate(maker fileMaker) {
	input, err := ioutil.ReadAll(g.reader)
	if err != nil {
		g.Error(err, "reading input")
	}

	request := g.Request
	if err := proto.Unmarshal(input, request); err != nil {
		g.Error(err, "parsing input proto")
	}

	if len(request.FileToGenerate) == 0 {
		g.Fail("no files to generate")
	}

	g.sideEffect()

	response, err := g.generate(maker, request)
	if err != nil {
		g.Error(err, "failed to generate files from proto")
	}

	output, err := proto.Marshal(response)
	if err != nil {
		g.Error(err, "failed to marshal output proto")
	}
	_, err = g.writer.Write(output)
	if err != nil {
		g.Error(err, "failed to write output proto")
	}
}

func (g *Generator) BuildEnumNameMap(request *plugin.CodeGeneratorRequest) {
	g.enumNameToObject = make(map[string]*google_protobuf.EnumDescriptorProto)
	for _, f := range request.ProtoFile {
		// The names in this loop are defined by the proto world, not us, so the
		// package name may be empty.  If so, the dotted package name of X will
		// be ".X"; otherwise it will be ".pkg.X".
		dottedPkg := "." + f.GetPackage()
		if dottedPkg != "." {
			dottedPkg += "."
		}

		for _, desc := range f.EnumType {
			name := dottedPkg + *desc.Name
			g.enumNameToObject[name] = desc
		}
	}
}

func (g *Generator) BuildImportsMap(request *plugin.CodeGeneratorRequest) {
	var exists = struct{}{}
	g.dependencyNameImportMap = make(map[Dependency]string)
	for _, protoFile := range request.ProtoFile {
		usedImportNames := make(map[string]struct{})
		for _, dependency := range protoFile.Dependency {
			var depProtoFile *google_protobuf.FileDescriptorProto
			for _, extProtoFile := range request.ProtoFile {
				depProtoFileName := *extProtoFile.Name

				if depProtoFileName == dependency {
					depProtoFile = extProtoFile
				}
			}

			packageName := depProtoFile.GetPackage()
			namespaceName := g.namespaceName(packageName)
			importName := namespaceName
			if _, ok := usedImportNames[namespaceName]; ok {
				splits := strings.Split(*depProtoFile.Name, "/")
				fileNameWithExt := splits[len(splits)-1]
				fileNameSplit := strings.Split(fileNameWithExt, ".")
				importName = importName + gen.CamelCase(fileNameSplit[0])
			}
			usedImportNames[namespaceName] = exists
			dep := Dependency{protoFileName: *protoFile.Name, depFileName: *depProtoFile.Name}
			g.dependencyNameImportMap[dep] = importName
		}
	}
}

func (g *Generator) BuildMessageOrEnumToFileMap(request *plugin.CodeGeneratorRequest) {
	g.messageToFileMap = make(map[string]string)
	for _, f := range request.ProtoFile {
		// The names in this loop are defined by the proto world, not us, so the
		// package name may be empty.  If so, the dotted package name of X will
		// be ".X"; otherwise it will be ".pkg.X".
		dottedPkg := "." + f.GetPackage()
		if dottedPkg != "." {
			dottedPkg += "."
		}

		for _, desc := range f.MessageType {
			name := dottedPkg + *desc.Name
			g.messageToFileMap[name] = f.GetName()
		}

		for _, desc := range f.EnumType {
			name := dottedPkg + *desc.Name
			g.messageToFileMap[name] = f.GetName()
		}
	}
}

func (g *Generator) namespaceName(packageName string) string {
	splits := strings.Split(packageName, ".")
	camelCaseName := ""
	for _, name := range splits {
		a := []string{camelCaseName, gen.CamelCase(name)}
		camelCaseName = strings.Join(a, "")
	}

	return camelCaseName
}

func (g *Generator) GetEnumTypeByName(name string) *google_protobuf.EnumDescriptorProto {
	return g.enumNameToObject[name]
}

func (g *Generator) GetImportName(protoFileName string, depFileName string) string {
	dep := Dependency{protoFileName: protoFileName, depFileName: depFileName}
	return g.dependencyNameImportMap[dep]
}

func (g *Generator) GetImportNameForMessage(protoFileName string, message string) string {
	depFileName := g.messageToFileMap[message]
	dep := Dependency{protoFileName: protoFileName, depFileName: depFileName}
	return g.dependencyNameImportMap[dep]
}
