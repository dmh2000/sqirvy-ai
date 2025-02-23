// Package api provides integration with deepseek's GPT models.
//
// This file implements the Client interface for deepseek's API, supporting
// both text and JSON queries to GPT models. It handles authentication,
// request formatting, and response parsing specific to deepseek's requirements.
package sqirvy

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

const (
	// DeepSeekTempScale is the scaling factor for DeepSeek's 0-2 temperature range
	DeepSeekTempScale = 2.0
)

// DeepSeekClient implements the Client interface for DeepSeek's API.
// It provides methods for querying DeepSeek's language models through
// their HTTP API.
type DeepSeekClient struct {
	apiKey  string       // DeepSeek API authentication key
	baseURL string       // DeepSeek API base URL
	client  *http.Client // HTTP client for making API requests
}

// Ensure DeepSeekClient implements the Client interface
var _ Client = (*DeepSeekClient)(nil)

// NewDeepSeekClient creates a new instance of DeepSeekClient.
// It returns an error if the required environment variables are not set.
func NewDeepSeekClient() (*DeepSeekClient, error) {
	apiKey := os.Getenv("DEEPSEEK_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("DEEPSEEK_API_KEY environment variable not set")
	}

	baseURL := os.Getenv("DEEPSEEK_BASE_URL")
	if baseURL == "" {
		return nil, fmt.Errorf("DEEPSEEK_BASE_URL environment variable not set")
	}

	return &DeepSeekClient{
		apiKey:  apiKey,
		baseURL: baseURL,
		client:  &http.Client{},
	}, nil
}

// deepseekRequest represents the structure of a request to deepseek's chat completion API
type deepseekRequest struct {
	Model          string            `json:"model"`                           // Model identifier
	Messages       []deepseekMessage `json:"messages"`                        // Conversation messages
	MaxTokens      int               `json:"max_completion_tokens,omitempty"` // Max response length
	ResponseFormat string            `json:"response_format,omitempty"`       // Desired response format
	Temperature    float32           `json:"temperature,omitempty"`           // Controls the randomness of the output
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

func (c *DeepSeekClient) QueryText(ctx context.Context, system string, prompts []string, model string, options Options) (string, error) {
	if len(prompts) == 0 {
		return "", fmt.Errorf("prompts cannot be empty for text query")
	}

	// Set default and validate temperature
	if options.Temperature < MinTemperature {
		options.Temperature = MinTemperature
	}
	if options.Temperature > MaxTemperature {
		return "", fmt.Errorf("temperature must be between %.1f and %.1f", MinTemperature, MaxTemperature)
	}
	// Scale temperature for DeepSeek's 0-2 range
	options.Temperature = (options.Temperature * DeepSeekTempScale) / MaxTemperature

	// Set default max tokens if not specified
	maxTokens := options.MaxTokens
	if maxTokens == 0 {
		maxTokens = MaxTokensDefault
	}

	messages := make([]deepseekMessage, 0, len(prompts))
	// First prompt is system prompt
	messages = append(messages, deepseekMessage{Role: "system", Content: system})
	for i := 1; i < len(prompts); i++ {
		messages = append(messages, deepseekMessage{Role: "user", Content: prompts[i]})
	}

	// Construct the request body with the prompt as a user message
	reqBody := deepseekRequest{
		Model:       model,
		Messages:    messages,
		MaxTokens:   int(maxTokens),      // Limit response length
		Temperature: options.Temperature, // Set temperature
	}

	// Send request and return response
	return c.makeRequest(ctx, reqBody)
}

func (c *DeepSeekClient) makeRequest(ctx context.Context, reqBody deepseekRequest) (string, error) {
	// Convert request body to JSON
	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request: %w", err)
	}

	// Create new HTTP request with JSON body
	endpoint := c.baseURL + "/chat/completions"

	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	// Set required headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.apiKey)

	// Send the request
	resp, err := c.client.Do(req.WithContext(ctx))
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

// Close implements the Close method for the Client interface.
func (c *DeepSeekClient) Close() error {
	c.client.CloseIdleConnections()
	return nil
}
