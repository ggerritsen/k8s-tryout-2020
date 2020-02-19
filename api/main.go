package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.Printf("Start")

	http.HandleFunc("/hello", sayHello)
	if err := http.ListenAndServe(":8080", http.DefaultServeMux); err != nil {
		log.Fatal(err)
	}

	log.Printf("Done")
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	if _, err := fmt.Fprint(w, "Hello, world"); err != nil {
		log.Printf("Could not write response: %v", err)
	}
}
