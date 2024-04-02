package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	pb "grpc-go-course/greet/proto"
)

var (
	addr string = "0.0.0.0:50051"
)

type Server struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen on : %v\n", err)
	}

	log.Printf("Listening on %s\n", addr)

	svr := grpc.NewServer()
	pb.RegisterGreetServiceServer(svr, Server{})

	if err = svr.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
