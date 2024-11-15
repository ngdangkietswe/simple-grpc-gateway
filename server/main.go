package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	gen "simple-grpc-gateway/generated/hello_world"
)

// HelloWorldServerImpl is the implementation of the RPC service defined in protocol definitions.
type HelloWorldServerImpl struct {
	gen.UnimplementedHelloServiceServer
}

// SayHello is the implementation of RPC call defined in protocol definitions.
func (g *HelloWorldServerImpl) SayHello(ctx context.Context, req *gen.HelloReq) (*gen.HelloResp, error) {
	return &gen.HelloResp{
		Message: "Hello " + req.Name,
	}, nil
}

func main() {
	log.Printf("Starting server on port: 8080")

	server := grpc.NewServer()
	gen.RegisterHelloServiceServer(server, &HelloWorldServerImpl{})
	if l, err := net.Listen("tcp", ":8080"); err != nil {
		log.Fatalf("Failed to listen: %v", err)
	} else {
		if err := server.Serve(l); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}
}
