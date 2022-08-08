package main

import (
	"context"
	"fmt"
	"io"
	"multiverse/calculator/calculatorpb"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
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
	err = client.FindMaximum()
	if err != nil {
		panic(err)
	}
	err = client.Divide(10, 2)
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

func (client *Client) FindMaximum() error {
	stream, err := client.grpcClient.FindMaximum(context.Background())
	if err != nil {
		return err
	}
	done := make(chan bool)
	errChan := make(chan error)
	go func() {
		for i := 0; i < 10; i++ {
			err = stream.Send(&calculatorpb.FindMaximumRequest{
				Number: int32(i),
			})
			if err != nil {
				errChan <- err
				return
			}
			time.Sleep(time.Millisecond * 200)
		}
		stream.CloseSend()
	}()
	go func() {
		for {
			response, err := stream.Recv()
			if err == io.EOF {
				done <- true
				return
			}
			if err != nil {
				errChan <- err
				return
			}
			fmt.Println("current", response)
		}
	}()
	for {
		select {
		case <-done:
			return nil
		case err := <-errChan:
			return err
		}
	}
}

func (client *Client) Divide(num, advisor int) error {
	resp, err := client.grpcClient.Divide(context.Background(), &calculatorpb.DivideRequest{
		Numerator:   int32(num),
		Denominator: int32(advisor),
	})
	if err != nil {
		formattedError, ok := status.FromError(err)
		if ok {
			if formattedError.Code() == codes.InvalidArgument {
				fmt.Println("Divide by zero")
				return nil
			}
			fmt.Println(formattedError.Message())
			return nil
		}
		return err
	}
	fmt.Printf("Divide response: quotient= %d, reminder= %d\n", resp.GetQuotient(), resp.GetRemainder())
	return nil
}
