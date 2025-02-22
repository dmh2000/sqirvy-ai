// Package api provides integration with Google's Gemini AI models.
//
// This file implements the Client interface for Google's Gemini API, supporting
// both text and JSON queries. It handles authentication, request formatting,
// and response parsing specific to the Gemini API requirements.
package sqirvy

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

const (
	// GeminiTempScale is the scaling factor for Gemini's 0-2 temperature range
	GeminiTempScale = 2.0
)

// GeminiClient implements the Client interface for Google's Gemini API.
// It provides methods for querying Google's Gemini language models through
// their official API client.
type GeminiClient struct {
	client *genai.Client   // Google Gemini API client
	ctx    context.Context // Context for API requests
}

// Ensure GeminiClient implements the Client interface
var _ Client = (*GeminiClient)(nil)

// NewGeminiClient creates a new instance of GeminiClient.
// It returns an error if the required GEMINI_API_KEY environment variable is not set.
func NewGeminiClient() (*GeminiClient, error) {
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("GEMINI_API_KEY environment variable not set")
	}

	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %w", err)
	}

	return &GeminiClient{
		client: client,
		ctx:    ctx,
	}, nil
}

func (c *GeminiClient) QueryText(ctx context.Context, prompts []string, model string, options Options) (string, error) {
	if len(prompts) == 0 {
		return "", fmt.Errorf("prompts cannot be empty for text query")
	}

	// Create a generative model instance with the specified model name
	genModel := c.client.GenerativeModel(model)
	// Set response type to plain text
	genModel.ResponseMIMEType = "text/plain"

	// Set default and validate temperature
	if options.Temperature < MinTemperature {
		options.Temperature = MinTemperature
	}
	if options.Temperature > MaxTemperature {
		return "", fmt.Errorf("temperature must be between %.1f and %.1f", MinTemperature, MaxTemperature)
	}
	// Scale temperature for Gemini's 0-2 range
	options.Temperature = (options.Temperature * GeminiTempScale) / MaxTemperature
	genModel.Temperature = &options.Temperature

	// Generate content from the prompt
	resp, err := genModel.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		return "", fmt.Errorf("failed to generate content: %w", err)
	}

	// Build response using strings.Builder for better performance
	var result strings.Builder
	for _, candidate := range resp.Candidates {
		for _, part := range candidate.Content.Parts {
			if textValue, ok := part.(genai.Text); ok {
				result.WriteString(string(textValue))
			}
		}
	}

	return result.String(), nil
}

// Close implements the Close method for the Client interface.
func (c *GeminiClient) Close() error {
	return c.client.Close()
}
