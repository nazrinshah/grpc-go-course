package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "grpc-go-course/sum/proto"
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
}
