#!/bin/bash

# Generate Go code for the greet.proto
protoc --go_out=. --go_opt=paths=source_relative \
  --go-grpc_out=. --go-grpc_opt=paths=source_relative greet/greetpb/greet.proto

# Generate Go code for the calculator.proto
protoc --go_out=. --go_opt=paths=source_relative \
  --go-grpc_out=. --go-grpc_opt=paths=source_relative calculator/calculatorpb/calculator.proto