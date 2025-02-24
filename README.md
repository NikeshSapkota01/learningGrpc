# Learning gRPC

This repository contains a project for learning gRPC and building services using Protocol Buffers and Go.

## Project Structure

- `greet/`: Contains the greeting service (`greet.proto`) and generated code.
- `calculator/`: Contains the calculator service (`calculator.proto`) and generated code.
- `generate.sh`: A shell script for generating Go code from `.proto` files using `protoc`.
- `go.mod`: Go module file for dependency management.
- `server.go`: gRPC server implementation.
- `client.go`: gRPC client implementation.

## Prerequisites

Before running the project, make sure you have the following installed:

- Go (preferably version 1.16+)
- `protoc` (Protocol Buffers Compiler)
- `protoc-gen-go` (Go plugin for protoc)
- `protoc-gen-go-grpc` (Go plugin for gRPC)

To install the necessary tools, run the following:

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

