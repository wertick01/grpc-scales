FROM golang:1.20-alpine 

ADD . /go/src/github.com/wertick01/grpc-scales/server

RUN apk update && apk add git \
    && GIT_TERMINAL_PROMPT=1 go install github.com/wertick01/grpc-scales/server@latest

ENTRYPOINT ["/go/bin/server"]

EXPOSE 6060