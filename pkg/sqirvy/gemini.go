// Package api provides integration with Google's Gemini AI models.
//
// This file implements the Client interface for Google's Gemini API, supporting
// both text and JSON queries. It handles authentication, request formatting,
// and response parsing specific to the Gemini API requirements.
package api

import (
	"context"
	"fmt"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

// GeminiClient implements the Client interface for Google's Gemini API
type GeminiClient struct {
	client *genai.Client   // Google Gemini API client
	ctx    context.Context // Context for API requests
}

func (c *GeminiClient) initClient() error {
	var err error
	// Create a background context for API operations
	c.ctx = context.Background()
	// Initialize the Gemini client with API key from environment
	c.client, err = genai.NewClient(c.ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	if err != nil {
		return fmt.Errorf("failed to create client: %w", err)
	}
	return nil
}

func (c *GeminiClient) QueryText(prompt string, model string, options Options) (string, error) {
	if prompt == "" {
		return "", fmt.Errorf("prompt cannot be empty for text query")
	}

	// Initialize the client if not already done
	if err := c.initClient(); err != nil {
		return "", err
	}
	// Ensure client is closed after we're done
	defer c.client.Close()

	// Create a generative model instance with the specified model name
	genModel := c.client.GenerativeModel(model)
	// Set response type to plain text
	genModel.ResponseMIMEType = "text/plain"
	// Set temperature
	if options.Temperature < 0.0 {
		options.Temperature = 0.0
	}
	if options.Temperature > 100.0 {
		return "", fmt.Errorf("temperature must be between 1 and 100")
	}
	// scale temperature for gemini 0..2.0
	options.Temperature = (options.Temperature * 2) / 100.0
	genModel.Temperature = &options.Temperature

	// Generate content from the prompt
	resp, err := genModel.GenerateContent(c.ctx, genai.Text(prompt))
	if err != nil {
		return "", fmt.Errorf("failed to generate content: %w", err)
	}

	// Concatenate all text parts from all candidates into a single string
	var result string
	for _, candidate := range resp.Candidates {
		for _, part := range candidate.Content.Parts {
			if textValue, ok := part.(genai.Text); ok {
				result += string(textValue)
			}
		}
	}

	return result, nil
}
