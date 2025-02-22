/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	_ "embed"
	"fmt"

	sqirvy "sqirvy-ai/pkg/sqirvy"

	"github.com/spf13/cobra"
)

// codeCmd represents the code command
func executeQuery(cmd *cobra.Command, sysprompt string, args []string) (string, error) {
	model, err := cmd.Flags().GetString("model")
	if err != nil {
		return "", fmt.Errorf("error getting model: %v", err)
	}
	temperature, err := cmd.Flags().GetInt("temperature")
	if err != nil {
		return "", fmt.Errorf("error getting temperature: %v", err)
	}

	prompt, err := ReadPrompt(sysprompt, args)
	if err != nil {
		return "", fmt.Errorf("error reading prompt:[]string{\n%v", err)
	}
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

	// Make the query
	options := sqirvy.Options{Temperature: float32(temperature), MaxTokens: sqirvy.GetMaxTokens(model)}
	ctx := context.Background()
	response, err := client.QueryText(ctx, []string{prompt}, model, options)
	if err != nil {
		return "", fmt.Errorf("error querying model %s: %v", model, err)
	}

	return response, nil

}
