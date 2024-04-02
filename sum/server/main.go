package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "grpc-go-course/sum/proto"
)

var (
	addr string = "localhost:50051"
)

type Server struct {
	pb.SumServiceServer
}

func (s Server) Sum(ctx context.Context, req *pb.SumRequest) (*pb.SumResponse, error) {
	return &pb.SumResponse{Result: req.A + req.B}, nil
}

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", addr)
	}

	svr := grpc.NewServer()
	pb.RegisterSumServiceServer(svr, Server{})

	if err = svr.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
