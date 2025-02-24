package main

import (
	"context"
	"fmt"
	"log"

	greet "github.com/NikeshSapkota01/learningGrpc/greet/greetpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	fmt.Println("Hello from client!!!")

	// Establishing connection with the server
	client, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer client.Close()

	// Creating a new client for the GreetService
	c := greet.NewGreetServiceClient(client)

	unary(c)
}

// this is for unary test cases
func unary(c greet.GreetServiceClient) {
	in := &greet.GreetRequest{
		FirstName: "John",
		LastName:  "Doe",
	}

	res, err := c.Greet(context.Background(), in)
	if err != nil {
		log.Fatalf("Error calling Greet: %v", err)
	}

	fmt.Printf("Server Response: %s\n", res.GetMessage())
}
