package main

import (
	"context"
	"log"

	"github.com/wertick01/grpc-scales/stream"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
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
