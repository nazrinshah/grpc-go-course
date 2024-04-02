package main

import (
	"context"
	"io"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"

	pb "grpc-go-course/blog/proto"
)

func listBlogs(cl pb.BlogServiceClient) {
	log.Printf("listBlogs was invoked")

	stream, err := cl.ListBlogs(context.Background(), &emptypb.Empty{})

	if err != nil {
		log.Fatalf("Error while calling ListBlogs: %v\n", err)
	}

	blogs := []*pb.Blog{}
	for {
		blog, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Something happened: %v\n", err)
		}

		blogs = append(blogs, blog)
	}

	log.Println(blogs)
}
