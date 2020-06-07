package main

import (
	"log"
	"net"

	"github.com/sabarivig/go/grpc/chat"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":4545")
	if err != nil {
		log.Fatal(err)
	}

	s := chat.Server{}

	grpcServer := grpc.NewServer()

	chat.RegisterChatServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}

}
