package main

import (
	"flag"
	"fmt"
	"github.com/yusufcanb/grpc-monorepo/pkg/greet"
	"log"
	"net"

	pb "github.com/yusufcanb/grpc-monorepo/pkg/protocol"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 5001, "The server port")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &greet.GreeterServer{})
	log.Printf("GRPC Greet Server started at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
