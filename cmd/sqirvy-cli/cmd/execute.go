// Package cmd implements the command-line interface commands for the sqirvy-cli tool.
// It provides functionality for executing queries against various AI models and
// handling command-line arguments and flags.
package cmd

import (
	"context"
	_ "embed"
	"fmt"

	sqirvy "sqirvy-ai/pkg/sqirvy"

	"github.com/spf13/cobra"
)

// executeQuery processes and executes an AI model query with the given system prompt and arguments.
// It handles model selection, temperature settings, and communication with the AI provider.
//
// Parameters:
//   - cmd: The Cobra command instance containing parsed flags
//   - sysprompt: The system prompt to provide context to the AI model
//   - args: Additional arguments to be processed as part of the query
//
// Returns:
//   - string: The model's response text
//   - error: Any error encountered during execution
func executeQuery(cmd *cobra.Command, system string, args []string) (string, error) {
	// Extract model name from command flags
	model, err := cmd.Flags().GetString("model")
	if err != nil {
		return "", fmt.Errorf("error getting model: %v", err)
	}
	// Extract temperature setting from command flags
	temperature, err := cmd.Flags().GetInt("temperature")
	if err != nil {
		return "", fmt.Errorf("error getting temperature: %v", err)
	}

	// Process system prompt and arguments into query prompts
	prompts, err := ReadPrompt(args)
	if err != nil {
		return "", fmt.Errorf("error reading prompt:[]string{\n%v", err)
	}
	// Determine the AI provider based on the selected model
	provider, err := sqirvy.GetProviderName(model)
	if err != nil {
		return "", fmt.Errorf("error getting provider for model %s: %v", model, err)
	}

	// Create client for the provider
	client, err := sqirvy.NewClient(sqirvy.Provider(provider))
	if err != nil {
		return "", fmt.Errorf("error creating client for provider %s: %v", provider, err)
	}
	defer client.Close()

	// Configure query options and execute the query
	options := sqirvy.Options{Temperature: float32(temperature), MaxTokens: sqirvy.GetMaxTokens(model)}
	ctx := context.Background()
	response, err := client.QueryText(ctx, system, prompts, model, options)
	if err != nil {
		return "", fmt.Errorf("error querying model %s: %v", model, err)
	}

	return response, nil
}
