package main

import (
	"fmt"
	"google.golang.org/grpc"
	"grpc/my-simple-grpc/output"
	"log"
	"net"
)

const (
	port = ":8080"
)

func main() {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println("failed to listed", err)
	}

	server := grpc.NewServer()

	output.RegisterOutputServiceServer(server, &OutputServiceServer{})

	log.Printf("Server started at port %s", port)

	if err := server.Serve(listener); err != nil {
		fmt.Println("Failed to serve: ", err)
	}
}

//func main() {
//	lis, err := net.Listen("tcp", port)
//	if err != nil {
//		log.Fatalf("Failed to listen: %v", err)
//	}
//
//	s := grpc.NewServer()
//
//	bs := storage.NewInMemoryBlogStorage()
//
//	blog.RegisterBlogServiceServer(s, &BlogServer{blogStorage: bs})
//
//	log.Printf("Server started at port %s", port)
//
//	if err := s.Serve(lis); err != nil {
//		log.Fatalf("Failed to serve: %v", err)
//	}
