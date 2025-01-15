// Package main provides command line interface functionality for the code review tool.
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

// helpMessage prints usage information for the command line tool.
// If prefix is not empty, it is printed before the usage information.
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

// processCommandLine parses command line arguments and returns:
// - the concatenated content of all input files as a string
// - the specified AI model name (or empty string for default)
// - any error that occurred during processing
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
