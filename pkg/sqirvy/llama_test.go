package sqirvy

import (
	"context"
	"os"
	"testing"
)

func TestMetaLlamaClient_QueryText(t *testing.T) {
	if os.Getenv("LLAMA_API_KEY") == "" {
		t.Skip("LLAMA_API_KEY not set")
	}
	if os.Getenv("LLAMA_BASE_URL") == "" {
		t.Skip("LLAMA_BASE_URL not set")
	}

	client, err := NewClient(Provider("llama"))
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

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
			got, err := client.QueryText(context.Background(), tt.prompt, "llama3.3-70b", Options{})
			if tt.wantErr {
				if err == nil {
					t.Errorf("MetaLlamaClient.QueryText() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}
			if err != nil {
				t.Errorf("MetaLlamaClient.QueryText() error = %v", err)
				return
			}
			if len(got) == 0 {
				t.Error("MetaLlamaClient.QueryText() returned empty response")
			}
		})
	}
}
