package main

import (
	pb "github.com/kucherenkovova/marco-polo/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	address = "localhost:50051"
)

func main() {
	log.Println("client start")
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewAdapterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := c.Forward(ctx, &pb.Phrase{Word: "marco"})
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	log.Printf("%s", res)
}
