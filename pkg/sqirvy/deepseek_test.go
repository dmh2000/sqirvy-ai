package api

import (
	"os"
	"strings"
	"testing"
)

func TestDeepSeekClient_QueryText(t *testing.T) {
	if os.Getenv("DEEPSEEK_API_KEY") == "" {
		t.Skip("DEEPSEEK_API_KEY not set")
	}
	if os.Getenv("DEEPSEEK_API_BASE") == "" {
		t.Skip("DEEPSEEK_API_BASE not set")
	}

	client := &DeepSeekClient{}

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
			got, err := client.QueryText(tt.prompt, "deepseek-chat", Options{})
			if tt.wantErr {
				if err == nil {
					t.Errorf("DeepSeekClient.QueryText() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}
			if err != nil {
				t.Errorf("DeepSeekClient.QueryText() error = %v", err)
				return
			}
			if len(got) == 0 {
				t.Error("DeepSeekClient.QueryText() returned empty response")
			}
		})
	}
}

func TestDeepSeekClient_QueryJSON(t *testing.T) {
	if os.Getenv("DEEPSEEK_API_KEY") == "" {
		t.Skip("DEEPSEEK_API_KEY not set")
	}
	if os.Getenv("DEEPSEEK_API_BASE") == "" {
		t.Skip("DEEPSEEK_API_BASE not set")
	}

	client := &DeepSeekClient{}

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
			got, err := client.QueryJSON(tt.prompt, "deepseek-chat", Options{})
			if tt.wantErr {
				if err == nil {
					t.Errorf("DeepSeekClient.QueryJSON() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}
			if err != nil {
				t.Errorf("DeepSeekClient.QueryJSON() error = %v", err)
				return
			}
			if !strings.Contains(got, "{") || !strings.Contains(got, "}") {
				t.Errorf("DeepSeekClient.QueryJSON() = %v, expected JSON response", got)
			}
		})
	}
}
