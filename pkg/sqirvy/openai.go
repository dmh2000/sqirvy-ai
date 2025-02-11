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

// OpenAIClient implements the Client interface for OpenAI's API
type OpenAIClient struct {
	apiKey  string       // OpenAI API authentication key
	baseURL string       // OpenAI API base URL
	client  *http.Client // HTTP client for making API requests
}

// openAIRequest represents the structure of a request to OpenAI's chat completion API
type openAIRequest struct {
	Model          string          `json:"model"`                           // Model identifier
	Messages       []openAIMessage `json:"messages"`                        // Conversation messages
	MaxTokens      int             `json:"max_completion_tokens,omitempty"` // Max response length
	ResponseFormat string          `json:"response_format,omitempty"`       // Desired response format
	Temperature    float32         `json:"temperature,omitempty"`           // Controls the randomness of the output
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

	// Initialize HTTP client and API key if not already done
	if c.client == nil {
		c.client = &http.Client{}
		// Get API key from environment variable
		c.apiKey = os.Getenv("OPENAI_API_KEY")
		if c.apiKey == "" {
			return "", fmt.Errorf("OPENAI_API_KEY environment variable not set")
		}

		// Get base URL from environment variable, or use default
		c.baseURL = os.Getenv("OPENAI_BASE_URL")
		if c.baseURL == "" {
			c.baseURL = "https://api.openai.com" // Default OpenAI base URL
		}
	}

	// validate temperature
	if options.Temperature < 0.0 {
		options.Temperature = 0.0
	}
	if options.Temperature > 100.0 {
		return "", fmt.Errorf("temperature must be between 1 and 100")
	}
	// scale Temperature for openai 0..2.0
	options.Temperature = (options.Temperature * 2) / 100.0

	// Construct the request body with the prompt as a user message
	reqBody := openAIRequest{
		Model: model,
		Messages: []openAIMessage{
			{Role: "user", Content: prompt},
		},
		MaxTokens:   MAX_TOKENS,          // Limit response length
		Temperature: options.Temperature, // Set temperature
	}

	// Send request and return response
	return c.makeRequest(reqBody)
}

func (c *OpenAIClient) makeRequest(reqBody openAIRequest) (string, error) {
	// update the endpoing if OPENAI_BASE_URL is set
	endpoint := c.baseURL + "/v1/chat/completions"

	// Convert request body to JSON
	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request: %w", err)
	}

	// Create new HTTP request with JSON body
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonBody))
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
