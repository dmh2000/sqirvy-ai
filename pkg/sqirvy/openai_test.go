package sqirvy

import (
	"context"
	"os"
	"testing"
)

func TestOpenAIClient_QueryText(t *testing.T) {
	if os.Getenv("OPENAI_API_KEY") == "" {
		t.Skip("OPENAI_API_KEY not set")
	}

	client, err := NewOpenAIClient()
	if err != nil {
		t.Errorf("new client failed")
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
			got, err := client.QueryText(context.Background(), "You are a helpful AI assistant.", tt.prompt, "gpt-4-turbo", Options{MaxTokens: GetMaxTokens("gpt-4-turbo")})
			if tt.wantErr {
				if err == nil {
					t.Errorf("OpenAIClient.QueryText() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}
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
