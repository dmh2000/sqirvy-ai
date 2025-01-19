// Package api provides integration with OpenAI's GPT models.
//
// This file implements the Client interface for OpenAI's API, supporting
// both text and JSON queries to GPT models. It handles authentication,
// request formatting, and response parsing specific to OpenAI's requirements.
package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

var openAIEndpoint = "https://api.openai.com/v1/chat/completions"

// OpenAIClient implements the Client interface for OpenAI's API
type OpenAIClient struct {
	apiKey string       // OpenAI API authentication key
	client *http.Client // HTTP client for making API requests
}

// openAIRequest represents the structure of a request to OpenAI's chat completion API
type openAIRequest struct {
	Model          string          `json:"model"`                           // Model identifier
	Messages       []openAIMessage `json:"messages"`                        // Conversation messages
	MaxTokens      int             `json:"max_completion_tokens,omitempty"` // Max response length
	ResponseFormat string          `json:"response_format,omitempty"`       // Desired response format
}

type openAIMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type openAIResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

func (c *OpenAIClient) QueryText(prompt string, model string, options Options) (string, error) {
	if prompt == "" {
		return "", fmt.Errorf("prompt cannot be empty for text query")
	}

	// update the endpoing if OPENAI_API_BASE is set
	if base := os.Getenv("OPENAI_API_BASE"); base != "" {
		openAIEndpoint = base + "/v1/chat/completions"
	}

	// Initialize HTTP client and API key if not already done
	if c.client == nil {
		c.client = &http.Client{}
		// Get API key from environment variable
		c.apiKey = os.Getenv("OPENAI_API_KEY")
		if c.apiKey == "" {
			return "", fmt.Errorf("OPENAI_API_KEY environment variable not set")
		}
	}

	// Construct the request body with the prompt as a user message
	reqBody := openAIRequest{
		Model: model,
		Messages: []openAIMessage{
			{Role: "user", Content: prompt},
		},
		MaxTokens: 1024, // Limit response length
	}

	// Send request and return response
	return c.makeRequest(reqBody)
}

// QueryJSON sends a JSON query to OpenAI and returns the response
// using json has some options, see:
// https://platform.openai.com/docs/guides/structured-outputs#examples
func (c *OpenAIClient) QueryJSON(prompt string, model string, options Options) (string, error) {
	// Validate prompt is not empty
	if prompt == "" {
		return "", fmt.Errorf("prompt cannot be empty for json query")
	}

	// Initialize HTTP client and API key if not already done
	if c.client == nil {
		c.client = &http.Client{}
		// Get API key from environment variable
		c.apiKey = os.Getenv("OPENAI_API_KEY")
		if c.apiKey == "" {
			return "", fmt.Errorf("OPENAI_API_KEY environment variable not set")
		}
	}

	// Construct the request body with the prompt as a user message
	reqBody := openAIRequest{
		Model: model,
		Messages: []openAIMessage{
			{Role: "user", Content: prompt},
		},
		MaxTokens: 1024, // Limit response length
	}

	// Send request and return response
	return c.makeRequest(reqBody)
}

func (c *OpenAIClient) makeRequest(reqBody openAIRequest) (string, error) {
	// Convert request body to JSON
	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request: %w", err)
	}

	// Create new HTTP request with JSON body
	req, err := http.NewRequest("POST", openAIEndpoint, bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	// Set required headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.apiKey)

	// Send the request
	resp, err := c.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	// Check for non-200 status code
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(body))
	}

	// Parse response JSON
	var openAIResp openAIResponse
	if err := json.Unmarshal(body, &openAIResp); err != nil {
		return "", fmt.Errorf("failed to unmarshal response: %w", err)
	}

	// Ensure we got at least one choice back
	if len(openAIResp.Choices) == 0 {
		return "", fmt.Errorf("no content in response")
	}

	// Return the content of the first choice
	return openAIResp.Choices[0].Message.Content, nil
}
