package main

import (
	"fmt"
	"os"

	"github.com/join-com/protoc-gen-ts/legacy/generator"
)

func main() {
	generatorVersion, isGeneratorConfigured := os.LookupEnv("PROTOC_TS_GENERATOR")
	if !isGeneratorConfigured {
		generatorVersion = "LEGACY"
	}

	if generatorVersion == "LEGACY" {
		g := generator.New()
		g.Generate()
	} else if generatorVersion == "V2" {
		// TODO
	} else {
		fmt.Fprintf(os.Stderr, "Invalid protoc-gen-ts generator version (%s)\n", generatorVersion)
		os.Stderr.WriteString("Acceptable values for the 'PROTOC_TS_GENERATOR' environment variable are:\n\t- LEGACY\n\t- V2\n\n")
		os.Exit(1)
	}
}
