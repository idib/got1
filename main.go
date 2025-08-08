package main

import (
	"fmt"
	"log"
	"time"

	"github.com/idib/got1/pkg/client"
	got2_client "github.com/idib/got2/pkg/client"
)

func main() {
	// Create a new client
	c := client.NewClient("http://localhost:8080")

	fmt.Println("Sending ping request to server...")

	// Try to ping the server
	response, err := c.Ping()
	if err != nil {
		log.Fatalf("Failed to ping server: %v", err)
	}

	fmt.Printf("Server response: %s\n", response)

	// Example of how to use the client in a loop
	for i := 1; i <= 3; i++ {
		fmt.Printf("Ping attempt %d: ", i)

		response, err := c.Ping()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		} else {
			fmt.Printf("Success: %s\n", response)
		}

		time.Sleep(1 * time.Second)
	}

	_, err = got2_client.NewClient("http://localhost:8080").Ping()
	if err != nil {
		return
	}
}
