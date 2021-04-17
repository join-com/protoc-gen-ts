# Protobuf to TypeScript generator

Generates `ts` files from `proto` file definitions.

## Setup

To install the "new" generator (bear in mind that it could be installed in `${HOME}/.local/bin` too):
```
curl -Lo protoc-gen-ts \
    https://github.com/join-com/protoc-gen-ts/releases/download/[VERSION]/protoc-gen-tsx.darwin.x86_64 \
&& chmod +x protoc-gen-tsx \
&& sudo mv protoc-gen-tsx /usr/local/bin
```

In case you need to use the old generator as well, you can install it too in parallel:
```
curl -Lo protoc-gen-ts \
    https://github.com/join-com/protoc-gen-ts/releases/download/0.7.1/protoc-gen-ts.darwin.x86_64 \
&& chmod +x protoc-gen-ts \
&& sudo mv protoc-gen-ts /usr/local/bin
```

## Usage

will generate TS implementations into `proto-ts` folder for all your proto files inside `proto`:
```
protoc proto/*.proto -I proto --tsx_out=proto-ts
```

If you want to use the old generator (notice the `tsx_out` -> `ts_out` change):
```
protoc proto/*.proto -I proto --ts_out=proto-ts
```

## Develop

1. Follow [instructions](https://golang.org/doc/install) to install Go and add /usr/local/go/bin to the PATH environment variable
2. Clone and navigate to repository and build package with `go install .`
3. Inside `integrationTests` file run `yarn proto:build` to generate files for test proto files
4. `yarn test` to run tests
