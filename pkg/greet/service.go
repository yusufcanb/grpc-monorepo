package greet

import (
	"context"
	pb "github.com/yusufcanb/grpc-monorepo/pkg/protocol"
	"log"
)

type GreeterServer struct {
	pb.UnimplementedGreeterServer
}

func (it *GreeterServer) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}
