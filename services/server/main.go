package main

import (
	"github.com/kucherenkovova/marco-polo/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

const (
	port = ":50052"
)

type MarcoPoloServer struct {
	Dict map[string]string
}

func (s *MarcoPoloServer) Send(ctx context.Context, phrase *proto.Phrase) (*proto.Phrase, error) {
	return &proto.Phrase{Word: s.Dict[phrase.Word]}, nil
}

func main() {
	log.Println("start server")
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	m := &MarcoPoloServer{
		Dict: map[string]string{
			"monkey": "follow",
			"follow": "monkey",
		},
	}
	proto.RegisterServerServer(s, m)
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
