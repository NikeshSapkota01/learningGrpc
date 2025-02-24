package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	greet "github.com/NikeshSapkota01/learningGrpc/greet/greetpb"
	"google.golang.org/grpc"
)

type server struct {
	greet.UnimplementedGreetServiceServer
}

// Greet method for unary RPC
func (*server) Greet(ctx context.Context, req *greet.GreetRequest) (*greet.GreetResponse, error) {
	fmt.Printf("greet %v\n", req.String())
	fn := req.GetFirstName()
	ln := req.GetLastName()

	// Create a response
	res := &greet.GreetResponse{
		Message: fmt.Sprintf("Hello, %s %s!", fn, ln),
	}

	return res, nil
}

// ServerGreetStream method for server streaming RPC
func (*server) ServerGreetStream(req *greet.ServerGreetStreamRequest, stream greet.GreetService_ServerGreetStreamServer) error {
	fmt.Printf("ServerGreetStream function called with: %v\n", req.String())
	fn := req.GetFirstName()
	ln := req.GetLastName()

	// Send 5 streaming responses
	for i := 0; i < 5; i++ {
		res := &greet.ServerGreetStreamResponse{
			Message: fmt.Sprintf("Hello, %s %s! Message %d", fn, ln, i+1),
		}

		if err := stream.Send(res); err != nil {
			return fmt.Errorf("failed to send message: %v", err)
		}

		// Simulate some delay between messages
		time.Sleep(1 * time.Second)
	}

	return nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	greet.RegisterGreetServiceServer(s, &server{})

	fmt.Println("Server is running on port 50051")

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
