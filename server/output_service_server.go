package main

import (
	"context"
	"fmt"
	"google.golang.org/protobuf/types/known/emptypb"
	"grpc/my-simple-grpc/output"
)

type OutputServiceServer struct {
	output.UnimplementedOutputServiceServer
}

func (o OutputServiceServer) SendOutput(ctx context.Context, req *output.OutputRequest) (*emptypb.Empty, error) {
	fmt.Println(req.Title)
	return &emptypb.Empty{}, nil
}
