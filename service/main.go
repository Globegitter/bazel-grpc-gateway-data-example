package main

import (
	"context"
	"log"
	"net"

	pb "github.com/globegitter/bazel-grc-gateway-data-example/proto/service"
	"google.golang.org/grpc"
)

func (s *server) GetNewData(ctx context.Context, req *pb.GetRequest) (*pb.Mandate, error) {
	log.Println("Got request")
	return &pb.Mandate{DataId: 1}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterMyServiceServer(s, &server{})
	s.Serve(lis)
}
