package main

import (
	"fmt"
	pb "github.com/kucherenkovova/marco-polo/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"os"
	"time"
)

const (
	HostKey = "ADAPTER_HOST"
	PortKey = "ADAPTER_PORT"
)

func main() {
	log.Println("client start")
	host, ok := os.LookupEnv(HostKey)
	if !ok {
		log.Fatalf("Env variable %s not found", HostKey)
	}
	port, ok := os.LookupEnv(PortKey)
	if !ok {
		log.Fatalf("Env variable %s not found", PortKey)
	}
	address := fmt.Sprintf("%s:%s", host, port)

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
