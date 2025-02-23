package sqirvy

import (
	"context"
	"os"
	"testing"
)

func TestGeminiClient_QueryText(t *testing.T) {
	if os.Getenv("GEMINI_API_KEY") == "" {
		t.Skip("GEMINI_API_KEY not set")
	}

	tests := []struct {
		name    string
		prompt  []string
		wantErr bool
	}{
		{
			name:    "Basic prompt",
			prompt:  []string{"Say 'Hello, World!'"},
			wantErr: false,
		},
		{
			name:    "Empty prompt",
			prompt:  []string{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, err := NewGeminiClient()
			if err != nil {
				t.Errorf("Gemini.QueryText() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil {
				t.Errorf("new client failed")
			}
			got, err := client.QueryText(context.Background(), assistant, tt.prompt, "gemini-1.5-flash", Options{Temperature: 0.5, MaxTokens: 4096})
			if tt.wantErr {
				if err == nil {
					t.Errorf("Gemini.QueryText() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}
			if err != nil {
				t.Errorf("Gemini.QueryText() error = %v", err)
				return
			}
			if len(got) == 0 {
				t.Error("Gemini.QueryText() returned empty response")
			}
		})
	}
}
