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
	user := welcomepb.UserInfo{
		Name:    "Armin",
		Country: "Iran",
		Age:     21,
	}
	response, err := client.Welcome(user)
	if err != nil {
		panic(err)
	}
	fmt.Println("response:", response)
	err = client.GetGreetings(user)
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
