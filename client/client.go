package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc/my-simple-grpc/output"
	"time"
)

const address = "localhost:8080"

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		fmt.Println("did not connect: ", err)
	}

	defer conn.Close()
	outputService := output.NewOutputServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	outputRequest := &output.OutputRequest{
		Title: "GOGOGO",
	}
	_, err = outputService.SendOutput(ctx, outputRequest)
	if err != nil {
		fmt.Println("error send output: ", err)
	}

	fmt.Println("success send output: ", outputRequest.Title)
}
