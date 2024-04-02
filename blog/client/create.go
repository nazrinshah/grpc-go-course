package main

import (
	"context"
	"log"

	pb "grpc-go-course/blog/proto"
)

func createBlog(cl pb.BlogServiceClient) string {
	log.Println("createBlog invoked")

	res, err := cl.CreateBlog(context.Background(), &pb.Blog{
		AuthorId: "Bob",
		Title:    "eats meat",
		Content:  "lorem ipsum",
	})

	if err != nil {
		log.Fatalf("Unexpected error: %v\n", err)
	}

	log.Printf("Blog has been created: %s\n", res.Id)
	return res.Id
}
