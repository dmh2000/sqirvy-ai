package api

import (
	"os"
	"strings"
	"testing"
)

func TestGeminiClient_QueryText(t *testing.T) {
	if os.Getenv("GEMINI_API_KEY") == "" {
		t.Skip("GEMINI_API_KEY not set")
	}

	tests := []struct {
		name   string
		prompt string
	}{
		{
			name:   "Basic prompt",
			prompt: "Say 'Hello, World!'",
		},
		{
			name:   "Empty prompt",
			prompt: "hello world",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &GeminiClient{}
			got, err := client.QueryText(tt.prompt, "gemini-2.0-flash-exp", Options{})
			if err != nil {
				t.Errorf("GeminiClient.QueryText() error = %v", err)
				return
			}
			if len(got) == 0 {
				t.Error("GeminiClient.QueryText() returned empty response")
			}
		})
	}
}

