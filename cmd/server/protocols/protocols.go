package protocols

import (
	"context"
	"log"

	"github.com/tarm/serial"
	"github.com/wertick01/grpc-scales/pkg/logger"
	"github.com/wertick01/grpc-scales/stream"
	"google.golang.org/protobuf/proto"
)

type COMPortCaller struct {
	Port   *serial.Port
	logger logger.ILogger
}

func CreateCOMPortCaller(logger logger.ILogger) *COMPortCaller {
	port, err := serial.OpenPort(&serial.Config{
		Name:     "COM1",
		Baud:     57600,
		Parity:   serial.ParityNone,
		StopBits: serial.StopBits(1),
		Size:     8,
	})
	if err != nil {
		panic(err)
	}

	return &COMPortCaller{
		Port:   port,
		logger: logger,
	}
}

// функция чтения данных с COM-порта
func (c *COMPortCaller) ReadFromCOMPort(ctx context.Context, serverCaller stream.ApiCallerScaleServer) {
	bufSize := 1024
	buf := make([]byte, bufSize)

	// читаем данные с COM-порта
	for {
		n, err := c.Port.Read(buf)
		if err != nil {
			c.logger.Error(&logger.Fields{
				Detail: &logger.Detail{
					Backtrace: logger.GetBacktrace(err),
				},
			})
		}

		// парсим полученные данные в protobuf сообщение
		req := &stream.RequestScale{}
		if err := proto.Unmarshal(buf[:n], req); err != nil {
			c.logger.Error(&logger.Fields{
				Detail: &logger.Detail{
					Backtrace: logger.GetBacktrace(err),
				},
			})
		}

		// обрабатываем полученное сообщение
		if req.Type == "set_tare" {
			res, err := serverCaller.SetTare(ctx, &stream.Empty{})
			if err != nil {
				c.logger.Error(&logger.Fields{
					Detail: &logger.Detail{
						Backtrace: logger.GetBacktrace(err),
					},
				})
			}

			requestBytes, err := proto.Marshal(res)
			if err != nil {
				c.logger.Error(&logger.Fields{
					Detail: &logger.Detail{
						Backtrace: logger.GetBacktrace(err),
					},
				})
			}

			_, err = c.Port.Write(requestBytes)
			if err != nil {
				c.logger.Error(&logger.Fields{
					Detail: &logger.Detail{
						Backtrace: logger.GetBacktrace(err),
					},
				})
			}

		} else if req.Type == "get_instant_weight" {
			// обработка команды получения веса
			// ...
		}
		// ...
	}
}

func (c *COMPortCaller) ScalesMessageOutChannel(resp *stream.ResponseScale, strm stream.ApiCallerScale_ScalesMessageOutChannelServer) error {
	if err := strm.Send(resp); err != nil {
		c.logger.Error(&logger.Fields{
			Detail: &logger.Detail{
				Backtrace: logger.GetBacktrace(err),
			},
		})
	}
	return nil
}

type USBCaller struct {
}

func (s *USBCaller) ScalesMessageOutChannel(req *stream.RequestScale, ss stream.ApiCallerScale_ScalesMessageOutChannelServer) error {
	// Handle incoming messages from client
	for {
		select {
		case <-ss.Context().Done():
			log.Printf("Client disconnected")
			return nil
		default:
			// Process messages from USB and send response
			response := stream.ResponseScale{
				Message: "Weight value obtained",
				Type:    "response",
				Subtype: "weight",
			}
			err := ss.Send(&response)
			if err != nil {
				log.Fatalf("Error sending message: %v", err)
			}
		}
	}
}
