package main

import (
	"context"
	"fmt"
	"io"
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
		current := arrival + "passed " + fmt.Sprintf("%d", i) + " times"
		stream.Send(&welcomepb.WelcomeResponse{Result: "Hello " + in.User.Name + " from: " + in.User.Country + " you came at " + current})
		time.Sleep(time.Second)
	}
	return nil
}

func (s *server) ToManyPeopleComing(stream welcomepb.WelcomeService_ToManyPeopleComingServer) error {
	finalRes := ""
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			stream.SendAndClose(&welcomepb.WelcomeResponse{Result: finalRes})
			return nil
		}
		if err != nil {
			return err
		}
		arrival := in.GetArrival().String()
		finalRes += in.User.Name + "from " + in.User.Country + " you came at " + arrival + "\n"
	}
}

func (s *server) ManyPeopleComingAtTheMoment(stream welcomepb.WelcomeService_ManyPeopleComingAtTheMomentServer) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		err = stream.Send(&welcomepb.WelcomeResponse{Result: in.User.Name + "from " + in.User.Country + " you came at " + in.GetArrival().String()})
		if err != nil {
			return err
		}
	}
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
