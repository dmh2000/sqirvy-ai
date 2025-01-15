package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
)

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

	// Initialize prompt with system.md if it exists
	prompt = ""
	if systemData, err := os.ReadFile("./system.md"); err == nil {
		prompt = string(systemData) + "\n\n"
	}

	// Check if we have data from stdin
	p, err := inputIsFromPipe()
	if err != nil {
		return "", "", fmt.Errorf("error checking if input is from pipe: %v", err)
	}
	if p {
		stdinBytes, err := io.ReadAll(os.Stdin)
		if err != nil {
			return "", "", fmt.Errorf("error reading from stdin: %v", err)
		}
		prompt += string(stdinBytes)
	}

	// Process any file arguments. use a buffer to avoid reading the entire file into memory
	for _, filename := range flag.Args() {
		data, err := os.ReadFile(filename)
		if err != nil {
			return "", "", fmt.Errorf("error reading file %s: %v", filename, err)
		}
		prompt += string(data) + "\n\n"
	}

	// Check if we have any input
	if prompt == "" {
		return "", "", fmt.Errorf("no input provided via stdin or files")
	}

	return prompt, model, nil
}

//"hello\n\ngoodbye\n\n"
