package main

import (
	"context"
	"io"
	"multiverse/calculator/calculatorpb"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	calculatorpb.UnimplementedCalculatorServer
}

func (s *server) Add(ctx context.Context, in *calculatorpb.AddRequest) (*calculatorpb.AddResponse, error) {
	return &calculatorpb.AddResponse{Sum: in.GetA() + in.GetB()}, nil
}

func (s *server) PrimeNumberDecomposition(in *calculatorpb.PrimeNumberDecompositionRequest, stream calculatorpb.Calculator_PrimeNumberDecompositionServer) error {
	n := in.GetNumber()
	advisor := int64(2)
	for n > 1 {
		if n%advisor == 0 {
			n = n / advisor
			stream.Send(&calculatorpb.PrimeNumberDecompositionResponse{
				PrimeFactor: advisor,
			})
		} else {
			advisor++
		}
	}
	return nil
}

func (s *server) ComputeAverage(stream calculatorpb.Calculator_ComputeAverageServer) error {
	var sum int32
	var count int32
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&calculatorpb.ComputeAverageResponse{
				Average: float64(sum) / float64(count),
			})
		}
		if err != nil {
			return err
		}
		sum += in.GetNumbers()
		count++
	}
}

func (s *server) FindMaximum(stream calculatorpb.Calculator_FindMaximumServer) error {
	var max int32
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		if in.GetNumber() > max {
			max = in.GetNumber()
			stream.Send(&calculatorpb.FindMaximumResponse{
				Maximum: max,
			})
		}
	}
}

func main() {
	listen, err := net.Listen("tcp", ":8081")
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServer(s, &server{})
	if err := s.Serve(listen); err != nil {
		panic(err)
	}
}
