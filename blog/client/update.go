package main

import (
	"context"
	"log"

	pb "grpc-go-course/blog/proto"
)

func updateBlog(cl pb.BlogServiceClient, id string) {
	log.Println("updateBlog was invoked")

	newBlog := &pb.Blog{
		Id:       id,
		AuthorId: "others",
		Title:    "blob",
		Content:  "saufghsafhsa",
	}

	_, err := cl.UpdateBlog(context.Background(), newBlog)

	if err != nil {
		log.Fatalf("Error happened while updating: %v\n", err)
	}

	log.Println("Blog was updated")
}