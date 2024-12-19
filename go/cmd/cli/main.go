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
	length := flag.Int("l", defaultLength, "Length of password (minimum 20)")
	lengthLong := flag.Int("length", defaultLength, "Length of password (minimum 20)")
	count := flag.Int("c", defaultCount, "Number of passwords to generate")
	countLong := flag.Int("count", defaultCount, "Number of passwords to generate")
	symbols := flag.Bool("s", false, "Include symbols in password")
	symbolsLong := flag.Bool("symbols", false, "Include symbols in password")
	help := flag.Bool("h", false, "Show help message")
	helpLong := flag.Bool("help", false, "Show help message")

	flag.Parse()

	// Handle help flag
	if *help || *helpLong {
		flag.Usage()
		os.Exit(0)
	}

	// Use the long version if it's set, otherwise use short version
	passwordLength := *length
	if flag.Lookup("length").Value.String() != fmt.Sprint(defaultLength) {
		passwordLength = *lengthLong
	}

	passwordCount := *count
	if flag.Lookup("count").Value.String() != fmt.Sprint(defaultCount) {
		passwordCount = *countLong
	}

	// Validate minimum length
	if passwordLength < minLength {
		log.Fatalf("Password length must be at least %d characters\n", minLength)
	}

	// Generate passwords
	for i := 0; i < passwordCount; i++ {
		password, err := pwd.GeneratePassword(passwordLength, *symbols || *symbolsLong)
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
