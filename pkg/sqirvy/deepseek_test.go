package sqirvy

import (
	"context"
	"os"
	"testing"
)

func TestDeepSeekClient_Query_R1(t *testing.T) {
	if os.Getenv("DEEPSEEK_API_KEY") == "" {
		t.Skip("DEEPSEEK_API_KEY not set")
	}
	if os.Getenv("DEEPSEEK_BASE_URL") == "" {
		t.Skip("DEEPSEEK_BASE_URL not set")
	}

	client, err := NewDeepSeekClient()
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
			prompt:  []string{""},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := client.QueryText(context.Background(), tt.prompt, "deepseek-r1", Options{})
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

// func TestDeepSeekClient_Query_V3(t *testing.T) {
// 	if os.Getenv("DEEPSEEK_API_KEY") == "" {
// 		t.Skip("DEEPSEEK_API_KEY not set")
// 	}
// 	if os.Getenv("DEEPSEEK_BASE_URL") == "" {
// 		t.Skip("DEEPSEEK_BASE_URL not set")
// 	}

// 	client := &DeepSeekClient{}

// 	tests := []struct {
// 		name    string
// 		prompt  string
// 		wantErr bool
// 	}{
// 		{
// 			name:    "Basic prompt",
// 			prompt:[]string{ "Say 'Hello, World!'",
// 			wantErr: false,
// 		},
// 		{
// 			name:    "Empty prompt",
// 			prompt:[]string{ "",
// 			wantErr: true,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, err := client.QueryText(tt.prompt, "deepseek-v3", Options{})
// 			if tt.wantErr {
// 				if err == nil {
// 					t.Errorf("DeepSeekClient.QueryText() error = %v, wantErr %v", err, tt.wantErr)
// 				}
// 				return
// 			}
// 			if err != nil {
// 				t.Errorf("DeepSeekClient.QueryText() error = %v", err)
// 				return
// 			}
// 			if len(got) == 0 {
// 				t.Error("DeepSeekClient.QueryText() returned empty response")
// 			}
// 		})
// 	}
// }
