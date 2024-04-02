package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "grpc-go-course/blog/proto"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to dial in: %v\n", err)
	}

	defer conn.Close()

	cl := pb.NewBlogServiceClient(conn)
	//createBlog(cl)
	//readBlog(cl, "660c2f755695f9a3097d41f4")
	//updateBlog(cl, "660c2f755695f9a3097d41f4")
	//readBlog(cl, "660c2f755695f9a3097d41f4")
	listBlogs(cl)
}
