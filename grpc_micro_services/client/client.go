package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	reversev1 "github.com/Nikitastarikov/practice-on-golang/grpc_micro_services/proto"
)

func main() {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	//args := os.Args
	conn, err := grpc.Dial("127.0.0.1:5300", opts...)

	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}

	defer conn.Close()

	client := reversev1.NewReverseClient(conn)
	request := &reversev1.Request{
		Message: "this is a test",
	}
	response, err := client.Do(context.Background(), request)

	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}

	fmt.Println(response.Message)
}
