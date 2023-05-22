package app

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"

	"github.com/wertick01/grpc-scales/cmd/server/config"
	"github.com/wertick01/grpc-scales/cmd/server/protocols"
	"github.com/wertick01/grpc-scales/pkg/logger"
	"github.com/wertick01/grpc-scales/stream"
	"google.golang.org/grpc"
)

type server struct {
	cfg           *config.Config
	srv           *grpc.Server
	comPortCaller *protocols.COMPortCaller
	logger        logger.ILogger
}

func NewServer(cfg *config.Config) *server {
	server := new(server)

	server.cfg = cfg
	newLogger, err := logger.NewZapLogger(logger.Config{
		Level:   cfg.LogLevel,
		Context: "BMI",
		Version: "2.0.0",
	})
	if err != nil {
		panic(err)
	}
	server.logger = newLogger

	// Я знаю, что логгер надо прокидывать как тут (в закомменченом коде), но решил пока не уделять этому слишком много внимания
	// var grpcCode grpc_zap.CodeToLevel

	// opts := []grpc_zap.Option{
	// 	grpc_zap.WithLevels(grpcCode),
	// }

	// server.srv = grpc.NewServer(grpc_middleware.WithUnaryServerChain(
	// 	grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
	// 	grpc_zap.UnaryServerInterceptor(server.logger, opts...),
	// ))

	server.srv = grpc.NewServer()

	return server
}

func (server *server) Run(ctx context.Context) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	listener, err := net.Listen(server.cfg.NetProtokol, ":"+server.cfg.ServerProt)
	if err != nil {
		server.logger.Panic(&logger.Fields{
			Detail: &logger.Detail{
				Backtrace: logger.GetBacktrace(err),
			},
			Message: "Sorry, listener don't want to listen your requests :(((",
			Code:    "server"})
	}

	server.comPortCaller = protocols.CreateCOMPortCaller(server.logger)
	implemented := &stream.ImplementedApiCallerScaleServer{}

	stream.RegisterApiCallerScaleServer(server.srv, implemented)
	go func() {
		if err := server.Serve(listener); err != nil && err != grpc.ErrServerStopped {
			server.logger.Panic(&logger.Fields{
				Detail: &logger.Detail{
					Backtrace: logger.GetBacktrace(err),
				},
				Message: "Sorry, server was died :(((",
				Code:    "server"})
		}
	}()

	go server.comPortCaller.ReadFromCOMPort(ctx, implemented)

	oscall := <-c

	server.logger.Info(&logger.Fields{
		Message: fmt.Sprintf("System call: +%v", oscall),
		Code:    "server",
	})
	server.Shutdown()
}

func (server *server) Serve(listener net.Listener) error {
	server.logger.Info(&logger.Fields{
		Code:    "server",
		Message: "Server started successfully",
	})

	return server.srv.Serve(listener)
}

func (server *server) Shutdown() {
	server.logger.Info(&logger.Fields{
		Message: "Server stoped",
		Code:    "server",
	})

	// TODO necessary logic

	server.srv.Stop()
	server.comPortCaller.Port.Close()

	// TODO necessary logic

	server.logger.Info(&logger.Fields{
		Message: "server exited properly",
		Code:    "server",
	})
}
