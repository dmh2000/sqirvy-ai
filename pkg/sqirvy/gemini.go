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
	client *genai.Client // Google Gemini API client
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

	client, err := genai.NewClient(context.Background(), option.WithAPIKey(apiKey))
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %w", err)
	}

	return &GeminiClient{
		client: client,
	}, nil
}

// GeminiClient.QueryText implements the QueryText method for the Client interface.
// It sends a text query to Google's Gemini API and returns the generated text response.
func (c *GeminiClient) QueryText(ctx context.Context, system string, prompts []string, model string, options Options) (string, error) {
	if ctx.Err() != nil {
		return "", fmt.Errorf("request context error %w", ctx.Err())
	}

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

	parts := make([]genai.Part, 0, len(prompts))
	// First prompt is system prompt
	parts = append(parts, genai.Text(system))
	// rest of prompts
	for _, prompt := range prompts {
		parts = append(parts, genai.Text(prompt))
	}

	// Generate content from the prompt
	resp, err := genModel.GenerateContent(ctx, parts...)
	if err != nil {
		return "", fmt.Errorf("failed to generate content: %w", err)
	}

	// Build response using strings.Builder for better performance
	var response strings.Builder
	for _, candidate := range resp.Candidates {
		for _, part := range candidate.Content.Parts {
			if textValue, ok := part.(genai.Text); ok {
				response.WriteString(string(textValue))
			}
		}
	}

	return response.String(), nil
}

// Close implements the Close method for the Client interface.
func (c *GeminiClient) Close() error {
	return c.client.Close()
}
