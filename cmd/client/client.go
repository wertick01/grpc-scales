package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/wertick01/grpc-scales/stream"
	"google.golang.org/grpc"
)

func main() {
	address := ":6060"
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to %s: %v", address, err)
	}
	defer conn.Close()

	client := stream.NewApiCallerScaleClient(conn)

	// Пример обмена данными по протоколу RS-232
	resp, err := client.SetTare(context.Background(), &stream.Empty{})
	if err != nil {
		log.Fatalf("Error while calling SetTare: %v", err)
	}
	if resp.Error != "" {
		fmt.Println("Error:", resp.Error)
	} else {
		fmt.Println("Tare was set")
	}

	// Пример обмена данными по протоколу Ethernet или Wi-Fi
	resp, err = client.ScalesMessage(context.Background(), &stream.RequestScale{
		Message: "Message from client to server through Ethernet or Wi-Fi",
		Type:    "Ethernet or Wi-Fi",
		Subtype: "Scales",
	})
	if err != nil {
		log.Fatalf("Error while calling ScalesMessage: %v", err)
	}
	fmt.Printf("Response from server: %+v\n", resp)

	// Пример обмена данными по протоколу USB со stream из 10 запросов
	stream, err := client.ScalesMessageOutChannel(context.Background(), &stream.RequestScale{})
	if err != nil {
		log.Fatalf("Error while calling ScalesMessageOutChannel: %v", err)
	}
	for i := 1; i <= 10; i++ {
		fmt.Println("Sending message:", i)
		if err := stream.Send(&stream.RequestScale{Message: fmt.Sprintf("Message %d from client to server through USB", i)}); err != nil {
			log.Fatalf("Error while sending request to ScalesMessageOutChannel: %v", err)
		}
		time.Sleep(time.Second)
		resp, err := stream.Recv()
		if err != nil {
			log.Fatalf("Error while receiving response from ScalesMessageOutChannel: %v", err)
		}
		fmt.Printf("Response from server: %+v\n", resp)
	}
	if err := stream.CloseSend(); err != nil {
		log.Fatalf("Error while closing stream: %v", err)
	}
}
