package main

import (
	"fmt"
	"log"

	api "sqirvyllm/pkg/api"
)

func main() {
	prompt, modelFlag, err := processCommandLine()
	if err != nil {
		log.Fatal(err)
	}

	// Use default model if none specified
	model := "claude-3.5-sonnet"
	if modelFlag != "" {
		model = modelFlag
	}

	// Get the provider for the model
	provider, err := getProviderName(model)
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
}
