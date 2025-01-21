package main

import (
	"flag"
	"fmt"
	"os"
	util "sqirvyllm/pkg/util"
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

type CommandLine struct {
	Model     string
	Files     []string
	StdinData string
}

func helpMessage() {
	fmt.Println("Usage: sqirvy-query [options] files...")
}

func processCommandLine() (string, string, []string, error) {
	// suppress the default help message
	flag.Usage = func() {}
	// add a -h flag
	var help bool
	var model string

	// create command line flags
	flag.BoolVar(&help, "h", false, "print help message")
	flag.StringVar(&model, "m", "", "AI model to use (default: claude-3.5-sonnet-latest)")
	flag.Parse()
	if help {
		helpMessage()
		os.Exit(0)
	}

	// get remaining arguments as filenames
	files := flag.Args()

	// Check if we have data from stdin
	p, err := util.InputIsFromPipe()
	if err != nil {
		return "", "", nil, fmt.Errorf("error checking if input is from pipe: %w", err)
	}

	// Read stdin
	var stdinData string
	var stdinSize int64
	if p {
		stdinData, stdinSize, err = util.ReadStdin(MaxTotalBytes)
		if err != nil {
			return "", "", nil, fmt.Errorf("error reading from stdin: %w", err)
		}
	}

	if stdinSize > MaxTotalBytes {
		return "", "", nil, fmt.Errorf("stdin data exceeds maximum size of %d bytes: %d", MaxTotalBytes, stdinSize)
	}

	return model, stdinData, files, nil
}
