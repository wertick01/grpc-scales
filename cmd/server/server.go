package main

import (
	"context"

	"github.com/wertick01/grpc-scales/cmd/server/app"
	"github.com/wertick01/grpc-scales/cmd/server/config"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	config, err := config.GetConfig()
	if err != nil {
		panic(err)
	}

	server := app.NewServer(config)
	server.Run(ctx)
}
