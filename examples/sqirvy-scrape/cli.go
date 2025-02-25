// Package main provides command line interface functionality for web scraping.
//
// This package handles:
// - Command line argument parsing and validation
// - File reading and content aggregation
// - Web scraping of provided URLs
// - Size limit enforcement for API requests
// - Help message formatting and display
// - Model selection and validation
package main

import (
	"flag"
	"fmt"
	"sort"
	sqirvy "sqirvy-ai/pkg/sqirvy"
	util "sqirvy-ai/pkg/util"
	"strings"
)

const MaxTotalBytes = 262144 // 256KB limit

// helpMessage displays usage information for the command line tool.
func helpMessage(prefix string) {
	if prefix != "" {
		fmt.Println(prefix)
	}
	fmt.Println("Usage: sqirvy-scrape [options] urls...")
	fmt.Println("initializes the context from stdin, pipe or redirection (if any)")
	fmt.Println("scrapes content from URLs and sends it to the specified AI model")
	fmt.Println("Options:")
	fmt.Println("  -h     print this help message")
	fmt.Println("  -m     AI model to use (default: claude-3-5-sonnet-latest)")
	fmt.Println("")
	fmt.Println("Supported models:")
	keys := sqirvy.GetModelList()
	sort.Strings(keys)
	for _, key := range keys {
		fmt.Printf("  %s\n", key)
	}
}

// processCommandLine handles command line argument processing and input handling.
func processCommandLine() (string, string, error) {
	// Suppress default help message
	flag.Usage = func() {}

	var help bool
	var model string

	flag.BoolVar(&help, "h", false, "print help message")
	flag.StringVar(&model, "m", "", "AI model to use")
	flag.Parse()

	if help {
		helpMessage("")
		return "help", "", nil
	}

	var builder strings.Builder
	var totalSize int64

	// Read system.md if it exists
	sysprompt, size, err := util.ReadFile("./system.md", MaxTotalBytes)
	if err == nil && size > 0 {
		builder.Write(sysprompt)
		builder.WriteString("\n\n")
		totalSize += int64(builder.Len())
	}

	// Check for stdin input
	isPipe, err := util.InputIsFromPipe()
	if err != nil {
		return "", "", fmt.Errorf("error checking stdin: %w", err)
	}

	if isPipe {
		stdinData, stdinSize, err := util.ReadStdin(MaxTotalBytes - totalSize)
		if err != nil {
			return "", "", fmt.Errorf("error reading stdin: %w", err)
		}
		builder.WriteString(stdinData)
		builder.WriteString("\n\n")
		totalSize += stdinSize
	}

	// Process command line arguments as URLs or files
	for _, arg := range flag.Args() {
		// Try to read as file first
		fileData, fileSize, err := util.ReadFile(arg, MaxTotalBytes-totalSize)
		if err == nil {
			builder.Write(fileData)
			builder.WriteString("\n\n")
			totalSize += fileSize
			continue
		}

		// If not a file, try to scrape as URL
		content, err := util.ScrapeURL(arg)
		if err != nil {
			return "", "", fmt.Errorf("failed to process %s as file or URL: %w", arg, err)
		}
		builder.WriteString(content)
		builder.WriteString("\n\n")
	}

	// Check if we have any content
	if builder.Len() == 0 {
		return "", "", fmt.Errorf("no content from files, URLs, or stdin")
	}

	return builder.String(), model, nil
}
