package main

import (
	"context"
	"log"
	"net"
	"time"

	pb "github.com/mnao305/grpc-web-chat/messenger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
)

const (
	port = ":9090"
)

type server struct {
	pb.UnimplementedMessengerServer
	requests []*pb.MessageResponse
}

func (s *server) GetMessages(_ *emptypb.Empty, stream pb.Messenger_GetMessagesServer) error {
	for _, r := range s.requests {
		if err := stream.Send(r); err != nil {
			return err
		}
	}
	previousCount := len(s.requests)

	for {
		currentCount := len(s.requests)
		if previousCount < currentCount {
			r := s.requests[currentCount-1]
			log.Printf("Sent: %v", r)
			if err := stream.Send(r); err != nil {
				return err
			}
		}
		previousCount = currentCount
	}
}

func (s *server) CreateMessage(ctx context.Context, r *pb.MessageRequest) (*pb.MessageResponse, error) {
	log.Printf("Received: %v", r.GetMessage())
	newR := &pb.MessageResponse{Message: r.GetMessage(), Date: time.Now().Format("2006-01-02T15:04:05.000000Z"), Name: r.GetName()}
	s.requests = append(s.requests, newR)
	log.Printf("date: %v", s.requests)
	return newR, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterMessengerServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
