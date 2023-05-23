package protocols

import (
	"context"

	"github.com/tarm/serial"
	"github.com/wertick01/grpc-scales/stream"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
)

type ClientCOMPort struct {
	port *serial.Port
}

// Функция для создания экземпляра COM-порта
func CreateSerialPort() (*ClientCOMPort, error) {
	port, err := serial.OpenPort(&serial.Config{
		Name:     "COM1",
		Baud:     57600,
		Parity:   serial.ParityNone,
		StopBits: serial.StopBits(1),
		Size:     8,
	})
	if err != nil {
		return nil, err
	}
	return &ClientCOMPort{
		port: port,
	}, nil
}

// Функция для отправки запроса на сервер через COM-порт
func (c *ClientCOMPort) SendRequestScale(message string, typ string, subtype string) (*stream.ResponseScale, error) {
	// Формирование запроса
	request := &stream.RequestScale{
		Message: message,
		Type:    typ,
		Subtype: subtype,
	}

	// Создание клиента gRPC
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	// Создание клиента для вызова метода сервера ScalesMessageOutChannel
	client := stream.NewApiCallerScaleClient(conn)
	stream, err := client.ScalesMessageOutChannel(context.Background())
	if err != nil {
		return nil, err
	}

	// Отправка запроса на сервер через COM-порт
	requestBytes, err := proto.Marshal(request)
	if err != nil {
		return nil, err
	}

	_, err = c.port.Write(requestBytes)
	if err != nil {
		return nil, err
	}

	// Получение ответа от сервера через gRPC
	response, err := stream.Recv()
	if err != nil {
		return nil, err
	}

	return response, nil
}

type RS232 struct {
	port *serial.Port
}

func NewRS232() *RS232 {
	conn, err := serial.OpenPort(&serial.Config{
		Name:     "/dev/ttyS1",
		Baud:     4800,
		Parity:   serial.ParityEven,
		StopBits: serial.Stop1,
	})
	if err != nil {
		panic(err)
	}

	return &RS232{
		port: conn,
	}
}

// Функция для отправки запроса на сервер через COM-порт
func (c *RS232) SendRequestScale(message string, typ string, subtype string) (*stream.ResponseScale, error) {
	// Формирование запроса
	request := &stream.RequestScale{
		Message: message,
		Type:    typ,
		Subtype: subtype,
	}

	// Создание клиента gRPC
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	// Создание клиента для вызова метода сервера ScalesMessageOutChannel
	client := stream.NewApiCallerScaleClient(conn)
	stream, err := client.ScalesMessageOutChannel(context.Background())
	if err != nil {
		return nil, err
	}

	// Отправка запроса на сервер через COM-порт
	requestBytes, err := proto.Marshal(request)
	if err != nil {
		return nil, err
	}

	_, err = c.port.Write(requestBytes)
	if err != nil {
		return nil, err
	}

	// Получение ответа от сервера через gRPC
	response, err := stream.Recv()
	if err != nil {
		return nil, err
	}

	return response, nil
}
