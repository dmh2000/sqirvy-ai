// Package main implements a code review tool that uses AI models to review code.
package main

import (
	_ "embed"
	"fmt"
	"log"
	api "sqirvyllm/pkg/api"
)

// build the system prompt and review requirements into the binary

//go:embed system.md
var systemPrompt string

//go:embed review.md
var reviewPrompt string

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	prompt, modelFlag, err := processCommandLine()
	if err != nil {
		helpMessage("Error:" + err.Error())
		log.Fatal(err)
	}
	// -h ?
	if prompt == "help" {
		return
	}

	if prompt == "" {
		log.Fatal("no prompt provided")
	}

	// prepend the system prompt and review instructions
	prompt = systemPrompt + "\n\n" + reviewPrompt + "\n\n" + prompt

	// Use default model if none specified
	model := "gemini-1.5-flash"
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
}
