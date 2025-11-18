package main

import (
	"fmt"
	"instargam/protogen/userpb"
	userservice "instargam/user-service"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	grpcServer := grpc.NewServer()
	userpb.RegisterUserServiceServer(grpcServer, &userservice.UserServiceServer{})

	fmt.Print("test123")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Error: %v", err)
	}
}
