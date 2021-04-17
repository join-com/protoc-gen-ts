package main

import (
	"log"
	"os"

	"github.com/join-com/protoc-gen-ts/internal/generator"
)

func main() {
	log.SetOutput(os.Stderr)
	generator.Generate()
}
