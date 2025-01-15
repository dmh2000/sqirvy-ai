// Package main provides command line interface functionality for the code review tool.
//
// This package handles:
// - Command line argument parsing and validation
// - File reading and content aggregation
// - Size limit enforcement for API requests
// - Help message formatting and display
// - Model selection and validation
//
// The CLI supports reviewing multiple files in a single pass while ensuring the
// total content size stays within the AI model's context limits.
package main

import (
	"flag"
	"fmt"
	"sort"
	api "sqirvyllm/pkg/api"
	util "sqirvyllm/pkg/util"
	"strings"
)

const MaxTotalBytes = 262144 // 256KB limit

// helpMessage displays comprehensive usage information for the command line tool.
//
// Parameters:
//   - prefix: Optional message to display before the usage information.
//     Typically used for error messages or warnings.
//
// The help message includes:
// - Basic usage syntax
// - Available command line options
// - List of supported AI models
// - Size limits and other constraints
func helpMessage(prefix string) {
	if prefix != "" {
		fmt.Println(prefix)
	}
	fmt.Println("Usage: sqirvy-review [options] files...")
	fmt.Println("adds files to context and sends them to the specified AI model for review")
	fmt.Println("Options:")
	fmt.Println("  -h    print this help message")
	fmt.Println("  -model AI model to use (default: gemini-1.5-flash)")
	fmt.Println("")
	fmt.Println("Supported models:")
	keys := make([]string, 0, len(api.ModelToProvider))
	for key := range api.ModelToProvider {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		fmt.Printf("  %s\n", key)
	}
}

// processCommandLine handles command line argument processing and input file handling.
//
// The function:
// - Parses and validates command line flags
// - Processes help requests
// - Reads and concatenates input files
// - Enforces size limits (MaxTotalBytes)
// - Validates model selection
//
// Returns:
//   - string: Concatenated content from all input files
//   - string: Selected AI model name (or empty for default)
//   - error: Any error encountered during processing
//
// Size limits are enforced to ensure the total input stays within
// the AI model's context window (default 256KB).
// processCommandLine parses command line arguments and assembles the input prompt.
// It handles input from both files and stdin, enforcing size limits.
//
// Returns:
//   - string: The assembled prompt text
//   - string: The specified AI model name (or empty for default)
//   - error: Any error that occurred during processing
func processCommandLine() (string, string, error) {
	// suppress the default help message
	flag.Usage = func() {}

	var help bool
	var model string

	flag.BoolVar(&help, "h", false, "print help message")
	flag.StringVar(&model, "model", "", "AI model to use")
	flag.Parse()

	if help {
		helpMessage("")
		return "help", "", nil
	}

	var builder strings.Builder

	// Read all files (will return error if MaxTotalBytes exceeded)
	fileData, _, err := util.ReadFiles(flag.Args(), MaxTotalBytes)
	if err != nil {
		return "", "", fmt.Errorf("error reading files: %w", err)
	}

	builder.WriteString(fileData)
	if builder.Len() == 0 {
		return "", "", fmt.Errorf("no files specified or files have no data")
	}

	return builder.String(), model, nil
}
