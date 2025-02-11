// Package api provides integration with Meta's Llama models via langchaingo.
//
// This file implements the Client interface for Meta's Llama models using
// langchaingo's OpenAI-compatible interface. It handles model initialization,
// prompt formatting, and response parsing.
package api

import (
	"context"
	"fmt"
	"os"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
)

// LlamaClient implements the Client interface for Meta's Llama models
type LlamaClient struct {
	llm llms.Model // OpenAI-compatible LLM client
}

func (c *LlamaClient) QueryText(prompt string, model string, options Options) (string, error) {
	if prompt == "" {
		return "", fmt.Errorf("prompt cannot be empty for text query")
	}

	// validate temperature
	if options.Temperature < 0.0 {
		options.Temperature = 0.0
	}
	if options.Temperature > 100.0 {
		return "", fmt.Errorf("temperature must be between 1 and 100")
	}
	// scale Temperature for openai 0..2.0
	options.Temperature = (options.Temperature * 2.0) / 100.0

	// Initialize LLM if not already done
	if c.llm == nil {
		apiKey := os.Getenv("LLAMA_API_KEY")
		if apiKey == "" {
			return "", fmt.Errorf("LLAMA_API_KEY environment variable not set")
		}

		baseURL := os.Getenv("LLAMA_BASE_URL")
		if baseURL == "" {
			// baseurl must be provided for Llama models
			return "", fmt.Errorf("LLAMA_BASE_URL environment variable not set")
		}

		llm, err := openai.New(
			openai.WithBaseURL(os.Getenv("LLAMA_BASE_URL")),
			openai.WithToken(apiKey),
			openai.WithModel(model),
		)
		if err != nil {
			return "", fmt.Errorf("failed to create Together client: %w", err)
		}
		c.llm = llm
	}

	// Call the LLM with the prompt
	completion, err := llms.GenerateFromSinglePrompt(context.Background(), c.llm, prompt, llms.WithTemperature(float64(options.Temperature)))
	if err != nil {
		return "", fmt.Errorf("failed to generate completion: %w", err)
	}

	return completion, nil
}
