package main

import (
	"context"
	"fmt"
	"log"

	pb "grpc-go-course/greet/proto"
)

func (s Server) Greet(ctx context.Context, req *pb.GreetRequest) (*pb.GreetResponse, error) {

	log.Printf("Greet function was invoked with %v\n", req)

	return &pb.GreetResponse{
		Result: fmt.Sprintf("Hello %s", req.FirstName),
	}, nil
}
