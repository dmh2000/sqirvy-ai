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
	_ "embed"
	"flag"
	"fmt"
	"log"
	"os"

	sqirvy "sqirvy-ai/pkg/sqirvy"
)

// bind the embedded prompts

//go:embed prompts/system.md
var systemPrompt string

//go:embed prompts/query.md
var queryPrompt string

//go:embed prompts/plan.md
var planPrompt string

//go:embed prompts/review.md
var reviewPrompt string

//go:embed prompts/code.md
var codePrompt string

//go:embed prompts/scrape.md
var scrapePrompt string

// a map from function names to their corresponding prompts
var prompts = map[string]string{
	"query":  queryPrompt,
	"plan":   planPrompt,
	"review": reviewPrompt,
	"code":   codePrompt,
	"scrape": scrapePrompt,
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// get comand line flags and arguments
	f_h, f_model, f_function, f_temperature, args, err := CliFlags()
	if err != nil {
		flag.Usage()
		log.Fatal(fmt.Errorf("error parsing flags: %v", err))
	}

	if f_h {
		flag.Usage()
		return
	}

	// start with system prompt
	prompt := systemPrompt

	// get the function prompt
	if _, ok := prompts[f_function]; !ok {
		flag.Usage()
		log.Fatal(fmt.Errorf("invalid function: %s", f_function))
	}
	prompt += prompts[f_function]

	// read additional prompt from stdin and flags
	p, err := ReadPrompt(args)
	if err != nil {
		flag.Usage()
		log.Fatal(fmt.Errorf("error reading prompt: %v", err))
	}
	prompt = prompt + "\n\n" + p

	// Use default model if none specified
	model := f_model

	// Get the provider for the model
	provider, err := sqirvy.GetProviderName(model)
	if err != nil {
		log.Fatal(fmt.Errorf("error getting provider for model %s: %v", model, err))
	}

	// Create client for the provider
	client, err := sqirvy.NewClient(sqirvy.Provider(provider))
	if err != nil {
		log.Fatal(fmt.Errorf("error creating client for provider %s: %v", provider, err))
	}

	// Make the query
	options := sqirvy.Options{Temperature: float32(f_temperature), MaxTokens: sqirvy.MaxTokensDefault}
	response, err := client.QueryText(prompt, model, options)
	if err != nil {
		log.Fatal(fmt.Errorf("error querying model %s: %v", model, err))
	}

	// Print response to stdout
	fmt.Print(response)
	fmt.Println()

	os.Exit(0)
}
