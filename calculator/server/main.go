package main

import (
	"context"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "grpc-go-course/calculator/proto"
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

func (s Server) Primes(in *pb.PrimesRequest, stream pb.SumService_PrimesServer) error {
	k := int32(2)
	N := in.Input
	for N > 1 {
		if N%k == 0 {
			N = N / k
			stream.Send(&pb.PrimesResponse{Result: k})
		} else {
			k = k + 1
		}
	}

	return nil
}

func (s Server) Average(stream pb.SumService_AverageServer) error {
	sum := int32(0)
	i := int32(0)

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.AverageResponse{Result: float32(sum) / float32(i)})
		}

		if err != nil {
			log.Fatalf("Error while receiving requests: %v\n", err)
		}

		i = i + 1
		sum = sum + req.Input
	}

	return nil
}

func (s Server) Max(stream pb.SumService_MaxServer) error {
	max := int32(0)

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error reading stream")
		}

		if max < req.GetInput() {
			max = req.GetInput()
			stream.Send(&pb.MaxResponse{Result: max})
		}
	}

	return nil
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
