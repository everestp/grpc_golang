package main

import (
	"context"
	"log"
	"net"

	pb "simplegrpcserver/proto/gen"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnsafeCalculateServer
}

func (s *server) Add(ctx context.Context, req *pb.AddRequest) (*pb.AddResponse, error) {
	return &pb.AddResponse{
		Sum: req.A + req.B,
	}, nil
}

func main() {
	port := ":5001"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("Failed to listen", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterCalculateServer(grpcServer, &server{})

	log.Println("Server is running on", port)
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatal("Failed to server:", err)
	}
}
