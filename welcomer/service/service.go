package main

import (
	"context"
	"fmt"
	"multiverse/welcomer/welcomepb"
	"net"
	"time"

	"google.golang.org/grpc"
)

type server struct {
	welcomepb.UnimplementedWelcomeServiceServer
}

func (s *server) Welcome(ctx context.Context, in *welcomepb.WelcomeRequest) (*welcomepb.WelcomeResponse, error) {
	arrival := in.GetArrival().String()
	return &welcomepb.WelcomeResponse{Result: "Hello " + in.User.Name + " from: " + in.User.Country + " you came at " + arrival}, nil
}

func (s *server) GetGreetings(in *welcomepb.WelcomeRequest, stream welcomepb.WelcomeService_GetGreetingsServer) error {
	arrival := in.GetArrival().String()
	for i := 0; i < 10; i++ {
		arrival = arrival + "passed " + fmt.Sprintf("%d", i) + " times"
		stream.Send(&welcomepb.WelcomeResponse{Result: "Hello " + in.User.Name + " from: " + in.User.Country + " you came at " + arrival})
		time.Sleep(time.Second)
	}
	return nil
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
