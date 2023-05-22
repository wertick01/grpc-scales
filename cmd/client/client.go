package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"steram.com/stream"
)

func main() {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	conn, err := grpc.Dial(":6060", opts...)

	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}

	defer conn.Close()

	client := stream.NewApiCallerScaleClient(conn)
	res, err := client.SetTare(context.Background(), &stream.Empty{})

	log.Println(res, err)
}
