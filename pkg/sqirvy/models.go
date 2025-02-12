// Package api provides model management functionality for AI language models.
//
// This file contains model-to-provider mappings and utility functions for
// working with different AI models across supported providers.
package sqirvy

import "fmt"

// ModelToProvider maps model names to their respective providers.
// This mapping is used to determine which client implementation should handle
// requests for a given model.
var ModelToProvider = map[string]string{
	// anthropic models
	"claude-3-5-sonnet-latest": "anthropic",
	"claude-3-5-haiku-latest":  "anthropic",
	"claude-3-opus-latest":     "anthropic",
	// deepseek models
	"deepseek-r1": "deepseek",
	"deepseek-v3": "deepseek",
	// google gemini models
	// gemini-2.0-pro-exp-02-05
	// gemini-2.0-flash-thinking-exp-01-21
	"gemini-2.0-flash": "gemini",
	"gemini-1.5-flash": "gemini",
	"gemini-1.5-pro":   "gemini",
	// openai models
	"gpt-4o":      "openai",
	"gpt-4o-mini": "openai",
	"gpt-4-turbo": "openai",
	"o1-mini":     "openai",
	// llama models
	"llama3.3-70b": "llama",
}

// ModelToMaxTokens maps model names to their maximum token limits.
// If a model is not in this map, MaxTokensDefault will be used.
var ModelToMaxTokens = map[string]int64{
	// anthropic models
	"claude-3-5-sonnet-latest": MaxTokensDefault,
	"claude-3-5-haiku-latest":  MaxTokensDefault,
	"claude-3-opus-latest":     4096,
	// deepseek models
	"deepseek-r1": MaxTokensDefault,
	"deepseek-v3": MaxTokensDefault,
	// google gemini models
	"gemini-2.0-flash": MaxTokensDefault,
	"gemini-1.5-flash": MaxTokensDefault,
	"gemini-1.5-pro":   MaxTokensDefault,
	// openai models
	"gpt-4o":      MaxTokensDefault,
	"gpt-4o-mini": MaxTokensDefault,
	"gpt-4-turbo": 4096,
	"o1-mini":     MaxTokensDefault,
	// llama models
	"llama3.3-70b": MaxTokensDefault,
}

// GetProviderName returns the provider name for a given model identifier.
// Returns an error if the model is not recognized.
func GetProviderName(model string) (string, error) {
	if provider, ok := ModelToProvider[model]; ok {
		return provider, nil
	}
	return "", fmt.Errorf("unrecognized model: %s", model)
}

// GetMaxTokens returns the maximum token limit for a given model identifier.
// Returns MaxTokensDefault if the model is not in ModelToMaxTokens.
func GetMaxTokens(model string) int64 {
	if maxTokens, ok := ModelToMaxTokens[model]; ok {
		return maxTokens
	}
	return MaxTokensDefault
}
