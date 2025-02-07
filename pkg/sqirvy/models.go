// Package api provides model management functionality for AI language models.
//
// This file contains model-to-provider mappings and utility functions for
// working with different AI models across supported providers.
package api

import "fmt"

const MAX_TOKENS = 4096

// ModelToProvider maps model names to their respective providers.
// This mapping is used to determine which client implementation should handle
// requests for a given model.
var ModelToProvider = map[string]string{
	// anthropic models
	"claude-3-5-sonnet-latest": "anthropic",
	"claude-3-5-haiku-latest":  "anthropic",
	"claude-3-opus-latest":     "anthropic",
	// deepseek models
	//"deepseek-chat": "deepseek",
	//"deepseek-reasoner": "deepseek",
	// google gemini models
	// gemini-2.0-pro-exp-02-05
	// gemini-2.0-flash-thinking-exp-01-21
	"gemini-2.0-flash-exp": "gemini",
	"gemini-1.5-flash":     "gemini",
	"gemini-1.5-pro":       "gemini",
	// openai models
	"gpt-4o":      "openai",
	"gpt-4o-mini": "openai",
	"gpt-4-turbo": "openai",
	"o1-mini":     "openai",
	// meta-llama models
	"meta-llama/meta-llama-3.1-8b-instruct-turbo": "meta-llama",
	"meta-llama/Llama-3.3-70B-Instruct-Turbo":     "meta-llama",
}

// GetProviderName returns the provider name for a given model identifier.
// Returns an error if the model is not recognized.
func GetProviderName(model string) (string, error) {
	if provider, ok := ModelToProvider[model]; ok {
		return provider, nil
	}
	return "", fmt.Errorf("unrecognized model: %s", model)
}
