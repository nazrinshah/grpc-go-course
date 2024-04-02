package main

import (
	"context"
	"fmt"
	"io"
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

func (s Server) LongGreet(stream pb.GreetService_LongGreetServer) error {
	log.Printf("LongGreet function was invoked")

	res := ""

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.GreetResponse{Result: res})
		}

		if err != nil {
			log.Fatalf("Error whle reading client stream: %v\n", err)
		}

		res = fmt.Sprintf("%sHello %s!\n", res, req.FirstName)
	}

	return nil
}
