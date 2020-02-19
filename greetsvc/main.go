package main

import (
	"context"
	"log"
	"math/rand"
	"net"
	"time"

	"google.golang.org/grpc"

	pb "../proto"
)

func main() {
	log.Print("Starting greetservice...")

	rand.Seed(time.Now().Unix())
	
	l, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGreetServiceServer(s, &server{})
	if err := s.Serve(l); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	log.Print("Stopped greetservice.")
}

type server struct {
	pb.UnimplementedGreetServiceServer
}

var greetings = []string{"Sayonara", "Hello", "Bonjour", "Hallo", "Guten Tag"}

func (s *server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Got greetRequest")
	gr := greetings[rand.Int() % len(greetings)]
	return &pb.GreetResponse{Message: gr}, nil
}
