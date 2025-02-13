package sqirvy

import (
	"context"
	"os"
	"testing"
)

func TestAllModels(t *testing.T) {
	// Test cases for both QueryText and QueryJSON
	tests := []struct {
		name    string
		prompt  string
		wantErr bool
	}{
		{
			name:    "Basic prompt",
			prompt:  "Say 'Hello, World!'",
			wantErr: false,
		},
		{
			name:    "JSON request",
			prompt:  "Return a JSON object with a greeting field containing 'Hello, World!'",
			wantErr: false,
		},
		{
			name:    "Empty prompt",
			prompt:  "",
			wantErr: true,
		},
	}

	// Test each model from ModelToProvider
	for model, provider := range ModelToProvider {
		// Create client for this provider
		client, err := NewClient(Provider(provider))
		if err != nil {
			t.Errorf("Failed to create client for provider %s: %v", provider, err)
			continue
		}

		// Check if required API key is set
		var apiKey string
		switch provider {
		case "anthropic":
			apiKey = os.Getenv("ANTHROPIC_API_KEY")
		case "deepseek":
			apiKey = os.Getenv("LLAMA_API_KEY")
		case "gemini":
			apiKey = os.Getenv("GEMINI_API_KEY")
		case "openai":
			apiKey = os.Getenv("OPENAI_API_KEY")
		case "llama":
			apiKey = os.Getenv("LLAMA_API_KEY")
		}

		if apiKey == "" {
			t.Logf("Skipping tests for %s model %s: API key not set", provider, model)
			continue
		}

		// Test QueryText
		t.Run(model+"_QueryText", func(t *testing.T) {
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					ctx := context.Background()
					got, err := client.QueryText(ctx, tt.prompt, model, Options{MaxTokens: GetMaxTokens(model), Temperature: 50})
					if tt.wantErr {
						if err == nil {
							t.Errorf("QueryText() error = nil, wantErr %v", tt.wantErr)
						}
						return
					}
					if err != nil {
						t.Errorf("QueryText() error = %v", err)
						return
					}
					if len(got) == 0 {
						t.Error("QueryText() returned empty response")
					}
				})
			}
		})
	}
}
