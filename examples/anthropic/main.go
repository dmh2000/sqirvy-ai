package main

import (
	"context"
	"fmt"
	"log"
	sqirvy "sqirvy-ai/pkg/sqirvy"
)

func main() {
	// Create a new Anthropic client
	client, err := sqirvy.NewClient(sqirvy.Anthropic)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Make the query
	response, err := client.QueryText(context.Background(), []string{"say hello world"}, "claude-3-5-sonnet-latest", sqirvy.Options{Temperature: 50})
	if err != nil {
		log.Fatalf("Query failed: %v", err)
	}

	fmt.Println("Response:", response)
}
