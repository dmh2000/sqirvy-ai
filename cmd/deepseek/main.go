package main

import (
	"fmt"
	"log"
	sqirvy "sqirvyllm/pkg/sqirvy"
)

func main() {
	// Create a new DeepSeek client
	client, err := sqirvy.NewClient(sqirvy.DeepSeek)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Make the query
	response, err := client.QueryText("say hello world", "deepseek-chat", sqirvy.Options{})
	if err != nil {
		log.Fatalf("Query failed: %v", err)
	}

	fmt.Println("Response:", response)
}
