package main

import (
	"context"
	"fmt"
	"log"
	sqirvy "sqirvy-ai/pkg/sqirvy"
)

const assistant = "you are a helpful assistant"

func main() {
	// Create a new OpenAI client
	client, err := sqirvy.NewClient(sqirvy.OpenAI)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Make the query
	response, err := client.QueryText(context.Background(), assistant, []string{"say hello world"}, "gpt-4-turbo", sqirvy.Options{MaxTokens: 4096})
	if err != nil {
		log.Fatalf("Query failed: %v", err)
	}

	fmt.Println("Response:", response)
}
