package ai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

const (
	openAIEndpoint = "https://api.openai.com/v1/chat/completions"
)

type OpenAIClient struct {
	apiKey string
	client *http.Client
}

type openAIRequest struct {
	Model          string          `json:"model"`
	Messages       []openAIMessage `json:"messages"`
	MaxTokens      int             `json:"max_tokens"`
	ResponseFormat string          `json:"response_format,omitempty"`
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
	if c.client == nil {
		c.client = &http.Client{}
		c.apiKey = os.Getenv("OPENAI_API_KEY")
		if c.apiKey == "" {
			return "", fmt.Errorf("OPENAI_API_KEY environment variable not set")
		}
	}

	reqBody := openAIRequest{
		Model: model,
		Messages: []openAIMessage{
			{Role: "user", Content: prompt},
		},
		MaxTokens: 1024,
	}

	return c.makeRequest(reqBody)
}

// QueryJSON sends a JSON query to OpenAI and returns the response
// using json has some options, see:
// https://platform.openai.com/docs/guides/structured-outputs#examples
func (c *OpenAIClient) QueryJSON(prompt string, model string, options Options) (string, error) {
	if prompt == "" {
		return "", fmt.Errorf("prompt cannot be empty for json query")
	}

	if c.client == nil {
		c.client = &http.Client{}
		c.apiKey = os.Getenv("OPENAI_API_KEY")
		if c.apiKey == "" {
			return "", fmt.Errorf("OPENAI_API_KEY environment variable not set")
		}
	}

	reqBody := openAIRequest{
		Model: model,
		Messages: []openAIMessage{
			{Role: "user", Content: prompt},
		},
		MaxTokens:      1024,
		ResponseFormat: options.OpenAIResponseFormat,
	}

	return c.makeRequest(reqBody)
}

func (c *OpenAIClient) makeRequest(reqBody openAIRequest) (string, error) {
	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request: %v", err)
	}

	req, err := http.NewRequest("POST", openAIEndpoint, bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.apiKey)

	resp, err := c.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to make request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(body))
	}

	var openAIResp openAIResponse
	if err := json.Unmarshal(body, &openAIResp); err != nil {
		return "", fmt.Errorf("failed to unmarshal response: %v", err)
	}

	if len(openAIResp.Choices) == 0 {
		return "", fmt.Errorf("no content in response")
	}

	return openAIResp.Choices[0].Message.Content, nil
}
