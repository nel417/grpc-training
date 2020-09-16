package main

import (
	"context"
	"fmt"
	"io"
	"log"

	// "github.com/nel417/grpc-train/greet/greetpb"
	"github.com/nel417/grpc-train/calculator/calculatorpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Calculator Client")
	//no ssl so it will be with insecure
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("couldnt connect %v", err)
	}
	defer cc.Close()
	c := calculatorpb.NewCalculatorServiceClient(cc)

	doUnary(c)
	doServerStreaming(c)
}

func doUnary(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting...")
	req := &calculatorpb.SumRequest{
		FirstNumber:  100,
		SecondNumber: 225,
	}
	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Sum RPC: %v", err)
	}
	log.Printf("Response from Sum: %v", res.SumResult)
}

func doServerStreaming(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting prime decomp...")
	req := &calculatorpb.PrimeNumberDecompositionRequest{
		Number: 125647362,
	}

	stream, err := c.PrimeNumberDecomposition(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Prime RPC: %v", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("SOmething happened: %v", err)
		}
		fmt.Println(res.GetPrimeFactor())
	}
}
