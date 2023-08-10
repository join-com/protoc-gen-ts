# Protobuf to TypeScript generator

Generates `ts` files from `proto` file definitions.

## Setup

To install the generator (bear in mind that it could be installed in `${HOME}/.local/bin` too):
```
curl -Lo protoc-gen-tsx \
    https://github.com/join-com/protoc-gen-ts/releases/download/[VERSION]/protoc-gen-tsx.darwin.arm64 \
&& chmod +x protoc-gen-tsx \
&& sudo mv protoc-gen-tsx /usr/local/bin
```

- If you need a binary for Linux, you can change the word "darwin" in the url by "linux".
- If you need support for the Intel processor, or any other ARM processor, change "arm64" by "amd64".

## Usage

will generate TS implementations into `proto-ts` folder for all your proto files inside `proto`:
```
protoc proto/*.proto -I proto --tsx_out=proto-ts
```

## Advanced topics: Required fields & type flavors

In order to enable the "required fields" feature and the "type flavors" feature,
copy the following files from this repository into your `proto` directory:
- `google/protobuf/descriptor.proto`
- `options.proto`

### Required fields

We can mark all fields as required in a message, with a "message custom option":
```proto
message ExampleMessage {
  option (join.protobuf.typescript_required_fields) = true;
  int32 firstField = 1;
  int32 secondField = 2;
  string thirdField = 3 [(join.protobuf.typescript_optional) = true]; // We can mark fields as optional
}
```

We can also mark individual fields as required:
```proto
message ExampleMessage {
  int32 firstField = 1;
  int32 secondField = 2 [(join.protobuf.typescript_required) = true];
  string thirdField = 3;
}
```

### Type flavors

We can also generate flavor nominal types for our fields with primitive types:
```proto
message UserProfile {
  int32 id = 1 [(join.protobuf.typescript_flavor) = "UserId"];
  string username = 2;
  repeated string emails = 3 [(join.protobuf.typescript_flavor) = "Email"];
}
```

### Field masks

By default, the generated type for the field mask property will be `{ paths: string[] }`. To add clarity which fields can be send as a mask use:
```proto
message GetRequest {
  int32 id = 1;
  google.protobuf.FieldMask read_mask = 2 [(join.protobuf.typescript_mask) = "Company"];
}
```

## Develop

1. Follow [instructions](https://golang.org/doinc/install) to install Go and add /usr/local/go/bin to the PATH environment variable
2. Run `./build.sh` to compile the package
3. Or navigate to `tests` folder and run `yarn compile`
4. Run `yarn proto:build:package(1/2/3)` to generate one of the packages for tests
5. Run `yarn test` to run tests
