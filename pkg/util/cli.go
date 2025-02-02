// Package main provides command line interface functionality for the AI query tool.
//
// This file contains the command-line processing logic including:
// - Argument parsing
// - Input handling from both files and stdin
// - Size limit enforcement
// - Help message formatting
package util

import (
	"flag"
	"fmt"
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

func CliFlags() (help bool, model string, function string, files []string, err error) {
	// create a dummy help function
	flag.Usage = func() {
		// print the help message
	}

	// suppress the default help message
	// add a -h flag
	var h bool
	var m string
	var f string

	flag.BoolVar(&h, "h", false, "print help message")
	flag.StringVar(&m, "m", "", "AI model to use (default: claude-3.5-sonnet-latest)")
	flag.StringVar(&f, "f", "", "AI function to use (default: generate)")

	flag.Parse()

	// check if help flag is set
	help = h

	// check if model flag is set
	if m != "" {
		model = m
	}

	// check if function flag is set
	if f != "" {
		function = f
	}

	// get the files
	files = flag.Args()

	return help, model, function, files, nil
}

// ReadPrompt reads and concatenates the contents of the given files and stdin,
// input strings are concatenated in the following order:
// 1. ./system.md (if it exists)
// 2. stdin (if it has input on stding, like a pipe)
// 3. files (if any are provided)
func ReadPrompt() (string, error) {
	var builder strings.Builder
	var totalSize int64

	// Initialize prompt with system.md if it exists
	builder.WriteString("")
	sysprompt, totalSize, err := ReadFile("./system.md", MaxTotalBytes)
	if err != nil {
		// no system.md file, skip a system prompt
		totalSize = 0
	}
	if totalSize > 0 {
		builder.WriteString(string(sysprompt))
		builder.WriteString("\n\n")
		totalSize += int64(builder.Len())
	}

	// Check if we have data from stdin
	p, err := InputIsFromPipe()
	if err != nil {
		return "", fmt.Errorf("error checking if input is from pipe: %w", err)
	}

	// Read stdin
	var stdinData string
	var stdinSize int64
	if p {
		stdinData, stdinSize, err = ReadStdin(MaxTotalBytes)
		if err != nil {
			return "", fmt.Errorf("error reading from stdin: %w", err)
		}
	}

	// Check if total size of stdin exceeds MaxTotalBytes
	totalSize += stdinSize
	if totalSize > MaxTotalBytes {
		return "", fmt.Errorf("total size would exceed limit of %d bytes", MaxTotalBytes)
	}
	builder.WriteString(string(stdinData))

	// Read all files
	fileData, fileSize, err := ReadFiles(flag.Args(), MaxTotalBytes)
	if err != nil {
		return "", fmt.Errorf("error reading files: %w", err)
	}

	// check if total size of stdin + files exceeds MaxTotalBytes
	totalSize += fileSize
	if totalSize > MaxTotalBytes {
		return "", fmt.Errorf("total size would exceed limit of %d bytes", MaxTotalBytes)
	}

	builder.WriteString(fileData)
	if builder.Len() == 0 {
		return "", fmt.Errorf("no prompts specified, stdin and files have no data")
	}

	// return the consolidated prompt
	return builder.String(), nil
}
