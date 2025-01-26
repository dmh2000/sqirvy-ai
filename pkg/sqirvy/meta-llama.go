// Package api provides integration with Meta's Llama models via Together.ai.
//
// This file implements the Client interface for Meta's Llama models using Together.ai's
// OpenAI-compatible REST API. It handles authentication, request formatting,
// and response parsing specific to Together.ai's requirements.
package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

var togetherEndpoint = "https://api.together.xyz/v1/chat/completions"

// MetaLlamaClient implements the Client interface for Meta's Llama models
type MetaLlamaClient struct {
	apiKey string       // Together.ai API authentication key
	client *http.Client // HTTP client for making API requests
}

// metaLlamaRequest represents the structure of a request to Together.ai's chat completion API
type metaLlamaRequest struct {
	Model          string             `json:"model"`
	Messages       []metaLlamaMessage `json:"messages"`
	MaxTokens      int                `json:"max_tokens,omitempty"`
	ResponseFormat string             `json:"response_format,omitempty"`
}

type metaLlamaMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type metaLlamaResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

func (c *MetaLlamaClient) QueryText(prompt string, model string, options Options) (string, error) {
	if prompt == "" {
		return "", fmt.Errorf("prompt cannot be empty for text query")
	}

	// Initialize HTTP client and API key if not already done
	if c.client == nil {
		c.client = &http.Client{}
		c.apiKey = os.Getenv("TOGETHER_API_KEY")
		if c.apiKey == "" {
			return "", fmt.Errorf("TOGETHER_API_KEY environment variable not set")
		}
	}

	// Construct the request body
	reqBody := metaLlamaRequest{
		Model: model,
		Messages: []metaLlamaMessage{
			{Role: "user", Content: prompt},
		},
		MaxTokens: 1024, // Limit response length
	}

	return c.makeRequest(reqBody)
}

func (c *MetaLlamaClient) QueryJSON(prompt string, model string, options Options) (string, error) {
	if prompt == "" {
		return "", fmt.Errorf("prompt cannot be empty for json query")
	}

	// Initialize HTTP client and API key if not already done
	if c.client == nil {
		c.client = &http.Client{}
		c.apiKey = os.Getenv("TOGETHER_API_KEY")
		if c.apiKey == "" {
			return "", fmt.Errorf("TOGETHER_API_KEY environment variable not set")
		}
	}

	// Construct the request body
	reqBody := metaLlamaRequest{
		Model: model,
		Messages: []metaLlamaMessage{
			{Role: "user", Content: prompt},
		},
		MaxTokens:      1024, // Limit response length
		ResponseFormat: "json_object",
	}

	return c.makeRequest(reqBody)
}

func (c *MetaLlamaClient) makeRequest(reqBody metaLlamaRequest) (string, error) {
	// Convert request body to JSON
	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request: %w", err)
	}

	// Create new HTTP request with JSON body
	req, err := http.NewRequest("POST", togetherEndpoint, bytes.NewBuffer(jsonBody))
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
	var llamaResp metaLlamaResponse
	if err := json.Unmarshal(body, &llamaResp); err != nil {
		return "", fmt.Errorf("failed to unmarshal response: %w", err)
	}

	// Ensure we got at least one choice back
	if len(llamaResp.Choices) == 0 {
		return "", fmt.Errorf("no content in response")
	}

	// Return the content of the first choice
	return llamaResp.Choices[0].Message.Content, nil
}
