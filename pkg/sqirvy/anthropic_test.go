package api

import (
	"os"
	"strings"
	"testing"
)

func TestAnthropicClient_QueryText(t *testing.T) {
	// Skip test if ANTHROPIC_API_KEY not set
	if os.Getenv("ANTHROPIC_API_KEY") == "" {
		t.Skip("ANTHROPIC_API_KEY not set")
	}

	client := &AnthropicClient{}

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
			prompt: "",
		},
	}

	tt := tests[0]
	t.Run(tt.name, func(t *testing.T) {
		got, err := client.QueryText(tt.prompt, "claude-3-5-sonnet-latest", Options{})
		if err != nil {
			t.Errorf("AnthropicClient.QueryText() error = %v", err)
			return
		}
		if !strings.Contains(got, "Hello") {
			t.Errorf("AnthropicClient.QueryText() = %v, expected response containing 'Hello'", got)
		}
	})

	tt = tests[1]
	t.Run(tt.name, func(t *testing.T) {
		_, err := client.QueryText(tt.prompt, "claude-3-5-sonnet-latest", Options{})
		if err == nil {
			t.Errorf("AnthropicClient.QueryText() empty prompt should have failed")
			return
		}
	})
}
