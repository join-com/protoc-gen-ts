package generator

import (
	"google.golang.org/protobuf/compiler/protogen"
)

func (r *Runner) generateRoot(generatedFileStream *protogen.GeneratedFile) {
	generatedFileStream.P(
		"import { Root } from 'protobufjs'\n\n",
		"export const root = new Root()",
	)
}
