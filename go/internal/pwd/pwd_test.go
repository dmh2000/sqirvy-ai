// Package pwd provides functions for generating and formatting passwords.
package pwd

import (
	"strings"
	"testing"
)

// TestGeneratePassword tests the GeneratePassword function.
func TestGeneratePassword(t *testing.T) {
	tests := []struct {
		name           string
		length         int
		includeSymbols bool
	}{
		{"basic_24", 24, false},
		{"with_symbols_24", 24, true},
		{"longer_32", 32, false},
		{"longer_with_symbols_32", 32, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			password, err := GeneratePassword(tt.length, tt.includeSymbols)
			if err != nil {
				t.Errorf("GeneratePassword() error = %v", err)
			}

			// Check length
			if len(password) != tt.length {
				t.Errorf("GeneratePassword() length = %v, want %v", len(password), tt.length)
			}

			// Check character types
			hasLetters := strings.ContainsAny(password, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
			hasNumbers := strings.ContainsAny(password, "0123456789")
			hasSymbols := strings.ContainsAny(password, "!@#$%^&*()-_=+[]{}|;:,.<>?")

			if !hasLetters {
				t.Error("Password should contain letters")
			}
			if !hasNumbers {
				t.Error("Password should contain numbers")
			}
			if tt.includeSymbols && !hasSymbols {
				t.Error("Password should contain symbols when includeSymbols is true")
			}
			if !tt.includeSymbols && hasSymbols {
				t.Error("Password should not contain symbols when includeSymbols is false")
			}
		})
	}
}

// TestFormatPassword tests the FormatPassword function.
func TestFormatPassword(t *testing.T) {
	tests := []struct {
		name     string
		password string
		want     string
	}{
		{
			name:     "length_3",
			password: "abc",
			want:     "abc",
		},
		{
			name:     "length_6",
			password: "abcdef",
			want:     "abc def",
		},
		{
			name:     "length_9",
			password: "abcdefghi",
			want:     "abc def ghi",
		},
		{
			name:     "uneven_length_7",
			password: "abcdefg",
			want:     "abc def g",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FormatPassword(tt.password)
			if err != nil {
				t.Errorf("FormatPassword() error = %v", err)
			}

			if got != tt.want {
				t.Errorf("FormatPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}
