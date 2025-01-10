package main

import (
	"fmt"
	"log"

	api "sqirvyllm/pkg/api"
)

func main() {
	// Create a new OpenAI client
	client, err := api.NewClient(api.OpenAI)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Make the query
	response, err := client.QueryText("say hello world", "gpt-4-turbo-preview", api.Options{})
	if err != nil {
		log.Fatalf("Query failed: %v", err)
	}

	fmt.Println("Response:", response)
}
