# Proto (WIP)

API definitions for kubelt systems using OpenAPIv3 specifications to generate to protobufs and finally Go.

## Setup
- Install swagger-cli `npm install -g swagger-cli`
- Install gnostic `go install github.com/googleapis/gnostic@latest`
- Install the rpc `go install github.com/google/gnostic-grpc@latest` (you may need to clone this repo instead and build the binary and copy it to your `$GOPATH/bin`)
- Install [protoc](https://grpc.io/docs/protoc-installation/`)

## Run
`make`


# TODO
- Migrate from open api to cuelang for proto generation