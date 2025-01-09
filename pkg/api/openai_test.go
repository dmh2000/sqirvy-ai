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
			got, err := client.QueryText(tt.prompt)
			if (err != nil) != tt.wantErr {
				t.Errorf("OpenAIClient.QueryText() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && len(got) == 0 {
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
			name:    "Empty prompt",
			prompt:  "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := client.QueryJSON(tt.prompt)
			if (err != nil) != tt.wantErr {
				t.Errorf("OpenAIClient.QueryJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if !strings.Contains(got, "{") || !strings.Contains(got, "}") {
					t.Errorf("OpenAIClient.QueryJSON() = %v, expected JSON response", got)
				}
			}
		})
	}
}
