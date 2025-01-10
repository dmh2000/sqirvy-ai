package main

import (
	"fmt"
	"log"
	"os"

	"sqirvyllm/pkg/api"
)

func main() {
	// Create a new Anthropic client
	client, err := api.NewClient(api.Anthropic)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Make the query
	response, err := client.QueryText("say hello world", "claude-3-sonnet-20240229", api.Options{})
	if err != nil {
		log.Fatalf("Query failed: %v", err)
	}

	fmt.Println("Response:", response)
}
