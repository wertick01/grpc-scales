package main

import (
	"context"
	"fmt"
	"log"

	"github.com/wertick01/grpc-scales/client/protocols"
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
	tcpRes, err := client.SetTare(context.Background(), &stream.Empty{})
	log.Println(tcpRes, err)

	port, err := protocols.CreateSerialPort()
	// if err != nil {
	// 	panic(err)
	// }
	log.Println(port, err)
	comResp, err := port.SendRequestScale("", "", "")
	log.Println(comResp, err)

	rsPort := protocols.NewRS232()
	log.Println(rsPort)
	rsResp, err := rsPort.SendRequestScale("", "", "")
	log.Println(rsResp, err)

	req := &stream.RequestScale{
		Message: "Get State",
		Type:    "State",
		Subtype: "",
	}
	response, err := client.ScalesMessageOutChannel(context.Background(), req)
	if err != nil {
		panic(err)
	}

	// Обработка ответа
	for {
		res, err := response.Recv()
		if err != nil {
			panic(err)
		}
		if res.Type == "State" {
			fmt.Printf("Scales state: %s\n", res.Message)
			break
		}
	}
}
