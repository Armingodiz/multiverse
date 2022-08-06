package main

import (
	"context"
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
