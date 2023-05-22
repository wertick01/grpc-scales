.PHONY: generate

generate:
	protoc --go_out=. --go-grpc_out=. stream/stream.proto