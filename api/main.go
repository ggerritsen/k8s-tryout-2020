package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"google.golang.org/grpc"

	pb "github.com/ggerritsen/k8s-tryout-2020/proto"
)

var greetSvcClient pb.GreetServiceClient
var customerSvcClient pb.CustomerServiceClient

func main() {
	log.Printf("Starting api on http://localhost:8080/...")

	ownPort := "8080"
	if v := os.Getenv("PORT"); v != "" {
		ownPort = v
	}

	customerSvcHostPort := "localhost:8081"
	if v := os.Getenv("CUSTOMERSVC_HOSTPORT"); v != "" {
		customerSvcHostPort = v
	}

	greetSvcHostPort := "localhost:8082"
	if v := os.Getenv("GREETSVC_HOSTPORT"); v != "" {
		greetSvcHostPort = v
	}

	customerSvcClient = pb.NewCustomerServiceClient(dial(customerSvcHostPort))
	greetSvcClient = pb.NewGreetServiceClient(dial(greetSvcHostPort))

	http.HandleFunc("/hello", sayHello)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", ownPort), http.DefaultServeMux); err != nil {
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
