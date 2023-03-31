package greet

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"

	pb "github.com/yusufcanb/grpc-monorepo/pkg/protocol"
)

func TestGreeterServer_SayHello(t *testing.T) {
	server := GreeterServer{}

	sayHello, err := server.SayHello(context.Background(), &pb.HelloRequest{Name: "josef"})
	if err != nil {
		t.Errorf("Something went wrong on SayHello procedure")
	}

	assert.Equal(t, "Hello, josef", sayHello.GetMessage())
}
