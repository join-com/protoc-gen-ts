package main

import (
	"fmt"
	"log"
	"os"

	"github.com/join-com/protoc-gen-ts/internal/generator"
	"github.com/join-com/protoc-gen-ts/internal/utils"
	legacyGenerator "github.com/join-com/protoc-gen-ts/legacy/generator"
)

func main() {
	log.SetOutput(os.Stderr)
	generatorVersion, isGeneratorConfigured := os.LookupEnv("PROTOC_TS_GENERATOR")
	if !isGeneratorConfigured {
		generatorVersion = "LEGACY"
	}

	if generatorVersion == "LEGACY" {
		g := legacyGenerator.New()
		g.Generate()
	} else if generatorVersion == "V2" {
		generator.Generate()
	} else {
		utils.LogError(
			fmt.Sprintf("invalid generator version (%s)", generatorVersion),
			"acceptable values for the 'PROTOC_TS_GENERATOR' environment variable are:",
			"    - LEGACY",
			"    - V2",
		)
	}
}
