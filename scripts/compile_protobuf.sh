protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=path
s=source_relative ../pkg/protocol/*.proto
