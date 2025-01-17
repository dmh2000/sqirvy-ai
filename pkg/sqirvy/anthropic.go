// Package api provides integration with Anthropic's Claude AI models.
//
// This file implements the Client interface for Anthropic's API, supporting
// both text and JSON queries to Claude models. It handles authentication,
// request formatting, and response parsing specific to Anthropic's requirements.
package api

import (
	"context"
	"fmt"

	"github.com/anthropics/anthropic-sdk-go"
)

const MAX_TOKENS = 2048

// AnthropicClient implements the Client interface for Anthropic's API
type AnthropicClient struct {
	client *anthropic.Client // Anthropic API client
}

func (c *AnthropicClient) QueryText(prompt string, model string, options Options) (string, error) {
	if prompt == "" {
		return "", fmt.Errorf("prompt cannot be empty for text query")
	}

	// Initialize client if not already done
	if c.client == nil {
		c.client = anthropic.NewClient()
	}

	// Build response string
	answer := ""

	// Create new message request with the provided prompt
	message, err := c.client.Messages.New(context.TODO(), anthropic.MessageNewParams{
		Model:     anthropic.F(model),       // Specify which model to use
		MaxTokens: anthropic.F(int64(1024)), // Limit response length
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
		answer += fmt.Sprintf("%v", content.Text)
	}
	return answer, nil
}

func (c *AnthropicClient) QueryJSON(prompt string, model string, options Options) (string, error) {
	if prompt == "" {
		return "", fmt.Errorf("prompt cannot be empty for json query")
	}

	// Initialize client if not already done
	if c.client == nil {
		c.client = anthropic.NewClient()
	}

	// Create new message request expecting JSON response
	message, err := c.client.Messages.New(context.TODO(), anthropic.MessageNewParams{
		Model:     anthropic.F(model),             // Specify which model to use
		MaxTokens: anthropic.F(int64(MAX_TOKENS)), // Use maximum allowed tokens
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

	// Check for valid JSON text content
	if len(message.Content[0].Text) == 0 {
		return "", fmt.Errorf("no JSON content in response")
	}

	// Validate JSON structure
	if message.Content[0].JSON.Text.IsInvalid() {
		return "", fmt.Errorf("invalid JSON content in response")
	}

	// Extract raw JSON text
	r := message.Content[0].JSON.Text.Raw()
	return r, nil
}
