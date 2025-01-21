package util

import (
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestExtractPDF(t *testing.T) {
	// Skip if pdftotext is not installed
	if _, err := exec.LookPath("pdftotext"); err != nil {
		t.Skip("pdftotext not found in PATH")
	}

	tests := []struct {
		name    string
		pdfPath string
		want    string
		wantErr bool
	}{
		{
			name:    "Valid PDF",
			pdfPath: "README.pdf",
			want:    "Sqirvy",
			wantErr: false,
		},
		{
			name:    "Empty Data",
			pdfPath: "",
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var pdfData []byte
			var err error

			if tt.pdfPath != "" {
				pdfData, err = os.ReadFile(tt.pdfPath)
				if err != nil {
					t.Fatalf("failed to read test PDF: %v", err)
				}
			}

			got, err := ExtractPDF(pdfData)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExtractPDF() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && !strings.Contains(got, tt.want) {
				t.Errorf("ExtractPDF() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSummarizePDF(t *testing.T) {
	// Skip if no API key
	if os.Getenv("ANTHROPIC_API_KEY") == "" {
		t.Skip("ANTHROPIC_API_KEY not set")
	}

	// Skip if pdftotext is not installed
	if _, err := exec.LookPath("pdftotext"); err != nil {
		t.Skip("pdftotext not found in PATH")
	}

	tests := []struct {
		name      string
		pdfPath   string
		maxTokens int
		wantErr   bool
	}{
		{
			name:      "Valid PDF",
			pdfPath:   "README.pdf",
			maxTokens: 1024,
			wantErr:   false,
		},
		{
			name:      "Empty Data",
			pdfPath:   "",
			maxTokens: 1024,
			wantErr:   true,
		},
		{
			name:      "Invalid MaxTokens",
			pdfPath:   "README.pdf",
			maxTokens: 0,
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var pdfData []byte
			var err error

			if tt.pdfPath != "" {
				pdfData, err = os.ReadFile(tt.pdfPath)
				if err != nil {
					t.Fatalf("failed to read test PDF: %v", err)
				}
			}

			got, err := SummarizePDF(pdfData, tt.maxTokens)
			if (err != nil) != tt.wantErr {
				t.Errorf("SummarizePDF() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && len(got) == 0 {
				t.Error("SummarizePDF() returned empty summary")
			}
		})
	}
}
