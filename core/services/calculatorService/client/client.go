package client

import (
	"context"
	"fmt"
	"io"
	"multiverse/core/services/calculatorService/calculatorpb"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

type CalculatorClient interface {
	Add(a, b int32) (int32, error)
	PrimeNumberDecomposition(number int64) (factores []int64, err error)
	ComputeAverage(nums []int32) (avg float64, err error)
	FindMaximum(numbers []int32) (max int32, err error)
	Divide(num, advisor int32) (q, r int32, err error)
}

type Client struct {
	grpcClient calculatorpb.CalculatorClient
}

func NewCalculatorConnection(useSsl bool, host, port string) (*grpc.ClientConn, error) {
	var creds credentials.TransportCredentials
	var err error
	if useSsl {
		creds, err = credentials.NewClientTLSFromFile("ssl/ca.crt", "") // Certificate Authority Trust certificate
	} else {
		creds = insecure.NewCredentials()
	}
	if err != nil {
		return nil, err
	}
	return grpc.Dial(fmt.Sprintf("%s:%s", host, port), grpc.WithTransportCredentials(creds)) // for now because we are dont have a certificate
}

func NewCalculatorClient(conn *grpc.ClientConn) (CalculatorClient, error) {
	cli := calculatorpb.NewCalculatorClient(conn)
	return &Client{cli}, nil
}

// func main() {
// 	conn, err := grpc.Dial("localhost:8081", grpc.WithTransportCredentials(insecure.NewCredentials())) // for now because we are dont have a certificate
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer conn.Close()
// 	cli := calculatorpb.NewCalculatorClient(conn)
// 	client := &Client{cli}
// 	response, err := client.Add(10, 20)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("response:", response)
// 	err = client.PrimeNumberDecomposition(int64(26))
// 	if err != nil {
// 		panic(err)
// 	}
// 	err = client.ComputeAverage()
// 	if err != nil {
// 		panic(err)
// 	}
// 	err = client.FindMaximum()
// 	if err != nil {
// 		panic(err)
// 	}
// 	err = client.Divide(10, 2)
// 	if err != nil {
// 		panic(err)
// 	}
// }

func (client *Client) Add(a, b int32) (int32, error) {
	response, err := client.grpcClient.Add(context.Background(), &calculatorpb.AddRequest{
		A: a,
		B: b,
	})
	if err != nil {
		return 0, err
	}
	return response.GetSum(), nil
}

func (client *Client) PrimeNumberDecomposition(number int64) (factores []int64, err error) {
	factores = make([]int64, 0)
	responseStream, err := client.grpcClient.PrimeNumberDecomposition(context.Background(), &calculatorpb.PrimeNumberDecompositionRequest{
		Number: number,
	})
	if err != nil {
		return
	}
	for {
		response, err := responseStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		factores = append(factores, response.GetPrimeFactor())
	}
	return factores, nil
}

func (client *Client) ComputeAverage(nums []int32) (avg float64, err error) {
	stream, err := client.grpcClient.ComputeAverage(context.Background())
	if err != nil {
		return
	}
	for _, num := range nums {
		stream.Send(&calculatorpb.ComputeAverageRequest{
			Numbers: num,
		})
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		return
	}
	return res.GetAverage(), nil
}

func (client *Client) FindMaximum(numbers []int32) (max int32, err error) {
	stream, err := client.grpcClient.FindMaximum(context.Background())
	if err != nil {
		return
	}
	done := make(chan bool)
	errChan := make(chan error)
	go func() {
		for _, num := range numbers {
			err = stream.Send(&calculatorpb.FindMaximumRequest{
				Number: num,
			})
			if err != nil {
				errChan <- err
				return
			}
			time.Sleep(time.Millisecond * 200)
		}
		stream.CloseSend()
	}()
	max = int32(0)
	go func(maxNumber *int32) {
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
			*maxNumber = response.GetMaximum()
		}
	}(&max)
	for {
		select {
		case <-done:
			return
		case err := <-errChan:
			return 0, err
		}
	}
}

func (client *Client) Divide(num, advisor int32) (q, r int32, err error) {
	resp, err := client.grpcClient.Divide(context.Background(), &calculatorpb.DivideRequest{
		Numerator:   num,
		Denominator: advisor,
	})
	if err != nil {
		formattedError, ok := status.FromError(err)
		if ok {
			if formattedError.Code() == codes.InvalidArgument {
				fmt.Println("Divide by zero")
				return
			}
			fmt.Println(formattedError.Message())
			return
		}
		return
	}
	return resp.GetQuotient(), resp.GetRemainder(), nil
}
