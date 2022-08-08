package main

import (
	"context"
	"fmt"
	"io"
	"multiverse/welcomer/welcomepb"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
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

func (s *server) LongWelcome(ctx context.Context, in *welcomepb.WelcomeRequest) (*welcomepb.WelcomeResponse, error) {
	for i := 0; i < 5; i++ {
		// we dont call sleep(5 * seconds) because we want to check the context every second to be able to cancel the request and stop using server recources
		time.Sleep(time.Second)
		if ctx.Err() == context.DeadlineExceeded {
			return nil, status.Errorf(codes.DeadlineExceeded, "deadline exceeded")
		}
	}
	return &welcomepb.WelcomeResponse{Result: "Hello " + in.User.Name + " from: " + in.User.Country + " you came at " + in.GetArrival().String()}, nil
}

// all requaments are met
func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	useSsl := true // TODO make this configurable
	var creds credentials.TransportCredentials
	if useSsl {
		creds, err = credentials.NewServerTLSFromFile("ssl/server.crt", "ssl/server.pem")
		if err != nil {
			panic(err)
		}
	}
	s := grpc.NewServer(grpc.Creds(creds))
	welcomepb.RegisterWelcomeServiceServer(s, &server{})
	if err := s.Serve(listen); err != nil {
		panic(err)
	}
}
