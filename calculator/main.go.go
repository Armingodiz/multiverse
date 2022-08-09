package main

import (
	"context"
	"io"
	"multiverse/calculator/calculatorpb"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
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

func (s *server) Divide(ctx context.Context, in *calculatorpb.DivideRequest) (*calculatorpb.DivideResponse, error) {
	divisor := in.GetDenominator()
	if divisor == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "divisor cannot be zero")
	}
	return &calculatorpb.DivideResponse{
		Quotient:  in.GetNumerator() / divisor,
		Remainder: in.GetNumerator() % divisor,
	}, nil
}

func main() {
	connectionType := os.Getenv("CONNECTION_TYPE")
	port := os.Getenv("CONNECTION_PORT")
	useSSl := os.Getenv("USE_SSL")
	useSsl := false
	if connectionType == "" {
		connectionType = "tcp"
	}
	if port == "" {
		port = ":8082"
	}
	if useSSl == "" || useSSl == "false" {
		useSsl = false
	} else {
		useSsl = true
	}
	listen, err := net.Listen(connectionType, port)
	if err != nil {
		panic(err)
	}
	var creds credentials.TransportCredentials
	if useSsl {
		creds, err = credentials.NewServerTLSFromFile("ssl/server.crt", "ssl/server.pem")
		if err != nil {
			panic(err)
		}
	}
	s := grpc.NewServer(grpc.Creds(creds))
	calculatorpb.RegisterCalculatorServer(s, &server{})
	if err := s.Serve(listen); err != nil {
		panic(err)
	}
}
