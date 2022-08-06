package main

import (
	"context"
	"fmt"
	"io"
	"multiverse/welcomer/welcomepb"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Client struct {
	grpcClient welcomepb.WelcomeServiceClient
}

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials())) // for now because we are dont have a certificate
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	cli := welcomepb.NewWelcomeServiceClient(conn)
	client := &Client{cli}
	// user := welcomepb.UserInfo{
	// 	Name:    "Armin",
	// 	Country: "Iran",
	// 	Age:     21,
	// }
	// response, err := client.Welcome(user)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("response:", response)
	// err = client.GetGreetings(user)
	// if err != nil {
	// 	if err == io.EOF {
	// 		fmt.Println("All creetings reveived")
	// 	} else {
	// 		panic(err)
	// 	}
	// }
	// err = client.ToManyPeopleComing()
	// if err != nil {
	// 	if err == io.EOF {
	// 		fmt.Println("All creetings reveived At the moment")
	// 	} else {
	// 		panic(err)
	// 	}
	// }
	err = client.ManyPeopleComingAtTheMoment()
	if err != nil {
		if err == io.EOF {
			fmt.Println("All creetings reveived")
		} else {
			panic(err)
		}
	}
}

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

func (client *Client) GetGreetings(user welcomepb.UserInfo) error {
	resStream, err := client.grpcClient.GetGreetings(context.Background(), &welcomepb.WelcomeRequest{
		User:    &user,
		Arrival: timestamppb.New(time.Now()),
	})
	if err != nil {
		return err
	}
	for {
		response, err := resStream.Recv()
		if err != nil {
			return err
		}
		fmt.Println("response:", response)
	}
}

func (client *Client) ToManyPeopleComing() error {
	stream, err := client.grpcClient.ToManyPeopleComing(context.Background())
	if err != nil {
		return err
	}
	for i := 0; i < 10; i++ {
		user := welcomepb.UserInfo{
			Name:    "Armin" + fmt.Sprintf("%d", i),
			Country: "Iran",
			Age:     20 + int32(i),
		}
		if err != nil {
			return err
		}
		time.Sleep(time.Millisecond * 200)
		stream.Send(&welcomepb.WelcomeRequest{
			User:    &user,
			Arrival: timestamppb.New(time.Now()),
		})
	}
	response, err := stream.CloseAndRecv()
	if err != nil {
		return err
	}
	fmt.Println("response of to many people comeing is: ", response)
	return nil
}

func (client *Client) ManyPeopleComingAtTheMoment() error {
	stream, err := client.grpcClient.ManyPeopleComingAtTheMoment(context.Background())
	if err != nil {
		return err
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
		}
	}()
	for {
		select {
		case <-done:
			return nil
		case err := <-errChan:
			return err
		}
	}
}
