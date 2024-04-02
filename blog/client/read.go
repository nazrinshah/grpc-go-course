package main

import (
	"context"
	"log"

	pb "grpc-go-course/blog/proto"
)

func readBlog(cl pb.BlogServiceClient, id string) *pb.Blog {
	log.Println("readBlog was invoked")

	req := &pb.BlogId{Id: id}
	res, err := cl.ReadBlog(context.Background(), req)

	if err != nil {
		log.Printf("Error happened while readin: %v\n", err)
	}

	log.Printf("Blog was read: %v\n", res)
	return res
}
