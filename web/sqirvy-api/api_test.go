package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestModelsEndpoint(t *testing.T) {
	// Create a test server
	ts := httptest.NewServer(http.HandlerFunc(handleModels))
	defer ts.Close()

	// Make GET request
	resp, err := http.Get(ts.URL)
	if err != nil {
		t.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	// Check status code
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK; got %v", resp.StatusCode)
	}

	// Decode response
	var response ModelsResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	// Verify we got some models
	if len(response.Models) == 0 {
		t.Error("Expected models in response, got none")
	}

	// Check for expected models
	expectedModels := map[string]string{
		"claude-3-5-sonnet-latest": "anthropic",
		"gemini-1.5-pro":           "gemini",
		"gpt-4-turbo":              "openai",
	}

	for model, provider := range expectedModels {
		found := false
		for _, m := range response.Models {
			if m.Model == model && m.Provider == provider {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Expected to find model %s from provider %s", model, provider)
		}
	}
}

func TestQueryEndpoint(t *testing.T) {
	// Skip if no API keys are set
	if os.Getenv("ANTHROPIC_API_KEY") == "" &&
		os.Getenv("GEMINI_API_KEY") == "" &&
		os.Getenv("OPENAI_API_KEY") == "" {
		t.Skip("Skipping test as no API keys are set")
	}

	// Create a test server
	ts := httptest.NewServer(http.HandlerFunc(handleQuery))
	defer ts.Close()

	tests := []struct {
		name        string
		request     QueryRequest
		wantStatus  int
		wantErr     bool
		skipIfNoKey string
	}{
		{
			name: "Valid Anthropic Query",
			request: QueryRequest{
				Model:       "claude-3-5-sonnet-latest",
				Prompt:      "Say hello",
				Temperature: 50,
			},
			wantStatus:  http.StatusOK,
			skipIfNoKey: "ANTHROPIC_API_KEY",
		},
		{
			name: "Valid Gemini Query",
			request: QueryRequest{
				Model:       "gemini-1.5-pro",
				Prompt:      "Say hello",
				Temperature: 50,
			},
			wantStatus:  http.StatusOK,
			skipIfNoKey: "GEMINI_API_KEY",
		},
		{
			name: "Valid OpenAI Query",
			request: QueryRequest{
				Model:       "gpt-4-turbo",
				Prompt:      "Say hello",
				Temperature: 50,
			},
			wantStatus:  http.StatusOK,
			skipIfNoKey: "OPENAI_API_KEY",
		},
		{
			name: "Empty Prompt",
			request: QueryRequest{
				Model:       "claude-3-5-sonnet-latest",
				Prompt:      "",
				Temperature: 50,
			},
			wantStatus: http.StatusBadRequest,
			wantErr:    true,
		},
		{
			name: "Invalid Model",
			request: QueryRequest{
				Model:       "invalid-model",
				Prompt:      "Say hello",
				Temperature: 50,
			},
			wantStatus: http.StatusBadRequest,
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Skip if required API key is not set
			if tt.skipIfNoKey != "" && os.Getenv(tt.skipIfNoKey) == "" {
				t.Skipf("Skipping test as %s is not set", tt.skipIfNoKey)
			}

			// Create request body
			body, err := json.Marshal(tt.request)
			if err != nil {
				t.Fatalf("Failed to marshal request: %v", err)
			}

			// Make POST request
			s := string(body)
			t.Log(s)
			resp, err := http.Post(ts.URL, "application/json", bytes.NewBuffer(body))
			if err != nil {
				t.Fatalf("Failed to make request: %v", err)
			}
			defer resp.Body.Close()

			// Check status code
			if resp.StatusCode != tt.wantStatus {
				t.Errorf("Expected status %v; got %v", tt.wantStatus, resp.StatusCode)
			}

			// For successful requests, verify response structure
			if !tt.wantErr {
				var response QueryResponse
				if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
					t.Fatalf("Failed to decode response: %v:%v", s, err)
				}

				if response.Result == "" {
					t.Error("Expected non-empty result in response")
				}
			}
		})
	}
}
