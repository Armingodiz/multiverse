package main

import (
	"context"
	"multiverse/welcomer/welcomepb"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	welcomepb.UnimplementedWelcomeServiceServer
}

func (s *server) Welcome(ctx context.Context, in *welcomepb.WelcomeRequest) (*welcomepb.WelcomeResponse, error) {
	arrival := in.GetArrival().String()
	return &welcomepb.WelcomeResponse{Result: "Hello " + in.User.Name + " from: " + in.User.Country + " you came at " + arrival}, nil
}

func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	welcomepb.RegisterWelcomeServiceServer(s, &server{})
	if err := s.Serve(listen); err != nil {
		panic(err)
	}
}
