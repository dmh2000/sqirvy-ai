package ai

import (
	"os"
	"strings"
	"testing"
)

func TestAnthropicClient_QueryText(t *testing.T) {
	if os.Getenv("ANTHROPIC_API_KEY") == "" {
		t.Skip("ANTHROPIC_API_KEY not set")
	}

	client := &AnthropicClient{}
	
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
			name:    "Empty prompt",
			prompt:  "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := client.QueryText(tt.prompt)
			if (err != nil) != tt.wantErr {
				t.Errorf("AnthropicClient.QueryText() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !strings.Contains(got, "Hello") {
				t.Errorf("AnthropicClient.QueryText() = %v, expected response containing 'Hello'", got)
			}
		})
	}
}

func TestAnthropicClient_QueryJSON(t *testing.T) {
	if os.Getenv("ANTHROPIC_API_KEY") == "" {
		t.Skip("ANTHROPIC_API_KEY not set")
	}

	client := &AnthropicClient{}
	
	tests := []struct {
		name    string
		prompt  string
		wantErr bool
	}{
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

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := client.QueryJSON(tt.prompt)
			if (err != nil) != tt.wantErr {
				t.Errorf("AnthropicClient.QueryJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !strings.Contains(got, "{") {
				t.Errorf("AnthropicClient.QueryJSON() = %v, expected JSON response", got)
			}
		})
	}
}
