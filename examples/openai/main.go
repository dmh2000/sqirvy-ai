package main

import (
	"fmt"
	"log"
	sqirvy "sqirvy-ai/pkg/sqirvy"
)

func main() {
	// Create a new OpenAI client
	client, err := sqirvy.NewClient(sqirvy.OpenAI)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Make the query
	response, err := client.QueryText("say hello world", "gpt-4-turbo-preview", sqirvy.Options{})
	if err != nil {
		log.Fatalf("Query failed: %v", err)
	}

	fmt.Println("Response:", response)
}
