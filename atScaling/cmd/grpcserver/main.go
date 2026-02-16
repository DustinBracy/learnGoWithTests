package main

import (
	"context"
	"log"
	"net"

	"github.com/dustinbracy/learnGoWithTests/atScaling/adapters/grpcserver"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	grpcserver.RegisterGreeterServer(s, &GreetServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

type GreetServer struct {
	grpcserver.UnimplementedGreeterServer
}

func (s *GreetServer) Greet(ctx context.Context, req *grpcserver.GreetRequest) (*grpcserver.GreetReply, error) {
	return &grpcserver.GreetReply{Message: "Hello, " + req.Name}, nil
}

func (s *GreetServer) Curse(ctx context.Context, req *grpcserver.CurseRequest) (*grpcserver.CurseReply, error) {
	return &grpcserver.CurseReply{Message: "Go away, " + req.Name}, nil
}

// GreeterServer is the server API for Greeter service.
// All implementations must embed UnimplementedGreeterServer
// for forward compatibility.
type GreeterServer interface {
	Greet(context.Context, *grpcserver.GreetRequest) (*grpcserver.GreetReply, error)
	mustEmbedUnimplementedGreeterServer()
}
