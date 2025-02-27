// Package api provides integration with OpenAI's GPT models.
//
// This file implements the Client interface for OpenAI's API, supporting
// both text and JSON queries to GPT models. It handles authentication,
// request formatting, and response parsing specific to OpenAI's requirements.
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
	// OpenAITempScale is the scaling factor for OpenAI's 0-2 temperature range
	OpenAITempScale = 2.0
)

// OpenAIClient implements the Client interface for OpenAI's API.
// It provides methods for querying OpenAI's language models through
// their HTTP API.
type OpenAIClient struct {
	apiKey  string       // OpenAI API authentication key
	baseURL string       // OpenAI API base URL
	client  *http.Client // HTTP client for making API requests
}

// Ensure OpenAIClient implements the Client interface
var _ Client = (*OpenAIClient)(nil)

// NewOpenAIClient creates a new instance of OpenAIClient.
// It returns an error if the required OPENAI_API_KEY environment variable is not set.
func NewOpenAIClient() (*OpenAIClient, error) {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("OPENAI_API_KEY environment variable not set")
	}

	baseURL := os.Getenv("OPENAI_BASE_URL")
	if baseURL == "" {
		return nil, fmt.Errorf("OPENAI_BASE_URL environment variable not set")
	}

	return &OpenAIClient{
		apiKey:  apiKey,
		baseURL: baseURL,
		client:  &http.Client{},
	}, nil
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

// OpenAIClient.QueryText implements the QueryText method for the Client interface.
// It sends a text query to OpenAI's API and returns the generated text response.
func (c *OpenAIClient) QueryText(ctx context.Context, system string, prompts []string, model string, options Options) (string, error) {
	if ctx.Err() != nil {
		return "", fmt.Errorf("request context error %w", ctx.Err())
	}

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
	// Scale temperature for OpenAI's 0-2 range
	options.Temperature = (options.Temperature * OpenAITempScale) / MaxTemperature

	// Set default max tokens if not specified
	maxTokens := options.MaxTokens
	if maxTokens == 0 {
		maxTokens = MaxTokensDefault
	}

	// create the slice of prompts
	messages := make([]openAIMessage, 0, len(prompts))
	messages = append(messages, openAIMessage{Role: "system", Content: system})
	for _, prompt := range prompts {
		messages = append(messages, openAIMessage{Role: "user", Content: prompt})
	}

	// Construct the request body with the prompt as a user message
	reqBody := openAIRequest{
		Model:       model,
		Messages:    messages,
		MaxTokens:   int(maxTokens),      // Limit response length
		Temperature: options.Temperature, // Set temperature
	}

	// Send request and return response
	return c.makeRequest(ctx, reqBody)
}

func (c *OpenAIClient) makeRequest(ctx context.Context, reqBody openAIRequest) (string, error) {
	// update the endpoing if OPENAI_BASE_URL is set
	endpoint := c.baseURL + "/v1/chat/completions"

	// Convert request body to JSON
	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request: %w", err)
	}

	ctx, cancel := context.WithTimeout(ctx, RequestTimeout)
	defer cancel()

	// Create new HTTP request with JSON body
	req, err := http.NewRequestWithContext(ctx, "POST", endpoint, bytes.NewBuffer(jsonBody))
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

// Close implements the Close method for the Client interface.
func (c *OpenAIClient) Close() error {
	// http.client does not require explicit close
	return nil
}
