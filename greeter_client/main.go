package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
)

const (
	defaultName = "world"
)

var (
	addr1 = flag.String("addr1", "localhost:50051", "the address to connect to for case 1")
	addr2 = flag.String("addr2", "localhost:50052", "the address to connect to for case 2")
	addr3 = flag.String("addr3", "localhost:50053", "the address to connect to for case 3")
	addr4 = flag.String("addr4", "localhost:50054", "the address to connect to for case 4")
	addr5 = flag.String("addr5", "localhost:50055", "the address to connect to for case 4")
	name  = flag.String("name", defaultName, "Name to greet")
)

func main() {
	flag.Parse()

	// Set up an HTTP server to listen for incoming requests.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Parse the "case" field from the request query.
		caseValue := r.URL.Query().Get("case")
		fmt.Printf("Received HTTP request with case=%s\n", caseValue)

		// Choose the gRPC server addresses based on the "case" value.
		var selectedAddresses []string
		var caseName string

		switch caseValue {
		case "1":
			selectedAddresses = []string{*addr1, *addr2, *addr3}
			caseName = "Case 1"
		case "2":
			selectedAddresses = []string{*addr1, *addr2}
			caseName = "Case 2"
		case "3":
			selectedAddresses = []string{*addr1, *addr2, *addr3, *addr4, *addr5}
			caseName = "Case 3"
		case "4":
			selectedAddresses = []string{*addr1, *addr2, *addr3, *addr5}
			caseName = "Case 4"
		default:
			log.Printf("Invalid case value: %s", caseValue)
			http.Error(w, "Invalid case value", http.StatusBadRequest)
			return
		}

		// Create a variable to store the combined gRPC response content.
		var responseContent string

		// Iterate over selectedAddresses and send gRPC requests to each server.
		for _, selectedAddress := range selectedAddresses {
			conn, err := grpc.Dial(selectedAddress, grpc.WithInsecure())
			if err != nil {
				log.Fatalf("did not connect to server: %v", err)
			}
			defer conn.Close()

			client := pb.NewGreeterClient(conn)

			r, err := client.SayHello(context.Background(), &pb.HelloRequest{Name: *name})
			if err != nil {
				log.Fatalf("could not greet %s: %v", selectedAddress, err)
			}

			fmt.Printf("%s Greeting: %s\n", selectedAddress, r.GetMessage())

			// Append gRPC response content to the variable.
			responseContent += r.GetMessage() + "\n"
		}

		// Write the combined responseContent to the HTTP response.
		_, err := fmt.Fprintln(w, responseContent)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		fmt.Printf("Selected %s\n", caseName)
	})

	serverAddr := "localhost:8082" // Set the desired port for the HTTP server.
	server := &http.Server{
		Addr:         serverAddr,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Printf("Starting HTTP server on %s...\n", serverAddr)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("HTTP server error: %v", err)
	}
}
