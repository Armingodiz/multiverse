package main

import (
	"context"
	"fmt"
	"multiverse/calculator/calculatorpb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	grpcClient calculatorpb.CalculatorClient
}

func main() {
	conn, err := grpc.Dial("localhost:8081", grpc.WithTransportCredentials(insecure.NewCredentials())) // for now because we are dont have a certificate
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	cli := calculatorpb.NewCalculatorClient(conn)
	client := &Client{cli}
	response, err := client.Add(10, 20)
	if err != nil {
		panic(err)
	}
	fmt.Println("response:", response)
}

func (client *Client) Add(a, b int32) (*calculatorpb.AddResponse, error) {
	response, err := client.grpcClient.Add(context.Background(), &calculatorpb.AddRequest{
		A: a,
		B: b,
	})
	return response, err
}
