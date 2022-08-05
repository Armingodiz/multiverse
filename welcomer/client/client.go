package main

import (
	"context"
	"fmt"
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
	response, err := client.Welcome(welcomepb.UserInfo{
		Name:    "Armin",
		Country: "Iran",
		Age:     21,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("response:", response)
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
