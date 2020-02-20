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
var customerSvcClient pb.CustomerServiceClient

func main() {
	log.Printf("Starting api...")

	conn, err := grpc.Dial("localhost:8081", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	greetSvcClient = pb.NewGreetServiceClient(conn)

	conn2, err := grpc.Dial("localhost:8082", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn2.Close()
	customerSvcClient = pb.NewCustomerServiceClient(conn2)

	http.HandleFunc("/hello", sayHello)
	if err := http.ListenAndServe(":8080", http.DefaultServeMux); err != nil {
		log.Fatal(err)
	}

	log.Printf("Stopped api.")
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	greeting, err := greetSvcClient.Greet(context.Background(), &pb.GreetRequest{})
	if err != nil {
		log.Printf("Error encountered: %v", err)
		fmt.Fprintf(w, "Error encountered: %v", err)
		return
	}

	c, err := customerSvcClient.GetCustomer(context.Background(), &pb.GetCustomerRequest{})
	if err != nil {
		log.Printf("Error encountered: %v", err)
		fmt.Fprintf(w, "Error encountered: %v", err)
		return
	}

	fmt.Fprintf(w, "<h1>%s, %s %s (id: %d)</h1>", greeting.Message, c.Customer.FirstName, c.Customer.LastName, c.Customer.Id)
}
