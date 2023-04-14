# Go /w gRPC

A Go project setup to develop server-client applications with [Protocol Buffers](https://protobuf.dev/) and [Cobra](https://github.com/spf13/cobra). 

## Structure

```
├── cmd
│   ├── cli           
│   │   ├── project.go
│   │   └── root.go   
│   └── server        
│       └── main.go
├── pkg
    ├── protocol
        └── greet.proto

```

- `cmd/cli`: A CLI program entrypoint with Cobra
- `cmd/server`: GRPC Server entrypoint
- `pkg/protocol`: Contains Protocol Buffer files and generated code.

## Compile Proto Files

```
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ../pkg/protocol/*.proto 
```

