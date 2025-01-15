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
	c.ctx = context.Background()
	c.client, err = genai.NewClient(c.ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	if err != nil {
		return fmt.Errorf("failed to create client: %v", err)
	}
	return nil
}

func (c *GeminiClient) QueryText(prompt string, model string, options Options) (string, error) {
	if err := c.initClient(); err != nil {
		return "", err
	}
	defer c.client.Close()

	genModel := c.client.GenerativeModel(model)
	genModel.ResponseMIMEType = "text/plain"

	resp, err := genModel.GenerateContent(c.ctx, genai.Text(prompt))
	if err != nil {
		return "", fmt.Errorf("failed to generate content: %v", err)
	}

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

func (c *GeminiClient) QueryJSON(prompt string, model string, options Options) (string, error) {
	if err := c.initClient(); err != nil {
		return "", err
	}
	defer c.client.Close()

	genModel := c.client.GenerativeModel(model)
	genModel.ResponseMIMEType = "application/json"

	resp, err := genModel.GenerateContent(c.ctx, genai.Text(prompt))
	if err != nil {
		return "", fmt.Errorf("failed to generate content: %v", err)
	}

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
