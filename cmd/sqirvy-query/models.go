package main

import "fmt"

var modelToProvider = map[string]string{
	"claude-3-5-sonnet-latest": "anthropic",
	"gemini-2.0-flash-exp":     "gemini",
	"gpt-4o-2024-11-20":        "openai",
}

func getProviderName(model string) (string, error) {
	if provider, ok := modelToProvider[model]; ok {
		return provider, nil
	}
	return "", fmt.Errorf("unrecognized model: %s", model)
}
