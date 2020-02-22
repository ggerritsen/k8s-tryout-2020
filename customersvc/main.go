package main

import (
	"context"
	"log"
	"math/rand"
	"net"

	"google.golang.org/grpc"

	pb "github.com/ggerritsen/k8s-tryout-2020/proto"
)

func main() {
	log.Print("Starting customerservice...")

	l, err := net.Listen("tcp", ":8082")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterCustomerServiceServer(s, &server{})
	if err := s.Serve(l); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	log.Print("Stopped customerservice...")
}

type server struct {
	pb.UnimplementedGreetServiceServer
}

var db = []*pb.Customer{
	{FirstName: "Kid", LastName: "Danger"},
	{FirstName: "Captain", LastName: "Man"},
	{FirstName: "Bert", LastName: "van Sesamstraat"},
	{FirstName: "Ernie", LastName: "van Sesamstraat"},
	{FirstName: "Dirk", LastName: "Scheele"},
}

func (s *server) GetCustomer(ctx context.Context, in *pb.GetCustomerRequest) (*pb.GetCustomerResponse, error) {
	log.Printf("Got getCustomerRequest: %v", in)

	id := rand.Int() % len(db)
	c := db[id]
	c.Id = int32(id) + 1

	return &pb.GetCustomerResponse{Customer: c}, nil
}
