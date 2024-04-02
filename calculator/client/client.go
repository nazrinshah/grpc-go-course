package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "grpc-go-course/calculator/proto"
)

var addr = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to dial in: %v\n", err)
	}

	defer conn.Close()

	cl := pb.NewSumServiceClient(conn)

	resp, err := cl.Sum(context.Background(), &pb.SumRequest{
		A: 3,
		B: 10,
	})

	fmt.Println(resp.Result)

	stream, err := cl.Primes(context.Background(), &pb.PrimesRequest{Input: 120})

	if err != nil {
		log.Fatalf("Error calling Primes: %v\n", err)
	}

	for {
		resp, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading stream: %v\n", err)
		}

		log.Printf("Primes: %v\n", resp.Result)
	}
}
