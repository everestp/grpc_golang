package main

import (
	"context"
	mainapipb "grpcstreamclient/proto/gen"
	"io"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("Did not connect :", err)
	}
	defer conn.Close()
	client := mainapipb.NewCalculatorClient(conn)
	ctx := context.Background()

	req := &mainapipb.FibonacciRequest{
		N: 10,
	}
	stream, err := client.GenerateFibonacci(ctx, req)
	if err != nil {
		log.Fatalln("Error calling GenerateFibonacci func", err)
	}
	for {
		response, err := stream.Recv()
		if err == io.EOF {
			log.Println("End of stream")
		}
		if err != nil {
			log.Fatal("Error revieving data from GenerateFibonacci func", err)
		}
		log.Println("Fibonacci number:", response)
	}

}
