package sqirvy

import (
	"context"
	"os"
	"strings"
	"testing"
)

const assistant = "you are a helpful assistant"

func TestAnthropicClient_QueryText(t *testing.T) {
	// Skip test if ANTHROPIC_API_KEY not set
	if os.Getenv("ANTHROPIC_API_KEY") == "" {
		t.Skip("ANTHROPIC_API_KEY not set")
	}

	client, err := NewAnthropicClient()
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

	tt := tests[0]
	t.Run(tt.name, func(t *testing.T) {
		ctx := context.Background()
		got, err := client.QueryText(ctx, assistant, tt.prompt, "claude-3-5-sonnet-latest", Options{})
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
		response, err := client.QueryText(context.Background(), assistant, tt.prompt, "claude-3-5-sonnet-latest", Options{})
		if err == nil {
			t.Errorf("AnthropicClient.QueryText() empty prompt should have failed")
			return
		}
		t.Log(response)
	})
}
