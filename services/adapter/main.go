package main

import (
	"github.com/kucherenkovova/marco-polo/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type MarcoPoloAdapter struct {
	Dict         map[string]string
	ServerClient proto.ServerClient
}

const (
	port          = ":50051"
	serverAddress = "localhost:50052"
)

func (a *MarcoPoloAdapter) Forward(ctx context.Context, phrase *proto.Phrase) (*proto.Phrase, error) {
	res, err := a.ServerClient.Send(ctx, &proto.Phrase{Word: a.Dict[phrase.Word]})
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	return &proto.Phrase{Word: a.Dict[res.Word]}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	// client init
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := proto.NewServerClient(conn)

	// client init end

	a := &MarcoPoloAdapter{
		Dict: map[string]string{
			"marco":  "monkey",
			"polo":   "follow",
			"monkey": "marco",
			"follow": "polo",
		},
		ServerClient: c,
	}

	proto.RegisterAdapterServer(s, a)
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
