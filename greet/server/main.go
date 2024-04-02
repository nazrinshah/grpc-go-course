package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

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

	tls := true // change to false if needed
	opts := []grpc.ServerOption{}

	if tls {
		certFile := "ssl/server.crt"
		key := "ssl/server.pem"
		creds, err := credentials.NewServerTLSFromFile(certFile, key)

		if err != nil {
			log.Fatalf("Failed loading certs: %v\n", err)
		}

		opts = append(opts, grpc.Creds(creds))
	}

	svr := grpc.NewServer(opts...)
	pb.RegisterGreetServiceServer(svr, Server{})

	if err = svr.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
