package main

import (
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"steram.com/stream"
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
