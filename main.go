package main

import (
	"github.com/antonversal/protoc-gen-ts/generator"
)

func main() {
	g := generator.New()
	g.Generate()
}
