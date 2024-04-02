package main

import (
	"context"
	"log"
	"net"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"

	pb "grpc-go-course/blog/proto"
)

type Server struct {
	pb.BlogServiceServer
}

var (
	addr       string = "0.0.0.0:50051"
	collection *mongo.Collection
)

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:root@localhost:27017/"))

	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(context.Background())

	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database("blogdb").Collection("blog")

	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", addr)
	}

	svr := grpc.NewServer()
	pb.RegisterBlogServiceServer(svr, Server{})

	if err = svr.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
