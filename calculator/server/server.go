package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/NikeshSapkota01/learningGrpc/calculator/calculatorpb"
	"google.golang.org/grpc"
)

type server struct {
	calculatorpb.UnimplementedCalculatorServiceServer
}

// why i did this learn from here => https://grpc.io/docs/languages/go/quickstart/
func (*server) Calculate(ctx context.Context, req *calculatorpb.CalculatorRequest) (*calculatorpb.CalculatorResponse, error) {
	fmt.Printf("calculate %v\n", req.String())
	fn := req.GetFirstDigit()
	ln := req.GetSecondDigit()
	op := req.GetOperation()

	var result int64 // Use int64 to avoid overflow during large calculations

	// Perform the operation based on the selected operation
	switch op {
	case calculatorpb.Operation_ADD:
		result = int64(fn) + int64(ln)
	case calculatorpb.Operation_SUBTRACT:
		result = int64(fn) - int64(ln)
	case calculatorpb.Operation_MULTIPLY:
		result = int64(fn) * int64(ln)
	case calculatorpb.Operation_DIVIDE:
		if ln != 0 {
			result = int64(fn) / int64(ln)
		} else {
			return nil, fmt.Errorf("division by zero is not allowed")
		}
	default:
		return nil, fmt.Errorf("unsupported operation")
	}

	res := &calculatorpb.CalculatorResponse{
		Result: result,
	}

	return res, nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	fmt.Println("Server is running on port 50051")

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
