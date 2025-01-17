// Package main implements a command-line tool for web scraping and AI analysis.
//
// The program accepts URLs and files as input, scrapes web content, and sends
// the combined content to an AI model for analysis. It supports multiple AI
// providers and models, with configurable options for processing and output.
package main

import (
	"fmt"
	"log"
	api "sqirvyllm/pkg/api"
)

func main() {
	// Process command line arguments
	prompt, modelFlag, err := processCommandLine()
	if err != nil {
		log.Fatal(err)
	}
	if prompt == "help" {
		return
	}

	// Use default model if none specified
	model := "claude-3-5-sonnet-latest"
	if modelFlag != "" {
		model = modelFlag
	}

	// Get the provider for the model
	provider, err := api.GetProviderName(model)
	if err != nil {
		log.Fatal(err)
	}

	// Create client for the provider
	client, err := api.NewClient(api.Provider(provider))
	if err != nil {
		log.Fatal(err)
	}

	// Make the query
	response, err := client.QueryText(prompt, model, api.Options{})
	if err != nil {
		log.Fatal(err)
	}

	// Print response to stdout
	fmt.Print(response)
	fmt.Println()
}
