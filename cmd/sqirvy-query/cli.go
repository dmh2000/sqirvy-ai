package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func processCommandLine() (prompt string, model string, err error) {
	flag.StringVar(&model, "model", "", "AI model to use (default: claude-3.5-sonnet)")
	flag.Parse()

	// Check if we have data from stdin
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		stdinBytes, err := io.ReadAll(os.Stdin)
		if err != nil {
			return "", "", fmt.Errorf("error reading from stdin: %v", err)
		}
		prompt += string(stdinBytes)
	}

	// Process any file arguments
	for _, filename := range flag.Args() {
		data, err := os.ReadFile(filename)
		if err != nil {
			return "", "", fmt.Errorf("error reading file %s: %v", filename, err)
		}
		prompt += string(data)
	}

	// Check if we have any input
	if prompt == "" {
		return "", "", fmt.Errorf("no input provided via stdin or files")
	}

	return prompt, model, nil
}
