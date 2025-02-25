// Package cmd implements command-line interface functionality for the sqirvy-cli tool.
// It provides commands for handling various types of prompts and input processing.
package cmd

import (
	_ "embed"
	"fmt"
	"net/url"
	util "sqirvy-ai/pkg/util"
)

// queryPrompt contains the embedded content of the query.md file,
// which defines the system prompt for query operations.
//
//go:embed prompts/query.md
var queryPrompt string

// planPrompt contains the embedded content of the plan.md file,
// which defines the system prompt for planning operations.
//
//go:embed prompts/plan.md
var planPrompt string

// codePrompt contains the embedded content of the code.md file,
// which defines the system prompt for code generation operations.
//
//go:embed prompts/code.md
var codePrompt string

// reviewPrompt contains the embedded content of the review.md file,
// which defines the system prompt for code review operations.
//
//go:embed prompts/review.md
var reviewPrompt string

// ReadPrompt processes input from multiple sources and combines them into a slice of prompts.
// It handles input from:
//   - A base system prompt
//   - Standard input (stdin)
//   - URLs (which are scraped for content)
//   - Local files
//
// The function ensures that the total size of all inputs does not exceed MaxInputTotalBytes.
//
// Parameters:
//   - prompt: The initial system prompt to use
//   - args: A slice of strings that can be either URLs or file paths
//
// Returns:
//   - []string: A slice containing all processed prompts
//   - error: An error if any operation fails or if size limits are exceeded
func ReadPrompt(args []string) ([]string, error) {

	var prompts []string
	var length int64

	// Initialize with the system prompt and check size limit

	// Process standard input and check size limit
	var stdinData string
	stdinData, _, err := util.ReadStdin(MaxInputTotalBytes)
	if err != nil {
		return []string{""}, fmt.Errorf("error: reading from stdin: %w", err)
	}
	prompts = append(prompts, stdinData)
	length += int64(len(stdinData))
	if length > MaxInputTotalBytes {
		return []string{""}, fmt.Errorf("error: total size would exceed limit of %d bytes (stdin)", MaxInputTotalBytes)
	}

	// Process each argument which can be either a URL or a file path
	for _, arg := range args {
		// Attempt to parse argument as URL
		_, err := url.ParseRequestURI(arg)
		if err == nil {
			// Handle URL content
			content, err := util.ScrapeURL(arg)
			if err != nil {
				return []string{""}, fmt.Errorf("error: failed to scrape URL %s: %w", arg, err)
			}
			content += "\n\n"
			prompts = append(prompts, content)
			length += int64(len(content))
			if length > MaxInputTotalBytes {
				return []string{""}, fmt.Errorf("error: total size would exceed limit of %d bytes (urls)", MaxInputTotalBytes)
			}
			continue
		}

		// Handle file content if not a URL
		fileData, _, err := util.ReadFile(arg, MaxInputTotalBytes)
		if err != nil {
			return []string{""}, fmt.Errorf("error: failed to read file %s: %w", arg, err)
		}
		prompts = append(prompts, string(fileData))
		length += int64(len(fileData))
		if length > MaxInputTotalBytes {
			return []string{""}, fmt.Errorf("error: total size would exceed limit of %d bytes (files)", MaxInputTotalBytes)
		}
	}

	// use default prompt if no other ones are specified
	if len(prompts) == 0 || (len(prompts) == 1 && prompts[0] == "") {
		prompts = []string{defaultPrompt}
	}

	return prompts, nil
}
