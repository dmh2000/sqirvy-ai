// Package api provides model management functionality for AI language models.
//
// This file contains model-to-provider mappings and utility functions for
// working with different AI models across supported providers.
package sqirvy

import "fmt"

var modelAlias = map[string]string{
	"claude-3-7-sonnet": "claude-3-7-sonnet-latest",
	"claude-3-5-sonnet": "claude-3-5-sonnet-latest",
	"claude-3-5-haiku":  "claude-3-5-haiku-latest",
	"claude-3-opus":     "claude-3-opus-latest",
}

func GetModelAlias(model string) string {
	if alias, ok := modelAlias[model]; ok {
		return alias
	}
	return model
}

// Supported AI providers
const (
	Anthropic string = "anthropic" // Anthropic's Claude models
	DeepSeek  string = "deepseek"  // DeepSeek's models
	Gemini    string = "gemini"    // Google's Gemini models
	OpenAI    string = "openai"    // OpenAI's GPT models
	Llama     string = "llama"     // Meta's Llama models
)

// ModelToProvider maps model names to their respective providers.
// This mapping is used to determine which client implementation should handle
// requests for a given model.
var modelToProvider = map[string]string{
	// anthropic models
	"claude-3-7-sonnet-20250219": Anthropic,
	"claude-3-5-sonnet-20241022": Anthropic,
	"claude-3-7-sonnet-latest":   Anthropic,
	"claude-3-5-sonnet-latest":   Anthropic,
	"claude-3-5-haiku-latest":    Anthropic,
	"claude-3-haiku-20240307":    Anthropic,
	"claude-3-opus-latest":       Anthropic,
	"claude-3-opus-20240229":     Anthropic,
	// deepseek models
	"deepseek-r1": DeepSeek,
	"deepseek-v3": DeepSeek,
	// google gemini models
	// gemini-2.0-pro-exp-02-05
	// gemini-2.0-flash-thinking-exp-01-21
	"gemini-2.0-flash":              Gemini,
	"gemini-1.5-flash":              Gemini,
	"gemini-1.5-pro":                Gemini,
	"gemini-2.0-flash-thinking-exp": Gemini,
	// openai models
	"gpt-4o":      OpenAI,
	"gpt-4o-mini": OpenAI,
	"gpt-4-turbo": OpenAI,
	"o1-mini":     OpenAI,
	// "o3-mini":     "openai",
	// llama models
	"llama3.3-70b": Llama,
}

// ModelToMaxTokens maps model names to their maximum token limits.
// If a model is not in this map, MaxTokensDefault will be used.
var modelToMaxTokens = map[string]int64{
	// anthropic models
	"claude-3-7-sonnet-latest": MaxTokensDefault,
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
	"gpt-4o":      4096,
	"gpt-4o-mini": 4096,
	"gpt-4-turbo": 4096,
	"o1-mini":     MaxTokensDefault,
	// llama models
	"llama3.3-70b": MaxTokensDefault,
}

func GetModelList() []string {
	var models []string
	for model := range modelToProvider {
		models = append(models, model)
	}
	return models
}

type ModelProvider struct {
	Model    string
	Provider string
}

func GetModelProviderList() []ModelProvider {
	var mp []ModelProvider
	for model, provider := range modelToProvider {
		mp = append(mp, ModelProvider{Model: model, Provider: provider})
	}
	return mp
}

// GetProviderName returns the provider name for a given model identifier.
// Returns an error if the model is not recognized.
func GetProviderName(model string) (string, error) {
	if provider, ok := modelToProvider[model]; ok {
		return provider, nil
	}
	return "", fmt.Errorf("unrecognized model: %s", model)
}

// GetMaxTokens returns the maximum token limit for a given model identifier.
// Returns MaxTokensDefault if the model is not in ModelToMaxTokens.
func GetMaxTokens(model string) int64 {
	if maxTokens, ok := modelToMaxTokens[model]; ok {
		return maxTokens
	}
	return MaxTokensDefault
}
