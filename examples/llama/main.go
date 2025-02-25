package main

import (
	"context"
	"fmt"
	"log"
	sqirvy "sqirvy-ai/pkg/sqirvy"
)

const assistant = "you are a helpful assistant"

func main() {
	// Create a new Anthropic client
	client, err := sqirvy.NewClient(sqirvy.Llama)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Make the query
	response, err := client.QueryText(context.Background(), assistant, []string{"say hello world"}, "llama3.3-70b", sqirvy.Options{Temperature: 50})
	if err != nil {
		log.Fatalf("Query failed: %v", err)
	}

	fmt.Println("Response:", response)
}
