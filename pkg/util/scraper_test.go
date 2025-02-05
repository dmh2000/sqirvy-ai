package util

import (
	"strings"
	"testing"
)

func TestScrapeURL(t *testing.T) {
	tests := []struct {
		name    string
		url     string
		wantErr bool
		errMsg  string
		want    string
	}{
		{
			name:    "Valid URL - Example.com",
			url:     "https://example.com",
			wantErr: false,
			want:    "Example Domain",
		},
		{
			name:    "Empty URL",
			url:     "",
			wantErr: true,
			errMsg:  "URL cannot be empty",
		},
		{
			name:    "Invalid URL format",
			url:     "not-a-url",
			wantErr: true,
			errMsg:  "failed to scrape",
		},
		{
			name:    "Non-existent domain",
			url:     "https://this-domain-should-not-exist-123456789.com",
			wantErr: true,
			errMsg:  "failed to scrape",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ScrapeURL(tt.url)
			if tt.wantErr {
				if err == nil {
					t.Errorf("ScrapeURL() error = nil, wantErr %v", tt.wantErr)
					return
				}
				if !strings.Contains(err.Error(), tt.errMsg) {
					t.Errorf("ScrapeURL() error = %v, want error containing %v", err, tt.errMsg)
				}
				return
			}
			if err != nil {
				t.Errorf("ScrapeURL() error = %v", err)
				return
			}
			if !strings.Contains(got, tt.want) {
				t.Errorf("ScrapeURL() = %v, want %v", got, tt.want)
			}
			t.Log(got)
		})
	}
}

func TestScrapeAll(t *testing.T) {
	tests := []struct {
		name    string
		urls    []string
		wantErr bool
		errMsg  string
		want    string
	}{
		{
			name: "Multiple valid URLs",
			urls: []string{
				"https://example.com",
				"https://example.org",
			},
			wantErr: false,
			want:    "Example Domain",
		},
		{
			name:    "Empty URL list",
			urls:    []string{},
			wantErr: true,
			errMsg:  "URLs list cannot be empty",
		},
		{
			name: "Mix of valid and invalid URLs",
			urls: []string{
				"https://example.com",
				"not-a-url",
				"https://this-domain-should-not-exist-123456789.com",
			},
			wantErr: true,
			want:    "Example Domain",
		},
		{
			name: "All invalid URLs",
			urls: []string{
				"not-a-url",
				"https://this-domain-should-not-exist-123456789.com",
			},
			wantErr: true,
			errMsg:  "failed to scrape",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ScrapeAll(tt.urls)
			if tt.wantErr {
				if err == nil {
					t.Errorf("ScrapeAll() error = nil, wantErr %v", tt.wantErr)
					return
				}
				if !strings.Contains(err.Error(), tt.errMsg) {
					t.Errorf("ScrapeAll() error = %v, want error containing %v", err, tt.errMsg)
				}
				return
			}
			if err != nil {
				t.Errorf("ScrapeAll() error = %v", err)
				return
			}
			if !strings.Contains(got, tt.want) {
				t.Errorf("ScrapeAll() = %v, want %v", got, tt.want)
			}
		})
	}
}
