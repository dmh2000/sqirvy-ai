package ai

import "fmt"

var ModelToProvider = map[string]string{
	// anthropic models
	"claude-3-5-sonnet-latest": "anthropic",
	"claude-3-5-haiku-latest":  "anthropic",
	"claude-3-opus-latest":     "anthropic",
	// google gemini models
	"gemini-2.0-flash-exp": "gemini",
	"gemini-1.5-flash":     "gemini",
	"gemini-1.5-pro":       "gemini",
	// openai models
	"gpt-4o":      "openai",
	"gpt-4o-mini": "openai",
	"gpt-4-turbo": "openai",
	"o1-mini":     "openai",
}

func GetProviderName(model string) (string, error) {
	if provider, ok := ModelToProvider[model]; ok {
		return provider, nil
	}
	return "", fmt.Errorf("unrecognized model: %s", model)
}
