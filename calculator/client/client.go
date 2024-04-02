package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "grpc-go-course/calculator/proto"
)

var addr = "localhost:50051"

func doSum(cl pb.SumServiceClient) {
	resp, _ := cl.Sum(context.Background(), &pb.SumRequest{
		A: 3,
		B: 10,
	})

	fmt.Println(resp.Result)
}

func doPrimes(cl pb.SumServiceClient) {
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

func doAverage(cl pb.SumServiceClient) {
	reqs := []*pb.AverageRequest{
		{
			Input: 1,
		},
		{
			Input: 2,
		},
		{
			Input: 3,
		},
		{
			Input: 4,
		},
	}

	stream1, err := cl.Average(context.Background())

	if err != nil {
		log.Fatalf("Error while calling Average: %v\n", err)
	}

	for _, req := range reqs {
		log.Println("Sending req: ", req)
		stream1.Send(req)
	}

	res1, err := stream1.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while receiving response from Average: %v\n", err)
	}

	log.Printf("Average: %v\n", res1.Result)
}

func doMax(cl pb.SumServiceClient) {
	stream, err := cl.Max(context.Background())

	if err != nil {
		log.Fatalf("Error while creating stream: %v", err)
	}

	reqs := []*pb.MaxRequest{
		{
			Input: 1,
		},
		{
			Input: 5,
		},
		{
			Input: 3,
		},
		{
			Input: 6,
		},
		{
			Input: 2,
		},
		{
			Input: 20,
		},
	}

	waitc := make(chan struct{})
	go func() {
		for _, req := range reqs {
			log.Printf("Send req: %v\n", req)
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}

		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalf("Error while receiving: %v", err)
			}

			log.Printf("Received: %v\n", res.Result)
		}

		close(waitc)
	}()

	<-waitc
}

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to dial in: %v\n", err)
	}

	defer conn.Close()

	cl := pb.NewSumServiceClient(conn)

	//doSum(cl)
	//doPrimes(cl)
	//doAverage(cl)
	doMax(cl)
}
