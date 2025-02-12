// Package api provides integration with Anthropic's Claude AI models.
//
// This file implements the Client interface for Anthropic's API, supporting
// both text and JSON queries to Claude models. It handles authentication,
// request formatting, and response parsing specific to Anthropic's requirements.
package sqirvy

import (
	"context"
	"fmt"
	"os"

	"github.com/anthropics/anthropic-sdk-go"
)

// AnthropicClient implements the Client interface for Anthropic's API
type AnthropicClient struct {
	client *anthropic.Client // Anthropic API client
}

func (c *AnthropicClient) QueryText(prompt string, model string, options Options) (string, error) {
	// test for ANTHROPIC_API_KEY in environment
	if os.Getenv("ANTHROPIC_API_KEY") == "" {
		return "", fmt.Errorf("ANTHROPIC_API_KEY environment variable not set")
	}

	if prompt == "" {
		return "", fmt.Errorf("prompt cannot be empty for text query")
	}

	// Initialize client if not already done
	if c.client == nil {
		c.client = anthropic.NewClient()
	}

	// set default for initial value
	if options.Temperature < 0.0 {
		options.Temperature = 0.0
	}
	if options.Temperature > 100.0 {
		return "", fmt.Errorf("temperature must be between 1 and 100")
	}
	// scale temperature for Claude 0..1.0
	options.Temperature /= 100.0

	// Build response string
	answer := ""

	// Set default max tokens if not specified
	maxTokens := options.MaxTokens
	if maxTokens == 0 {
		maxTokens = MaxTokensDefault
	}

	// Create new message request with the provided prompt and temperature
	message, err := c.client.Messages.New(context.TODO(), anthropic.MessageNewParams{
		Model:       anthropic.F(model),                        // Specify which model to use
		MaxTokens:   anthropic.F(maxTokens),                    // Limit response length
		Temperature: anthropic.F(float64(options.Temperature)), // Set temperature
		Messages: anthropic.F([]anthropic.MessageParam{
			anthropic.NewUserMessage(anthropic.NewTextBlock(prompt)), // Create user message
		}),
	})
	if err != nil {
		return "", err
	}

	// Verify we got a non-empty response
	if len(message.Content) == 0 {
		return "", fmt.Errorf("no content in response")
	}

	// Concatenate all text blocks from the response
	for _, content := range message.Content {
		answer += fmt.Sprintf("%v", string(content.Text))
	}
	return answer, nil
}
