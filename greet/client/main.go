package main

import (
	"log"

	"github.com/quistew/golang-microservices/greet/client/framework"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/quistew/golang-microservices/greet/proto"
)

var addr string = "127.0.0.1:5051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}

	defer conn.Close()

	c := pb.NewGreetServiceClient(conn)

	gs, err := framework.NewGreetService(c)
	if err != nil {
		log.Fatalf("error creating greet service: %v\n", err)
	}

	gs.DoGreet("eli")
}
