// Package api provides integration with DeepSeek's LLM models.
//
// This file implements the Client interface for DeepSeek's API, supporting
// both text and JSON queries using the OpenAI-compatible  interface.
// It handles authentication, request formatting, and response parsing.
package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

var deepseekEndpoint string

// DeepSeekClient implements the Client interface for DeepSeek's API
type DeepSeekClient struct {
	apiKey string       // DeepSeek API authentication key
	client *http.Client // HTTP client for making API requests
}

// deepseekRequest represents the structure of a request to DeepSeek's chat completion API
type deepseekRequest struct {
	Model          string             `json:"model"`                           
	Messages       []deepseekMessage  `json:"messages"`                        
	MaxTokens      int                `json:"max_completion_tokens,omitempty"` 
	ResponseFormat string             `json:"response_format,omitempty"`       
}

type deepseekMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type deepseekResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

func (c *DeepSeekClient) QueryText(prompt string, model string, options Options) (string, error) {
	if prompt == "" {
		return "", fmt.Errorf("prompt cannot be empty for text query")
	}

	// Initialize endpoint from environment if not already set
	if deepseekEndpoint == "" {
		base := os.Getenv("DEEPSEEK_API_BASE")
		if base == "" {
			return "", fmt.Errorf("DEEPSEEK_API_BASE environment variable not set")
		}
		deepseekEndpoint = base + "/v1/chat/completions"
	}

	// Initialize HTTP client and API key if not already done
	if c.client == nil {
		c.client = &http.Client{}
		c.apiKey = os.Getenv("DEEPSEEK_API_KEY")
		if c.apiKey == "" {
			return "", fmt.Errorf("DEEPSEEK_API_KEY environment variable not set")
		}
	}

	// Construct the request body
	reqBody := deepseekRequest{
		Model: model,
		Messages: []deepseekMessage{
			{Role: "user", Content: prompt},
		},
		MaxTokens: 1024, // Limit response length
	}

	return c.makeRequest(reqBody)
}

func (c *DeepSeekClient) QueryJSON(prompt string, model string, options Options) (string, error) {
	if prompt == "" {
		return "", fmt.Errorf("prompt cannot be empty for json query")
	}

	// Initialize HTTP client and API key if not already done
	if c.client == nil {
		c.client = &http.Client{}
		c.apiKey = os.Getenv("DEEPSEEK_API_KEY")
		if c.apiKey == "" {
			return "", fmt.Errorf("DEEPSEEK_API_KEY environment variable not set")
		}
	}

	// Construct the request body
	reqBody := deepseekRequest{
		Model: model,
		Messages: []deepseekMessage{
			{Role: "user", Content: prompt},
		},
		MaxTokens: 1024, // Limit response length
	}

	return c.makeRequest(reqBody)
}

func (c *DeepSeekClient) makeRequest(reqBody deepseekRequest) (string, error) {
	// Convert request body to JSON
	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request: %w", err)
	}

	// Create new HTTP request with JSON body
	req, err := http.NewRequest("POST", deepseekEndpoint, bytes.NewBuffer(jsonBody))
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
	var deepseekResp deepseekResponse
	if err := json.Unmarshal(body, &deepseekResp); err != nil {
		return "", fmt.Errorf("failed to unmarshal response: %w", err)
	}

	// Ensure we got at least one choice back
	if len(deepseekResp.Choices) == 0 {
		return "", fmt.Errorf("no content in response")
	}

	// Return the content of the first choice
	return deepseekResp.Choices[0].Message.Content, nil
}
