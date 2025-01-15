package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
	util "sqirvyllm/pkg/util"
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

func helpMessage() {
	fmt.Println("Usage: sqirvy-query [options] files...")
	fmt.Println("concatenates prompt from stdin and/or files and sends it to the specified AI model")
	fmt.Println("Options:")
	fmt.Println("  -h    print this help message")
	fmt.Println("  -m    AI model to use (default: claude-3-5-sonnet-latest)")
	fmt.Println("")
	fmt.Println("Supported models:")
	keys := make([]string, 0, len(modelToProvider))
	for key := range modelToProvider {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		fmt.Printf("  %s\n", key)
	}
}

// inputIsFromPipe determines if the program is receiving piped input on stdin
func inputIsFromPipe() (bool, error) {
	fileInfo, err := os.Stdin.Stat()
	return (fileInfo.Mode() & os.ModeCharDevice) == 0, err
}

func processCommandLine() (prompt string, model string, err error) {

	// suppress the default help message
	flag.Usage = func() {}
	// add a -h flag
	var help bool
	flag.BoolVar(&help, "h", false, "print help message")
	flag.StringVar(&model, "m", "", "AI model to use (default: claude-3.5-sonnet-latest)")
	flag.Parse()
	if help {
		helpMessage()
		os.Exit(0)
	}

	var builder strings.Builder
	var totalSize int64

	// Initialize prompt with system.md if it exists
	totalSize = 0
	builder.WriteString("")
	sysprompt, totalSize, err := util.ReadFile("./system.md", MaxTotalBytes)
	if err != nil {
		return "", "", fmt.Errorf("error reading system.md: %w", err)
	}
	builder.WriteString(string(sysprompt))
	builder.WriteString("\n\n")
	totalSize += int64(builder.Len())

	// Check if we have data from stdin
	p, err := inputIsFromPipe()
	if err != nil {
		return "", "", fmt.Errorf("error checking if input is from pipe: %v", err)
	}

	// Read stdin
	var stdinData string
	var stdinSize int64
	if p {
		stdinData, stdinSize, err = util.ReadStdin(MaxTotalBytes)
		if err != nil {
			return "", "", fmt.Errorf("error reading from stdin: %w", err)
		}
	}

	// Check if total size of stdin exceeds MaxTotalBytes
	totalSize += stdinSize
	if totalSize > MaxTotalBytes {
		return "", "", fmt.Errorf("total size would exceed limit of %d bytes", MaxTotalBytes)
	}
	prompt += string(stdinData)

	// Read all files
	fileData, fileSize, err := util.ReadFiles(flag.Args(), MaxTotalBytes)
	if err != nil {
		return "", "", fmt.Errorf("error reading files: %w", err)
	}

	// check if total size of stdin + files exceeds MaxTotalBytes
	totalSize += fileSize
	if totalSize > MaxTotalBytes {
		return "", "", fmt.Errorf("total size would exceed limit of %d bytes", MaxTotalBytes)
	}

	prompt += fileData

	// Check if we have any input
	if prompt == "" {
		return "", "", fmt.Errorf("no input provided via stdin or files")
	}

	return prompt, model, nil
}
