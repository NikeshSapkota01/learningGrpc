package main

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/NikeshSapkota01/learningGrpc/calculator/calculatorpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	fmt.Println("Hello from client of calculator!!!")

	// Establishing connection with the server
	client, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer client.Close()

	c := calculatorpb.NewCalculatorServiceClient(client)

	firstDigit, secondDigit, operation, err := getInput()
	if err != nil {
		log.Fatalf("Input error: %v", err)
	}

	// Call the unary function to perform the calculation
	unary(c, firstDigit, secondDigit, operation)
}

// Unary call to the server for the calculation
func unary(c calculatorpb.CalculatorServiceClient, firstDigit int32, secondDigit int32, op calculatorpb.Operation) {
	in := &calculatorpb.CalculatorRequest{
		FirstDigit:  firstDigit,
		SecondDigit: secondDigit,
		Operation:   op,
	}

	res, err := c.Calculate(context.Background(), in)
	if err != nil {
		log.Fatalf("Error calling calculator: %v", err)
	}

	fmt.Printf("Server Response: Result = %d\n", res.GetResult())
}

// Function to ask for input from the user and validate the input
func getInput() (int32, int32, calculatorpb.Operation, error) {
	var firstDigit, secondDigit int32
	var operationStr string

	fmt.Println("Enter the first number (int32): ")
	_, err := fmt.Scanf("%d", &firstDigit)
	if err != nil {
		return 0, 0, 0, fmt.Errorf("invalid input for first number: %v", err)
	}

	fmt.Println("Enter the second number (int32): ")
	_, err = fmt.Scanf("%d", &secondDigit)
	if err != nil {
		return 0, 0, 0, fmt.Errorf("invalid input for second number: %v", err)
	}

	// Get operation input (can be empty)
	fmt.Println("Enter the operation (1: ADD, 2: SUBTRACT, 3: MULTIPLY, 4: DIVIDE, leave blank for default): ")
	_, err = fmt.Scanf("%s", &operationStr)
	if err != nil {
		return 0, 0, 0, fmt.Errorf("invalid input for operation: %v", err)
	}

	// Default the operation to 0 (ADD) if input is empty
	var op calculatorpb.Operation
	if operationStr == "" {
		op = calculatorpb.Operation_ADD
	} else {
		// Try to parse operation input if it's not empty
		operation, err := strconv.Atoi(operationStr)
		if err != nil || operation < 1 || operation > 4 {
			return 0, 0, 0, fmt.Errorf("invalid input for operation: %v", err)
		}

		// Convert operation input to corresponding enum
		switch operation {
		case 1:
			op = calculatorpb.Operation_ADD
		case 2:
			op = calculatorpb.Operation_SUBTRACT
		case 3:
			op = calculatorpb.Operation_MULTIPLY
		case 4:
			op = calculatorpb.Operation_DIVIDE
		}
	}

	return firstDigit, secondDigit, op, nil
}
