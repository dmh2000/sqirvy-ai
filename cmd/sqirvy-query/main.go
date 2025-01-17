// Package main implements a command-line interface for querying AI language models.
//
// The program accepts input from both files and standard input, concatenates them,
// and sends the combined text as a prompt to the specified AI model. It supports
// multiple AI providers including Anthropic, OpenAI, and Google's Gemini.
//
// Usage:
//
//	sqirvy-query [options] files...
//
// The program will read from stdin if available, and concatenate any specified files.
// A system prompt can be provided via system.md in the current directory.
package main

import (
	"fmt"
	"log"
	"os"
	sqirvy "sqirvyllm/pkg/sqirvy"
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
	provider, err := sqirvy.GetProviderName(model)
	if err != nil {
		log.Fatal(err)
	}

	// Create client for the provider
	client, err := sqirvy.NewClient(sqirvy.Provider(provider))
	if err != nil {
		log.Fatal(err)
	}

	// Make the query
	response, err := client.QueryText(prompt, model, sqirvy.Options{})
	if err != nil {
		log.Fatal(err)
	}

	// Print response to stdout
	fmt.Print(response)
	fmt.Println()

	os.Exit(0)
}
