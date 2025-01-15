package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
	api "sqirvyllm/pkg/api"
	util "sqirvyllm/pkg/util"
	"strings"
)

const MaxTotalBytes = 262144 // 256KB limit

func helpMessage() {
	fmt.Println("Usage: sqirvy-review [options] files...")
	fmt.Println("concatenates files and sends them to the specified AI model for review")
	fmt.Println("Options:")
	fmt.Println("  -h    print this help message")
	fmt.Println("  -model AI model to use (default: claude-3-5-sonnet-latest)")
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

func processCommandLine() (string, string, error) {
	// suppress the default help message
	flag.Usage = func() {}
	
	var help bool
	var model string

	flag.BoolVar(&help, "h", false, "print help message")
	flag.StringVar(&model, "model", "", "AI model to use")
	flag.Parse()

	if help {
		helpMessage()
		os.Exit(0)
	}

	var builder strings.Builder
	
	// Read all files
	fileData, fileSize, err := util.ReadFiles(flag.Args(), MaxTotalBytes)
	if err != nil {
		return "", "", fmt.Errorf("error reading files: %w", err)
	}

	builder.WriteString(fileData)
	if builder.Len() == 0 {
		return "", "", fmt.Errorf("no files specified or files have no data")
	}

	return builder.String(), model, nil
}
