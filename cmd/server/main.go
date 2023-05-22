package main

import (
	"net"

	"github.com/wertick01/grpc-scales/stream"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

func main() {
	listener, err := net.Listen("tcp", ":6060")
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer([]grpc.ServerOption{})

	stream.RegisterApiCallerScaleServer(server, &stream.UnimplementedApiCallerScaleServer{})

	server.Serve(listener)
}
