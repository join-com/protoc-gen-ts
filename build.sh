#!/bin/sh

# Building binary
env GOOS=darwin GOARCH=amd64 go build -o ./dist/protoc-gen-tsx.darwin.amd64
env GOOS=darwin GOARCH=arm64 go build -o ./dist/protoc-gen-tsx.darwin.arm64
env GOOS=linux GOARCH=amd64 go build -o ./dist/protoc-gen-tsx.linux.amd64
env GOOS=linux GOARCH=arm64 go build -o ./dist/protoc-gen-tsx.linux.arm64

if [ "${CI}" != "true" ]; then
    # Stripping symbols to make the binary more lightweight
    strip -u -r -S ./dist/protoc-gen-tsx.darwin.amd64
fi
