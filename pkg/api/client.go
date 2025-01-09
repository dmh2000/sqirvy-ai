package ai

import (
	"fmt"
)

// Provider represents supported AI providers
type Provider string

const (
	Anthropic Provider = "anthropic"
	Gemini   Provider = "gemini"
	OpenAI   Provider = "openai"
)

// Client provides a unified interface for AI operations
type Client interface {
	QueryText(prompt string) (string, error)
	QueryJSON(prompt string) (string, error)
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
