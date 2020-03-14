package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"google.golang.org/grpc"

	pb "github.com/ggerritsen/k8s-tryout-2020/proto"
)

var greetSvcClient pb.GreetServiceClient
var customerSvcClient pb.CustomerServiceClient

func main() {
	log.Printf("Starting api on http://localhost:8080/...")

	greetSvcClient = pb.NewGreetServiceClient(dial("greetsvc:8081"))
	customerSvcClient = pb.NewCustomerServiceClient(dial("customersvc:8082"))

	http.HandleFunc("/hello", sayHello)
	if err := http.ListenAndServe(":8080", http.DefaultServeMux); err != nil {
		log.Fatal(err)
	}

	log.Printf("Stopped api.")
}

func dial(hostPort string) grpc.ClientConnInterface {
	conn, err := grpc.Dial(hostPort, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	return conn
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
