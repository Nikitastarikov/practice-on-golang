package main

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	reversev1 "github.com/Nikitastarikov/practice-on-golang/grpc_micro_services/proto"
)

func main() {
	fmt.Println("Hello, server!")

	listener, err := net.Listen("tcp", ":5300")

	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)

	reversev1.RegisterReverseServer(grpcServer, &MyServer{})
	grpcServer.Serve(listener)
}

type MyServer struct {
	reversev1.UnimplementedReverseServer
}

func (s *MyServer) Do(c context.Context, request *reversev1.Request) (response *reversev1.Response, err error) {
	n := 0
	// Сreate an array of runes to safely reverse a string.
	r1 := make([]rune, len(request.Message))

	for _, r := range request.Message {
		r1[n] = r
		n++
	}

	// Reverse using runes.
	r1 = r1[0:n]

	for i := 0; i < n/2; i++ {
		r1[i], r1[n-1-i] = r1[n-1-i], r1[i]
	}

	output := string(r1)
	response = &reversev1.Response{
		Message: output,
	}

	return response, nil
}
