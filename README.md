# Go /w gRPC

A Go project setup to develop server-client applications with [Protocol Buffers](https://protobuf.dev/) and [Cobra](https://github.com/spf13/cobra). 

## Overview

The general workflow of the project contains 3 steps;

- Compiling proto files
- Registering your service to gRPC server
- Implement CLI usage in Cobra

### Structure

```
├── cmd
│   ├── cli           
│   │   ├── project.go
│   │   └── root.go   
│   └── server        
│       └── main.go
├── pkg
│   ├── greet
│   │   ├── project.go
│   │   └── root.go        
│   └── protocol
│        └── greet.proto

```

- `cmd/cli`: CLI program entrypoint
- `cmd/server`: GRPC Server entrypoint
- `pkg/protocol`: Contains Protocol Buffer files and generated code.

## Getting Started

### Development

#### Compile Proto Files


```
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ../pkg/protocol/*.proto 
```

#### Start gRPC Server

```
go run cmd/server/main.go
```

#### Call gRPC with CLI

```
go run main.go greet
```


### Build Project

- Build CLI application with the following;

  ```
  go build -o grpc-cli main.go
  ```

- Build gRPC server application with below;

  ```
  go build -o grpc-server cmd/server/main.go
  ```
