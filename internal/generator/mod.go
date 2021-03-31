package generator

/*
 * Entry point for the 'generator' package.
 */

import (
	"flag"

	"google.golang.org/protobuf/compiler/protogen"
)

func Generate() {
	var flags flag.FlagSet // It will contain flags passed via protoc

	opts := protogen.Options{
		ParamFunc: flags.Set,
	}
	runner := NewRunner()
	opts.Run(runner.Run)
}
