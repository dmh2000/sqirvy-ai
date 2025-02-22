// Package api provides integration with Meta's Llama models via langchaingo.
//
// This file implements the Client interface for Meta's Llama models using
// langchaingo's OpenAI-compatible interface. It handles model initialization,
// prompt formatting, and response parsing.
package sqirvy

import (
	"context"
	"fmt"
	"os"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
)

const (
	// LlamaTempScale is the scaling factor for Llama's 0-2 temperature range
	LlamaTempScale = 2.0
)

// LlamaClient implements the Client interface for Meta's Llama models.
// It provides methods for querying Llama language models through
// an OpenAI-compatible interface.
type LlamaClient struct {
	llm llms.Model // OpenAI-compatible LLM client
}

// Ensure LlamaClient implements the Client interface
var _ Client = (*LlamaClient)(nil)

// NewLlamaClient creates a new instance of LlamaClient.
// It returns an error if the required environment variables are not set.
func NewLlamaClient() (*LlamaClient, error) {
	apiKey := os.Getenv("LLAMA_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("LLAMA_API_KEY environment variable not set")
	}

	baseURL := os.Getenv("LLAMA_BASE_URL")
	if baseURL == "" {
		return nil, fmt.Errorf("LLAMA_BASE_URL environment variable not set")
	}

	llm, err := openai.New(
		openai.WithBaseURL(baseURL),
		openai.WithToken(apiKey),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create Llama client: %w", err)
	}

	return &LlamaClient{
		llm: llm,
	}, nil
}

func (c *LlamaClient) QueryText(ctx context.Context, prompts []string, model string, options Options) (string, error) {
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
	// Scale temperature for Llama's 0-2 range
	options.Temperature = (options.Temperature * LlamaTempScale) / MaxTemperature

	// Call the LLM with the prompt
	completion, err := llms.GenerateFromSinglePrompt(
		ctx,
		c.llm,
		prompts[0],
		llms.WithTemperature(float64(options.Temperature)),
		llms.WithModel(model),
	)
	if err != nil {
		return "", fmt.Errorf("failed to generate completion: %w", err)
	}

	return completion, nil
}

// Close implements the Close method for the Client interface.
func (c *LlamaClient) Close() error {
	return nil
}
