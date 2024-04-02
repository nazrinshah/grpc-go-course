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

func (s Server) GreetManyTimes(in *pb.GreetRequest, stream pb.GreetService_GreetManyTimesServer) error {
	log.Printf("GreetManyTimes function was invoked with: %v\n", in)

	for i := 0; i < 10; i++ {
		res := fmt.Sprintf("Hello %s, number %d", in.FirstName, i)
		stream.Send(&pb.GreetResponse{Result: res})
	}

	return nil
}
