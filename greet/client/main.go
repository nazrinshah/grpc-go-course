package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "grpc-go-course/greet/proto"
)

var addr string = "localhost:50051"

func main() {
	tls := true
	opts := []grpc.DialOption{}

	if tls {
		certFile := "ssl/ca.crt"
		creds, err := credentials.NewClientTLSFromFile(certFile, "")

		if err != nil {
			log.Fatalf("Erropr while loading CA trust certificate: %v\n", err)
		}

		opts = append(opts, grpc.WithTransportCredentials(creds))
	}

	conn, err := grpc.Dial(addr, opts...)

	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}

	defer conn.Close()

	cl := pb.NewGreetServiceClient(conn)
	doGreet(cl)
	//doGreetManyTimes(cl)
	//doLongGreet(cl)
	//doGreetEveryone(cl)
	//doGreetWithDeadline(cl, 5*time.Second)
	//doGreetWithDeadline(cl, 1*time.Second)
}
