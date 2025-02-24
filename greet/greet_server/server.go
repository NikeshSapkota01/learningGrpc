package main

import (
	"context"
	"fmt"
	"log"
	"net"

	greet "example.com/grpc/greet/greetpb"
	"google.golang.org/grpc"
)

type server struct {
	greet.UnimplementedGreetServiceServer
}

// func (s *server) MethodName(ctx context.Context, req *pb.RequestType) (*pb.ResponseType, error)

func (*server) Greet(ctx context.Context, req *greet.GreetRequest) (*greet.GreetResponse, error) {
	fmt.Printf("greet hehehe %v\n", req.String())
	fn := req.GetFirstName()
	ln := req.GetLastName()

	// Create a response
	res := &greet.GreetResponse{
		Message: fmt.Sprintf("Hello, %s %s!", fn, ln),
	}

	return res, nil
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
