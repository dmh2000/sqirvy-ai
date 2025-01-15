package main

import (
	"fmt"
	"log"
	"os"

	api "sqirvyllm/pkg/api"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	prompt, modelFlag, err := processCommandLine()
	if err != nil {
		helpMessage()
		log.Fatal(err)
	}
	if prompt == "" {
		log.Fatal("no prompt provided")
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

	os.Exit(0)
}
