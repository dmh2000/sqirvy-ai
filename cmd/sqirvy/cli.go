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
	sqirvy "sqirvy-ai/pkg/sqirvy"
	util "sqirvy-ai/pkg/util"
	"strings"
)

const (
	maxTokens     = 65536
	bytesPerToken = 4
	MaxTotalBytes = maxTokens * bytesPerToken // 262144 bytes limit

	flagHelp        = "h"
	flagModel       = "m"
	flagFunction    = "f"
	flagTemperature = "t"
)

var validFunctions = map[string]bool{
	"query":  true,
	"plan":   true,
	"code":   true,
	"review": true,
	"scrape": true,
}

func CliFlags() (help bool, model string, function string, temperature int, files []string, err error) {
	var h bool
	var m string
	var f string
	var t int

	flag.BoolVar(&h, flagHelp, false, "print help message")
	flag.StringVar(&m, flagModel, "claude-3-5-sonnet-latest", "AI model to use (default: claude-3.5-sonnet-latest)")
	flag.StringVar(&f, flagFunction, "query", "AI function to use (default: query)")
	flag.IntVar(&t, flagTemperature, 50, "Temperature setting for the AI model (0-100, default: 50)")

	flag.Parse()

	help = h

	if m != "" {
		model = m
	}

	function = f
	if _, ok := validFunctions[function]; !ok {
		return help, model, function, temperature, files, fmt.Errorf("invalid function: %s", function)
	}

	if t < 1 || t > 100 {
		return help, model, function, temperature, files, fmt.Errorf("invalid temperature: %d", t)
	}
	temperature = t

	files = flag.Args()

	flag.Usage = printUsage

	return help, model, function, temperature, files, nil
}

func printUsage() {
	fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s [-h] [-m model] [-f function] [-t temperature] [files and/or urls  ...]\n", os.Args[0])
	fmt.Fprintf(flag.CommandLine.Output(), "  -%s  print this help message\n", flagHelp)
	fmt.Fprintf(flag.CommandLine.Output(), "  -%s  AI model to use (default: claude-3.5-sonnet-latest)\n", flagModel)
	fmt.Fprintf(flag.CommandLine.Output(), "  -%s  AI function to use (default: query)\n", flagFunction)
	fmt.Fprintf(flag.CommandLine.Output(), "    -%s query  : execute a generic query\n", flagFunction)
	fmt.Fprintf(flag.CommandLine.Output(), "    -%s plan   : generate a plan for further action\n", flagFunction)
	fmt.Fprintf(flag.CommandLine.Output(), "    -%s code   : generate code according to input specifications\n", flagFunction)
	fmt.Fprintf(flag.CommandLine.Output(), "    -%s review : review code or text\n", flagFunction)
	fmt.Fprintf(flag.CommandLine.Output(), "    -%s scrape : scrape a URL for text\n", flagFunction)
	fmt.Fprintf(flag.CommandLine.Output(), "  -%s  Temperature setting for the AI model (0-100, default: 50)\n", flagTemperature)
	fmt.Fprintf(flag.CommandLine.Output(), "  file1 file2 ...  input files to process\n")
	fmt.Fprintf(flag.CommandLine.Output(), "  URLs can be provided as arguments\n")
	fmt.Fprintf(flag.CommandLine.Output(), "  data from stdin will be read if there is any\n")
	fmt.Fprintf(flag.CommandLine.Output(), "  models:\n")

	models := make([]string, 0, len(sqirvy.ModelToProvider))
	for k := range sqirvy.ModelToProvider {
		models = append(models, k)
	}
	sort.Strings(models)
	for _, model := range models {
		fmt.Fprintf(flag.CommandLine.Output(), "    %s\n", model)
	}
}

func ReadPrompt(args []string) (string, error) {
	var builder strings.Builder
	var totalSize int64

	var stdinData string
	var stdinSize int64
	stdinData, stdinSize, err := util.ReadStdin(MaxTotalBytes)
	if err != nil {
		return "", fmt.Errorf("error reading from stdin: %w", err)
	}

	totalSize += stdinSize
	if totalSize > MaxTotalBytes {
		return "", fmt.Errorf("total size would exceed limit of %d bytes", MaxTotalBytes)
	}
	builder.WriteString(string(stdinData))

	for _, arg := range args {
		// Check if it's a URL
		_, err := url.ParseRequestURI(arg)
		if err == nil {
			content, err := util.ScrapeURL(arg)
			if err != nil {
				return "", fmt.Errorf("failed to scrape URL %s: %w", arg, err)
			}
			builder.WriteString(content)
			builder.WriteString("\n\n")
			totalSize += int64(len(content)) // Update totalSize with URL content length
			if totalSize > MaxTotalBytes {
				return "", fmt.Errorf("total size would exceed limit of %d bytes", MaxTotalBytes)
			}
			continue
		}

		// If not a URL, try to read as file
		fileData, fileSize, err := util.ReadFile(arg, MaxTotalBytes)
		if err != nil {
			return "", fmt.Errorf("failed to read file %s: %w", arg, err)
		}

		builder.Write(fileData)
		builder.WriteString("\n\n")
		totalSize += int64(fileSize) // Update totalSize with file size
		if totalSize > MaxTotalBytes {
			return "", fmt.Errorf("total size would exceed limit of %d bytes", MaxTotalBytes)
		}
	}

	return builder.String(), nil
}
