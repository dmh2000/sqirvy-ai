package main

import (
	"fmt"
	"log"
	sqirvy "sqirvy-llm/pkg/sqirvy"
)

func main() {
	// Create a new Anthropic client
	client, err := sqirvy.NewClient(sqirvy.MetaLlama)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Make the query
	response, err := client.QueryText("say hello world", "meta-llama/meta-llama-3.1-8b-instruct-turbo", sqirvy.Options{Temperature: 50})
	if err != nil {
		log.Fatalf("Query failed: %v", err)
	}

	fmt.Println("Response:", response)
}
