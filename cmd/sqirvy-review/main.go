// Package main implements an automated code review tool that leverages AI language models
// to perform comprehensive code analysis and provide detailed feedback.
//
// The tool analyzes source code for:
// - Potential bugs and logic errors
// - Security vulnerabilities and best practices
// - Performance optimizations
// - Style and idiomatic code usage
// - Documentation and comment quality
//
// It supports multiple AI providers (OpenAI, Anthropic, Google) and uses embedded
// system prompts to ensure consistent and thorough reviews.
//
// Usage:
//
//	sqirvy-review [options] files...
//
// The tool will analyze all provided files and generate a detailed markdown report
// containing findings and recommendations across multiple categories.
package main

import (
	_ "embed"
	"fmt"
	"log"
	sqirvy "sqirvy-llm/pkg/sqirvy"
)

// build the system prompt and review requirements into the binary

//go:embed system.md
var systemPrompt string

//go:embed review.md
var reviewPrompt string

var DEFAULT_MODEL = "gemini-1.5-flash"

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
	model := DEFAULT_MODEL
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
}
