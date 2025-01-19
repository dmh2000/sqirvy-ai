package main

import (
	"fmt"
	"log"
	sqirvy "sqirvyllm/pkg/sqirvy"
)

func main() {
	// Create a new Gemini client
	client, err := sqirvy.NewClient(sqirvy.Gemini)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Make the query
	response, err := client.QueryText("say hello world", "gemini-2.0-flash-exp", sqirvy.Options{})
	if err != nil {
		log.Fatalf("Query failed: %v", err)
	}

	fmt.Println("Response:", response)
}
