package main

import "fmt"

var modelToProvider = map[string]string{
	"claude-3.5-sonnet":         "anthropic",
	"gemini-2.0-flash-exp":      "gemini",
	"openai-gpt-4-turbo-preview": "openai",
}

func getProviderName(model string) (string, error) {
	if provider, ok := modelToProvider[model]; ok {
		return provider, nil
	}
	return "", fmt.Errorf("unrecognized model: %s", model)
}
