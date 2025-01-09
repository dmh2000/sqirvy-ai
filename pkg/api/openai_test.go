package ai

import (
	"os"
	"strings"
	"testing"
)

func TestOpenAIClient_QueryText(t *testing.T) {
	if os.Getenv("OPENAI_API_KEY") == "" {
		t.Skip("OPENAI_API_KEY not set")
	}

	client := &OpenAIClient{}

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
			got, err := client.QueryText(tt.prompt, "gpt-4-turbo-preview", Options{})
			if err != nil {
				t.Errorf("OpenAIClient.QueryText() error = %v", err)
				return
			}
			if len(got) == 0 {
				t.Error("OpenAIClient.QueryText() returned empty response")
			}
		})
	}
}

func TestOpenAIClient_QueryJSON(t *testing.T) {
	if os.Getenv("OPENAI_API_KEY") == "" {
		t.Skip("OPENAI_API_KEY not set")
	}

	client := &OpenAIClient{}

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
			name:    "JSON Empty prompt",
			prompt:  "",
			wantErr: true,
		},
	}

	tt := tests[0]
	t.Run(tt.name, func(t *testing.T) {
		got, err := client.QueryJSON(tt.prompt, "gpt-4-turbo-preview", Options{})
		if err != nil {
			t.Errorf("OpenAIClient.QueryJSON() error = %v", err)
			return
		}
		if !strings.Contains(got, "{") || !strings.Contains(got, "}") {
			t.Errorf("OpenAIClient.QueryJSON() = %v, expected JSON response", got)
		}
	})
	// empty prompt should fail
	tt = tests[1]
	t.Run(tt.name, func(t *testing.T) {
		_, err := client.QueryJSON(tt.prompt, "gpt-4-turbo-preview", Options{})
		if err == nil {
			t.Errorf("OpenAIClient.QueryJSON() : should have failed due to no 'json' in prompt")
			return
		}
	})
}
