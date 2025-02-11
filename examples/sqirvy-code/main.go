// Package main implements an automated code generation tool that leverages AI language models
// to perform comprehensive code analysis and provide detailed feedback.
// //
// It supports multiple AI providers (OpenAI, Anthropic, Google) and uses embedded
// system and generator prompts
//
// Usage:
//
//	sqirvy-code [options] files...
//
// The tool will analyze all provided files and generate a detailed markdown report
// containing findings and recommendations across multiple categories.
package main

import (
	_ "embed"
	"fmt"
	"log"
	sqirvy "sqirvy-ai/pkg/sqirvy"
)

// build the system prompt and code requirements into the binary

//go:embed system.md
var systemPrompt string

//go:embed code.md
var codePrompt string

var DEFAULT_MODEL = "claude-3-5-sonnet-latest"

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

	// prepend the system prompt and codegen instructions
	prompt = systemPrompt + "\n\n" + codePrompt + "\n\n" + prompt

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
