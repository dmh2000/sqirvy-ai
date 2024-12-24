// Package pwd provides functions for generating and formatting passwords.
package pwd

import (
	"crypto/rand"
	"fmt"
	"strings"
)

// minimum length of password is 24 characters
const minLength = 24

// GeneratePassword generates a random password of the specified length.
// It includes letters and numbers by default, and optionally includes symbols if includeSymbols is true.
func GeneratePassword(length int, includeSymbols bool) (string, error) {
	if length < minLength {
		return "", fmt.Errorf("password length must be at least %d characters", minLength)
	}

	const (
		letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
		numbers = "0123456789"
		symbols = "!@#$%^&*()-_=+[]{}|;:,.<>?"
	)

	// Create character set
	charset := letters + numbers
	if includeSymbols {
		charset += symbols
	}

	// Generate password
	password := make([]byte, length)
	charsetLen := len(charset)

    randomBytes := make([]byte, length)
    _, err := rand.Read(randomBytes)
    if err != nil {
        return "", fmt.Errorf("error reading random bytes: %w", err)
    }
    for i := 0; i < length; i++ {
        password[i] = charset[int(randomBytes[i])%charsetLen]
    }

	return string(password), nil
}

// FormatPassword takes a password string and formats it by adding spaces every 3 characters.
func FormatPassword(password string) (string, error) {
	var builder strings.Builder
	for i := 0; i < len(password); i += 3 {
		if i > 0 {
			_, err := builder.WriteString(" ")
			if err != nil {
				return "", fmt.Errorf("error writing space to builder: %w", err)
			}
		}
		end := i + 3
		if end > len(password) {
			end = len(password)
		}
		_, err := builder.WriteString(password[i:end])
		if err != nil {
			return "", fmt.Errorf("error writing password to builder: %w", err)
		}
	}
	return builder.String(), nil
}
