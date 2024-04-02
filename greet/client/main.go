package main

import (
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "grpc-go-course/greet/proto"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}

	defer conn.Close()

	cl := pb.NewGreetServiceClient(conn)
	//doGreet(cl)
	//doGreetManyTimes(cl)
	//doLongGreet(cl)
	//doGreetEveryone(cl)
	doGreetWithDeadline(cl, 5*time.Second)
	doGreetWithDeadline(cl, 1*time.Second)
}
