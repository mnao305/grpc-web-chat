FROM golang:1.15.6

WORKDIR /go/src/grpc-web-chat

COPY go.mod .
COPY go.sum .

RUN go mod download

CMD [ "go", "run", "/go/src/grpc-web-chat/main.go" ]