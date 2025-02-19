package cmd

import (
	_ "embed"
	"fmt"
	"net/url"
	util "sqirvy-ai/pkg/util"
	"strings"
)

//go:embed prompts/query.md
var queryPrompt string

//go:embed prompts/plan.md
var planPrompt string

//go:embed prompts/code.md
var codePrompt string

//go:embed prompts/review.md
var reviewPrompt string

func ReadPrompt(prompt string, args []string) (string, error) {
	var builder strings.Builder

	// start with the system prompt
	builder.WriteString(prompt)
	if builder.Len() > MaxInputTotalBytes {
		return "", fmt.Errorf("total size would exceed limit of %d bytes (prompt)", MaxInputTotalBytes)
	}

	// STDIN
	var stdinData string
	stdinData, _, err := util.ReadStdin(MaxInputTotalBytes)
	if err != nil {
		return "", fmt.Errorf("error reading from stdin: %w", err)
	}
	builder.WriteString(string(stdinData))
	if builder.Len() > MaxInputTotalBytes {
		return "", fmt.Errorf("total size would exceed limit of %d bytes (stdin)", MaxInputTotalBytes)
	}

	for _, arg := range args {
		// Check if it's a URL
		_, err := url.ParseRequestURI(arg)
		if err == nil {
			content, err := util.ScrapeURL(arg)
			if err != nil {
				return "", fmt.Errorf("failed to scrape URL %s: %w", arg, err)
			}
			builder.WriteString(content)
			builder.WriteString("\n\n")
			if builder.Len() > MaxInputTotalBytes {
				return "", fmt.Errorf("total size would exceed limit of %d bytes (urls)", MaxInputTotalBytes)
			}
			continue
		}

		// If not a URL, try to read as file
		fileData, _, err := util.ReadFile(arg, MaxInputTotalBytes)
		if err != nil {
			return "", fmt.Errorf("failed to read file %s: %w", arg, err)
		}

		builder.Write(fileData)
		builder.WriteString("\n\n")
		if builder.Len() > MaxInputTotalBytes {
			return "", fmt.Errorf("total size would exceed limit of %d bytes (files)", MaxInputTotalBytes)
		}
	}

	return builder.String(), nil
}
