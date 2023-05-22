FROM golang:1.20-alpine 

ADD . /go/src/github.com/wertick01/grpc-scales

RUN go install github.com/wertick01/grpc-scales@latest

ENTRYPOINT ["/go/bin/server"]

EXPOSE 6060