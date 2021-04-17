#!/bin/sh

# Building binary
go build -o ./dist/protoc-gen-tsx

# Stripping symbols to make the binary more lightweight
strip -u -r -S ./dist/protoc-gen-tsx
