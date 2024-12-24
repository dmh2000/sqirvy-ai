// Package main implements the command-line interface for the password generator.
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	pwd "sqirvy.xyz/sqirvy_ai/internal/pwd"
)

const (
	minLength     = 24 // Minimum password length
	defaultLength = 24 // Default password length
	defaultCount  = 1  // Default number of passwords to generate
)

func main() {
	// Define command line flags
	length := flag.Int("length", defaultLength, "Length of password (minimum 24)")
	flag.IntVar(length, "l", defaultLength, "Length of password (minimum 24)")

	count := flag.Int("count", defaultCount, "Number of passwords to generate")
	flag.IntVar(count, "c", defaultCount, "Number of passwords to generate")

	symbols := flag.Bool("symbols", false, "Include symbols in password")
	flag.BoolVar(symbols, "s", false, "Include symbols in password")

	help := flag.Bool("help", false, "Show help message")
	flag.BoolVar(help, "h", false, "Show help message")

	flag.Parse()

	// Handle help flag
	if *help {
		flag.Usage()
		os.Exit(0)
	}

	// Validate minimum length
	if *length < minLength {
		log.Fatalf("Password length must be at least %d characters\n", minLength)
	}

	// Generate passwords
	for i := 0; i < *count; i++ {
		password, err := pwd.GeneratePassword(*length, *symbols)
		if err != nil {
			log.Fatalf("Error generating password: %v\n", err)
		}
		fmt.Println()
		fmt.Println(strings.Repeat("-", 40))
		fmt.Println(password)
		formatted, err := pwd.FormatPassword(password)
		if err != nil {
			log.Fatalf("Error formatting password: %v\n", err)
		}
		fmt.Println(formatted)
	}
}
