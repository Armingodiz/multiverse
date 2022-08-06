package main

import (
	"context"
	"fmt"
	"io"
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
	err = client.PrimeNumberDecomposition(int64(26))
	if err != nil {
		panic(err)
	}
	err = client.ComputeAverage()
	if err != nil {
		panic(err)
	}
}

func (client *Client) Add(a, b int32) (*calculatorpb.AddResponse, error) {
	response, err := client.grpcClient.Add(context.Background(), &calculatorpb.AddRequest{
		A: a,
		B: b,
	})
	return response, err
}

func (client *Client) PrimeNumberDecomposition(number int64) error {
	responseStream, err := client.grpcClient.PrimeNumberDecomposition(context.Background(), &calculatorpb.PrimeNumberDecompositionRequest{
		Number: number,
	})
	if err != nil {
		return err
	}
	fmt.Print("Prime number decomposition: ")
	for {
		response, err := responseStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		fmt.Print(response.GetPrimeFactor())
		fmt.Print(" ")
	}
	fmt.Println()
	return nil
}

func (client *Client) ComputeAverage() error {
	stream, err := client.grpcClient.ComputeAverage(context.Background())
	if err != nil {
		return err
	}
	for i := 10; i < 200; i += 34 {
		stream.Send(&calculatorpb.ComputeAverageRequest{
			Numbers: int32(i),
		})
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		return err
	}
	fmt.Println("average of getAverage requests: ", res.GetAverage())
	return nil
}
