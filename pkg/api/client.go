package ai

import (
	"fmt"
)

// Provider represents supported AI providers
type Provider string

const (
	Anthropic Provider = "anthropic"
	Gemini    Provider = "gemini"
	OpenAI    Provider = "openai"
)

type AnthropicOptions struct {
	// Anthropic supports specifying a json response format

}
type GeminiOptions struct {
	// Gemini supports specifying a json response format

}

type OpenAIOptions struct {
	// open ai supports specifying a json response format
	OpenAIResponseFormat string
}

type Options struct {
	AnthropicOptions
	GeminiOptions
	OpenAIOptions
}

// Client provides a unified interface for AI operations
type Client interface {
	QueryText(prompt string, model string, options Options) (string, error)
	QueryJSON(prompt string, model string, options Options) (string, error)
}

// NewClient creates a new AI client for the specified provider
func NewClient(provider Provider) (Client, error) {
	switch provider {
	case Anthropic:
		return &AnthropicClient{}, nil
	case Gemini:
		return &GeminiClient{}, nil
	case OpenAI:
		return &OpenAIClient{}, nil
	default:
		return nil, fmt.Errorf("unsupported provider: %s", provider)
	}
}
