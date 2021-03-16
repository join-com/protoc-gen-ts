# Protobuf to TypeScript generator

Generates `ts` files from `proto` file definitions.

## Setup

```
curl -Lo protoc-gen-ts \
    https://github.com/join-com/protoc-gen-ts/releases/download/[VERSION]/protoc-gen-ts-[VERSION] \
&& chmod +x protoc-gen-ts \
&& sudo mv protoc-gen-ts /usr/local/bin
```

## Usage

```
protoc proto/*.proto -I proto --ts_out=proto-ts
```
will generate TS implementations into `proto-ts` folder for all your proto files inside `proto`

## Develop

1. Follow [instructions](https://golang.org/doc/install) to install Go and add /usr/local/go/bin to the PATH environment variable
2. Clone and navigate to repository and build package with `go install .`
3. Inside `integrationTests` file run `yarn proto:build` to generate files for test proto files
4. `yarn test` to run tests
