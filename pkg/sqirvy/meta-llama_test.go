package api

import (
	"os"
	"testing"
)

func TestMetaLlamaClient_QueryText(t *testing.T) {
	if os.Getenv("TOGETHER_API_KEY") == "" {
		t.Skip("TOGETHER_API_KEY not set")
	}

	client := &MetaLlamaClient{}

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
			got, err := client.QueryText(tt.prompt, "meta-llama/meta-llama-3.1-8b-instruct-turbo", Options{})
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
