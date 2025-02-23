package main

import (
	"context"
	"fmt"
	"log"
	sqirvy "sqirvy-ai/pkg/sqirvy"
)

const assistant = "you are a helpful assistant"

func main() {
	// Create a new Gemini client
	client, err := sqirvy.NewClient(sqirvy.Gemini)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Make the query
	response, err := client.QueryText(context.Background(), assistant, []string{"say hello world"}, "gemini-2.0-flash-exp", sqirvy.Options{})
	if err != nil {
		log.Fatalf("Query failed: %v", err)
	}

	fmt.Println("Response:", response)
}
