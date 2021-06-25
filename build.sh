#!/bin/sh

# Building binary
env GOOS=darwin GOARCH=amd64 go build -o ./dist/protoc-gen-tsx.darwin.amd64
env GOOS=darwin GOARCH=arm64 go build -o ./dist/protoc-gen-tsx.darwin.arm64
env GOOS=linux GOARCH=amd64 go build -o ./dist/protoc-gen-tsx.linux.amd64
env GOOS=linux GOARCH=arm64 go build -o ./dist/protoc-gen-tsx.linux.arm64

rm -f ./dist/protoc-gen-tsx;

if test "$(uname)" = "Darwin" ; then
  if test "$(uname -m)" = "x86_64"; then
      ln -s ./dist/protoc-gen-tsx.darwin.amd64 ./dist/protoc-gen-tsx
  else
      ln -s ./dist/protoc-gen-tsx.darwin.arm64 ./dist/protoc-gen-tsx ## Best effort guess
  fi
elif test "$(uname)" = "Linux" ; then
  if test "$(uname -m)" = "x86_64"; then
      ln -s ./dist/protoc-gen-tsx.linux.amd64 ./dist/protoc-gen-tsx
  else
      ln -s ./dist/protoc-gen-tsx.linux.arm64 ./dist/protoc-gen-tsx ## Best effort guess
  fi
fi

chmod a+x ./dist/*

if [ "${CI}" != "true" ]; then
    # Stripping symbols to make the binary more lightweight
    strip -u -r -S ./dist/protoc-gen-tsx.darwin.amd64
fi
