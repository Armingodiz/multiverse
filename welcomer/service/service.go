package main

import (
	"multiverse/welcomer/welcomepb"
	"net"

	"google.golang.org/grpc"
)

type server struct{}

func (s *server) mustEmbedUnimplementedWelcomeServiceServer() {}

func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	welcomepb.RegisterWelcomeServiceServer(s, &welcomepb.UnimplementedWelcomeServiceServer{})
	if err := s.Serve(listen); err != nil {
		panic(err)
	}
}
