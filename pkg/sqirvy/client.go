// Package api provides a unified interface for interacting with various AI language models.
//
// The package supports multiple AI providers including:
// - Anthropic (Claude models)
// - Google (Gemini models)
// - OpenAI (GPT models)
//
// It provides a consistent interface for making text and JSON queries while handling
// provider-specific implementation details internally.
package sqirvy

import (
	"context"
	"fmt"
)

const (
	// MaxTokensDefault is the default maximum number of tokens in responses
	MaxTokensDefault = 4096

	// Temperature limits for model queries (0-100 scale)
	MinTemperature = 0.0
	MaxTemperature = 100.0
)

// Provider represents supported AI providers.
// Currently supports Anthropic, DeepSeek, Gemini, and OpenAI.
// Provider identifies which AI service provider to use
type Provider string

// Supported AI providers
const (
	Anthropic Provider = "anthropic" // Anthropic's Claude models
	DeepSeek  Provider = "deepseek"  // DeepSeek's models
	Gemini    Provider = "gemini"    // Google's Gemini models
	OpenAI    Provider = "openai"    // OpenAI's GPT models
	MetaLlama Provider = "llama"     // Meta's Llama models
)

// Options combines all provider-specific options into a single structure.
// This allows for provider-specific configuration while maintaining a unified interface.
type Options struct {
	Temperature float32 // Controls the randomness of the output
	MaxTokens   int64   // Maximum number of tokens in the response
}

// Client provides a unified interface for AI operations.
// It abstracts away provider-specific implementations behind a common interface
// for making text and JSON queries to AI models.
type Client interface {
	QueryText(ctx context.Context, prompt string, model string, options Options) (string, error)
	Close() error
}

// NewClient creates a new AI client for the specified provider
func NewClient(provider Provider) (Client, error) {
	switch provider {
	case Anthropic:
		client, err := NewAnthropicClient()
		if err != nil {
			return nil, err
		}
		return client, nil
	case DeepSeek:
		client, err := NewDeepSeekClient()
		if err != nil {
			return nil, err
		}
		return client, nil
	case Gemini:
		client, err := NewGeminiClient()
		if err != nil {
			return nil, err
		}
		return client, nil
	case OpenAI:
		client, err := NewOpenAIClient()
		if err != nil {
			return nil, err
		}
		return client, nil
	case MetaLlama:
		client, err := NewLlamaClient()
		if err != nil {
			return nil, err
		}
		return client, nil
	default:
		return nil, fmt.Errorf("unsupported provider: %s", provider)
	}
}
