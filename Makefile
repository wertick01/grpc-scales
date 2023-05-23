.PHONY: generate build-server

generate:
	protoc --go_out=. --go-grpc_out=. stream/stream.proto

build-server:
	sudo docker build -t grpc-scales .

build-client:
	cd client && go build -o reverse