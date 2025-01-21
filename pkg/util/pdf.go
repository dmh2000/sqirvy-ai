// Package util provides utility functions for file operations, web scraping,
// and PDF processing. The PDF functionality includes text extraction and
// AI-powered summarization capabilities.
package util

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	api "sqirvyllm/pkg/sqirvy"
)

// ExtractPDF extracts text content from PDF binary data using pdftotext.
//
// Parameters:
//   - pdfData: []byte containing the raw PDF file data
//
// Returns:
//   - string: Extracted text content from the PDF
//   - error: Returns error if:
//   - pdfData is empty
//   - pdftotext command fails
//   - text extraction fails
//   - temporary file operations fail
//
// Example usage:
//
// pdfData, err := os.ReadFile("document.pdf")
//
//	if err != nil {
//	   log.Fatal(err)
//	}
//
// text, err := ExtractPDF(pdfData)
//
//	if err != nil {
//	   log.Fatal(err)
//	}
//
// fmt.Println(text)
func ExtractPDF(pdfData []byte) (string, error) {
	// Validate input
	if len(pdfData) == 0 {
		return "", fmt.Errorf("PDF data cannot be empty")
	}

	// Create temporary directory for processing
	tempDir, err := os.MkdirTemp("", "pdf_extract_*")
	if err != nil {
		return "", fmt.Errorf("failed to create temporary directory: %w", err)
	}
	defer os.RemoveAll(tempDir) // Clean up temp directory

	// Create temporary PDF file
	pdfPath := filepath.Join(tempDir, "input.pdf")
	if err := os.WriteFile(pdfPath, pdfData, 0600); err != nil {
		return "", fmt.Errorf("failed to write temporary PDF file: %w", err)
	}

	// Create output path for text
	txtPath := filepath.Join(tempDir, "output.txt")

	// Execute pdftotext command
	cmd := exec.Command("pdftotext", "-layout", pdfPath, txtPath)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("pdftotext failed: %w: %s", err, stderr.String())
	}

	// Read extracted text
	content, err := os.ReadFile(txtPath)
	if err != nil {
		return "", fmt.Errorf("failed to read extracted text: %w", err)
	}

	return string(content), nil
}

// SummarizePDF generates an AI-powered summary of PDF content using the sqirvy library.
//
// Parameters:
//   - pdfData: []byte containing the raw PDF file data
//   - maxTokens: Maximum number of tokens for the summary (must be > 0)
//
// Returns:
//   - string: AI-generated summary of the PDF content
//   - error: Returns error if:
//   - pdfData is empty
//   - maxTokens is invalid
//   - text extraction fails
//   - AI summarization fails
//
// Example usage:
//
// pdfData, err := os.ReadFile("document.pdf")
//
//	if err != nil {
//	   log.Fatal(err)
//	}
//
// summary, err := SummarizePDF(pdfData, 1024)
//
//	if err != nil {
//	   log.Fatal(err)
//	}
//
// fmt.Println(summary)
func SummarizePDF(pdfData []byte, maxTokens int) (string, error) {
	// Validate input
	if len(pdfData) == 0 {
		return "", fmt.Errorf("PDF data cannot be empty")
	}
	if maxTokens <= 0 {
		return "", fmt.Errorf("maxTokens must be greater than 0")
	}

	// Extract text from PDF
	text, err := ExtractPDF(pdfData)
	if err != nil {
		return "", fmt.Errorf("failed to extract PDF text: %w", err)
	}

	// Initialize AI client
	client, err := api.NewClient(api.Anthropic)
	if err != nil {
		return "", fmt.Errorf("failed to initialize AI client: %w", err)
	}

	// Create summarization prompt
	prompt := fmt.Sprintf("Please provide a concise summary of the following text, focusing on the main points and key information:\n\n%s", text)

	// Get summary from AI
	summary, err := client.QueryText(prompt, "claude-3-sonnet-20240229", api.Options{})
	if err != nil {
		return "", fmt.Errorf("failed to generate summary: %w", err)
	}

	return summary, nil
}
