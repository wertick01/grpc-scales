package main

import (
	"net"

	"github.com/wertick01/grpc-scales/cmd/server/config"
	"github.com/wertick01/grpc-scales/stream"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	config, err := config.GetConfig()
	if err != nil {
		panic(err)
	}

	listen, err := net.Listen("tcp", config.ServerProt)
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	stream.RegisterApiCallerScaleServer(s, &stream.ImplementedApiCallerScaleServer{})
	reflection.Register(s)
	if err := s.Serve(listen); err != nil {
		panic(err)
	}
}
