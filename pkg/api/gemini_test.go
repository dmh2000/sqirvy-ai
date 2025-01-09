package ai

import (
	"os"
	"strings"
	"testing"
)

func TestGeminiClient_QueryText(t *testing.T) {
	if os.Getenv("GEMINI_API_KEY") == "" {
		t.Skip("GEMINI_API_KEY not set")
	}

	client := &GeminiClient{}
	
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
				t.Errorf("GeminiClient.QueryText() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && len(got) == 0 {
				t.Error("GeminiClient.QueryText() returned empty response")
			}
		})
	}
}

func TestGeminiClient_QueryJSON(t *testing.T) {
	if os.Getenv("GEMINI_API_KEY") == "" {
		t.Skip("GEMINI_API_KEY not set")
	}

	client := &GeminiClient{}
	
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
				t.Errorf("GeminiClient.QueryJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !strings.Contains(got, "{") {
				t.Errorf("GeminiClient.QueryJSON() = %v, expected JSON response", got)
			}
		})
	}
}
