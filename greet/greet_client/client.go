package main

import (
	"context"
	"fmt"
	"log"

	"github.com/nel417/grpc-train/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	//no ssl so it will be with insecure
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("couldnt connect %v", err)
	}
	defer cc.Close()
	c := greetpb.NewGreetServiceClient(cc)

	doUnary(c)
}

func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("Starting...")
	req := &greetpb.GreetRequest {
		Greeting: &greetpb.Greeting{
			FirstName: "Nick",
			LastName: "Landreville",
		},
	}
	res, err:= c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Greet RPC: %v", err)
	}
	log.Printf("Response from Greet: %v", res.Result)
}