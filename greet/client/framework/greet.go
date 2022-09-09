package framework

import (
	"context"
	"fmt"
	"log"

	pb "github.com/quistew/golang-microservices/greet/proto"
)

type GreetService struct {
	c pb.GreetServiceClient
}

func NewGreetService(c pb.GreetServiceClient) (*GreetService, error) {
	if c == nil {
		return nil, fmt.Errorf("client cannot be nil")
	}

	return &GreetService{
		c: c,
	}, nil
}

func (gs *GreetService) DoGreet(name string) {
	log.Println("doGreet was invoked")
	res, err := gs.c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: name,
	})

	if err != nil {
		log.Fatalf("could not greet: %v\n", err)
	}

	log.Printf("Greeting: %s\n", res.Result)
}
