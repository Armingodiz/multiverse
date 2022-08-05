package service

import (
	"multiverse/welcomer/welcomepb"
	"net"

	"google.golang.org/grpc"
)

type server struct{}

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
