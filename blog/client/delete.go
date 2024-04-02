package main

import (
	"context"
	"log"

	pb "grpc-go-course/blog/proto"
)

func deleteBlog(cl pb.BlogServiceClient, id string) {
	log.Println("deleteBlog was called")

	_, err := cl.DeleteBlog(context.Background(), &pb.BlogId{Id: id})

	if err != nil {
		log.Fatalf("Error while deleting: %v\n", err)
	}

	log.Println("Blog was deleted")
}
