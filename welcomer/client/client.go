package main

import (
	"fmt"
	"multiverse/welcomer/welcomepb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials())) // for now because we are dont have a certificate
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := welcomepb.NewWelcomeServiceClient(conn)
	fmt.Println("client:", client)
}
