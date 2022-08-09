package client

import (
	"context"
	"fmt"
	"io"
	"log"
	"multiverse/welcomer/welcomepb"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type WelcomerClient interface {
	Welcome(user welcomepb.UserInfo) (*welcomepb.WelcomeResponse, error)
	GetGreetings(user welcomepb.UserInfo) ([]*welcomepb.WelcomeResponse, error)
	ToManyPeopleComing() (res *welcomepb.WelcomeResponse, err error)
	ManyPeopleComingAtTheMoment() (res []*welcomepb.WelcomeResponse, err error)
}

func NewWelcomerClient(useSsl bool, host, port string) (WelcomerClient, error) {
	var creds credentials.TransportCredentials
	var err error
	if useSsl {
		creds, err = credentials.NewClientTLSFromFile("ssl/ca.crt", "") // Certificate Authority Trust certificate
	}
	if err != nil {
		return nil, err
	}
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", host, port), grpc.WithTransportCredentials(creds)) // for now because we are dont have a certificate
	if err != nil {
		return nil, err
	}
	cli := welcomepb.NewWelcomeServiceClient(conn)
	return &Client{cli}, nil
}

type Client struct {
	grpcClient welcomepb.WelcomeServiceClient
}

// func main() {
// 	var creds credentials.TransportCredentials
// 	var err error
// 	useSsl := false // TODO make this configurable
// 	if useSsl {
// 		creds, err = credentials.NewClientTLSFromFile("ssl/ca.crt", "") // Certificate Authority Trust certificate
// 	}
// 	if err != nil {
// 		panic(err)
// 	}
// 	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(creds)) // for now because we are dont have a certificate
// 	if err != nil {
// 		panic(err)
// 	}
// 	cli := welcomepb.NewWelcomeServiceClient(conn)
// 	defer conn.Close()

// 	client, err := &Client{cli}
// 	user := welcomepb.UserInfo{
// 		Name:    "Armin",
// 		Country: "Iran",
// 		Age:     21,
// 	}
// 	// response, err := client.Welcome(user)
// 	// if err != nil {
// 	// 	panic(err)
// 	// }
// 	// fmt.Println("response:", response)
// 	// err = client.GetGreetings(user)
// 	// if err != nil {
// 	// 	if err == io.EOF {
// 	// 		fmt.Println("All welcomes reveived")
// 	// 	} else {
// 	// 		panic(err)
// 	// 	}
// 	// }
// 	// err = client.ToManyPeopleComing()
// 	// if err != nil {
// 	// 	if err == io.EOF {
// 	// 		fmt.Println("All welcomes reveived At the moment")
// 	// 	} else {
// 	// 		panic(err)
// 	// 	}
// 	// }
// 	//err = client.ManyPeopleComingAtTheMoment()
// 	err = client.LongWelcome(user, 6)
// 	if err != nil {
// 		panic(err)
// 	}
// }

func (client *Client) Welcome(user welcomepb.UserInfo) (*welcomepb.WelcomeResponse, error) {
	response, err := client.grpcClient.Welcome(context.Background(), &welcomepb.WelcomeRequest{
		User: &welcomepb.UserInfo{
			Name:    user.Name,
			Country: user.Country,
			Age:     user.Age,
		},
		Arrival: timestamppb.New(time.Now()),
	})
	return response, err
}

func (client *Client) GetGreetings(user welcomepb.UserInfo) ([]*welcomepb.WelcomeResponse, error) {
	responses := make([]*welcomepb.WelcomeResponse, 0)
	resStream, err := client.grpcClient.GetGreetings(context.Background(), &welcomepb.WelcomeRequest{
		User:    &user,
		Arrival: timestamppb.New(time.Now()),
	})
	if err != nil {
		return nil, err
	}
	for {
		response, err := resStream.Recv()
		if err != nil {
			if err == io.EOF {
				return responses, nil
			}
			return nil, err
		}
		responses = append(responses, response)
		fmt.Println("response:", response)
	}
}

func (client *Client) ToManyPeopleComing() (res *welcomepb.WelcomeResponse, err error) {
	stream, err := client.grpcClient.ToManyPeopleComing(context.Background())
	if err != nil {
		return
	}
	for i := 0; i < 10; i++ {
		user := welcomepb.UserInfo{
			Name:    "Armin" + fmt.Sprintf("%d", i),
			Country: "Iran",
			Age:     20 + int32(i),
		}
		if err != nil {
			return
		}
		time.Sleep(time.Millisecond * 200)
		stream.Send(&welcomepb.WelcomeRequest{
			User:    &user,
			Arrival: timestamppb.New(time.Now()),
		})
	}
	response, err := stream.CloseAndRecv()
	if err != nil {
		return
	}
	fmt.Println("response of to many people comeing is: ", response)
	return response, nil
}

func (client *Client) ManyPeopleComingAtTheMoment() (res []*welcomepb.WelcomeResponse, err error) {
	stream, err := client.grpcClient.ManyPeopleComingAtTheMoment(context.Background())
	res = make([]*welcomepb.WelcomeResponse, 0)
	if err != nil {
		return
	}
	done := make(chan bool)
	errChan := make(chan error)
	go func() {
		for i := 0; i < 10; i++ {
			user := welcomepb.UserInfo{
				Name:    "Armin" + fmt.Sprintf("%d", i),
				Country: "Iran",
				Age:     20 + int32(i),
			}
			if err != nil {
				errChan <- err
				return
			}
			time.Sleep(time.Millisecond * 200)
			stream.Send(&welcomepb.WelcomeRequest{
				User:    &user,
				Arrival: timestamppb.New(time.Now()),
			})
		}
		stream.CloseSend()
	}()

	go func() {
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
			fmt.Println("response:", response)
			res = append(res, response)
		}
	}()
	for {
		select {
		case <-done:
			return
		case err := <-errChan:
			return nil, err
		}
	}
}
func (client *Client) LongWelcome(user welcomepb.UserInfo, timeout int) error {
	contextWithTimeout, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(timeout))
	defer cancel()
	response, err := client.grpcClient.LongWelcome(contextWithTimeout, &welcomepb.WelcomeRequest{
		User: &welcomepb.UserInfo{
			Name:    user.Name,
			Country: user.Country,
			Age:     user.Age,
		},
		Arrival: timestamppb.New(time.Now()),
	})
	if err != nil {
		formattedErr, ok := status.FromError(err)
		if ok {
			if formattedErr.Code() == codes.DeadlineExceeded {
				fmt.Println("Timeout")
				return nil
			}
		}
		log.Println("unexpected error: ", formattedErr.Message())
		return err
	}
	fmt.Println("response:", response)
	return nil
}
