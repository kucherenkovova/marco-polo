package main

import (
	"github.com/kucherenkovova/marco-polo/proto"
	"github.com/kucherenkovova/marco-polo/services/adapter/adapters"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

const (
	port          = ":50051"
	serverAddress = "localhost:50052"
)

func main() {
	log.Println("start adapter")
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

	a, err := adapters.NewAdapter(c)
	inLookupAdapter, err := adapters.NewDictInLookupAdapter(a, adapters.Dictionary{
		"marco": "monkey",
		"polo":  "follow",
	})
	inOutLookupAdapter, err := adapters.NewDictOutLookupAdapter(inLookupAdapter, adapters.Dictionary{
		"monkey": "marco",
		"follow": "polo",
	})

	proto.RegisterAdapterServer(s, inOutLookupAdapter)
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
