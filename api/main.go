package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"google.golang.org/grpc"

	pb "../proto"
)

var greetSvcClient pb.GreetServiceClient

func main() {
	log.Printf("Starting api...")

	conn, err := grpc.Dial("localhost:8081", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	greetSvcClient = pb.NewGreetServiceClient(conn)

	http.HandleFunc("/hello", sayHello)
	if err := http.ListenAndServe(":8080", http.DefaultServeMux); err != nil {
		log.Fatal(err)
	}

	log.Printf("Stopped api.")
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	resp, err := greetSvcClient.Greet(context.Background(), &pb.GreetRequest{})
	if err != nil {
		log.Printf("Error encountered: %v", err)
		fmt.Fprintf(w, "Error encountered: %v", err)
		return
	}

	// NEXT: get name from customer service
	fmt.Fprintf(w, resp.Message)
}
