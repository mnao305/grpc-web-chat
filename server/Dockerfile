FROM golang:1.15.6

WORKDIR /go/src/grpc-web-chat

RUN gom mod download

CMD [ "go", "run", "/go/src/grpc-web-chat/main.go" ]