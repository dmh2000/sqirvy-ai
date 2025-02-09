// Package main provides command line interface functionality for the AI query tool.
//
// This file contains the command-line processing logic including:
// - Argument parsing
// - Input handling from both files and stdin
// - Size limit enforcement
// - Help message formatting
package main

import (
	"flag"
	"fmt"
	"net/url"
	"os"
	"sort"
	sqirvy "sqirvy-llm/pkg/sqirvy"
	util "sqirvy-llm/pkg/util"
	"strings"
)

// Max Bytes To Read
// MaxTotalBytes is the maximum allowed size for all input files combined
// assume the model will return an error if max context length is exceeded
// it is impractical to know the exact max context length beforehand for all models
// assume max 64k tokens
// assuming 4 characters per token
// total 262144 bytes
// since the bytes are converted to UTF-8, the upload could be larger than the byte limit
const maxTokens = 65536
const bytesPerToken = 4
const MaxTotalBytes = maxTokens * bytesPerToken // 262144 bytes limit

// a list of values that the function flag can take
var validFunctions = map[string]bool{
	"query":  true,
	"plan":   true,
	"code":   true,
	"review": true,
	"scrape": true,
}

func CliFlags() (help bool, model string, function string, temperature int, files []string, err error) {
	// suppress the default help message
	// add a -h flag
	var h bool
	var m string
	var f string
	var t int

	flag.BoolVar(&h, "h", false, "print help message")
	flag.StringVar(&m, "m", "claude-3-5-sonnet-latest", "AI model to use (default: claude-3.5-sonnet-latest)")
	flag.StringVar(&f, "f", "query", "AI function to use (default: query)")
	flag.IntVar(&t, "t", 50, "Temperature setting for the AI model (0-100, default: 50)")

	flag.Parse()

	// check if help flag is set
	help = h

	// check if model flag is set
	if m != "" {
		model = m
	}

	// check if function flag is set
	function = f
	if _, ok := validFunctions[function]; !ok {
		return help, model, function, temperature, files, fmt.Errorf("invalid function: %s", function)
	}

	// check if temperature flag is set
	if t < 1 || t > 100 {
		return help, model, function, temperature, files, fmt.Errorf("invalid temperature: %d", t)
	}
	temperature = t

	// get the files
	files = flag.Args()

	// supress default usage
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s [-h] [-m model] [-f function] [-t temperature] [file1 file2 ...]\n", os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(), "  -h  print this help message\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  -m  AI model to use (default: claude-3.5-sonnet-latest)\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  -f  AI function to use (default: query)\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  -t  Temperature setting for the AI model (0-100, default: 50)\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  valid functions: query, review, code, scrape\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  file1 file2 ...  input files to process\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  URLs can be provided as arguments\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  data from stdin will be read if there is any\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  models:\n")

		// print available models from map keys
		models := make([]string, 0, len(sqirvy.ModelToProvider))
		for k := range sqirvy.ModelToProvider {
			models = append(models, k)
		}
		sort.Strings(models)
		for model := range models {
			fmt.Fprintf(flag.CommandLine.Output(), "    %s\n", models[model])
		}
	}

	return help, model, function, temperature, files, nil
}

// ReadPrompt reads and concatenates the contents of the given files and stdin,
// input strings are concatenated in the following order:
// 1. ./system.md (if it exists)
// 2. stdin (if it has input on stding, like a pipe)
// 3. files (if any are provided)
// 4. URLs (if any are provided)
func ReadPrompt(args []string) (string, error) {
	var builder strings.Builder
	var totalSize int64

	// Initialize prompt with system.md if it exists
	builder.WriteString("")

	// Read stdin
	var stdinData string
	var stdinSize int64
	stdinData, stdinSize, err := util.ReadStdin(MaxTotalBytes)
	if err != nil {
		return "", fmt.Errorf("error reading from stdin: %w", err)
	}

	// Check if total size of stdin exceeds MaxTotalBytes
	totalSize += stdinSize
	if totalSize > MaxTotalBytes {
		return "", fmt.Errorf("total size would exceed limit of %d bytes", MaxTotalBytes)
	}
	builder.WriteString(string(stdinData))

	// Process command line arguments as URLs or files
	for _, arg := range args {
		// Try to read as file first
		fileData, fileSize, err := util.ReadFile(arg, MaxTotalBytes-totalSize)
		if err == nil {
			builder.Write(fileData)
			builder.WriteString("\n\n")
			totalSize += fileSize
			continue
		}

		// If not a file and is a valid url, scrape it
		_, err = url.ParseRequestURI(arg)
		if err == nil {
			content, err := util.ScrapeURL(arg)
			if err != nil {
				return "", fmt.Errorf("failed to process %s as file or URL: %w", arg, err)
			}
			builder.WriteString(content)
			builder.WriteString("\n\n")
			fileSize += int64(len(content))
		}

		// check if total size of stdin + files exceeds MaxTotalBytes
		totalSize += fileSize
		if totalSize > MaxTotalBytes {
			return "", fmt.Errorf("total size would exceed limit of %d bytes", MaxTotalBytes)
		}

	}

	// return the consolidated prompt
	return builder.String(), nil
}
