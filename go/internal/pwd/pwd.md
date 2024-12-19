# Password Generation Utility

A Go package that provides secure password generation and formatting capabilities.

## Features

- Generates cryptographically secure random passwords
- Enforces a minimum password length of 24 characters
- Optional inclusion of special symbols
- Password formatting with spaces for better readability

## Functions

### `GeneratePassword(length int, includeSymbols bool) (string, error)`

Generates a random password with the following characteristics:
- Minimum length of 24 characters
- Always includes letters (a-z, A-Z) and numbers (0-9)
- Optional special symbols (!@#$%^&*()-_=+[]{}|;:,.<>?)
- Uses crypto/rand for secure random generation

### `FormatPassword(password string) (string, error)`

Formats a password by:
- Adding spaces every 3 characters
- Making the password more readable and easier to transcribe

## Example Usage

```go
// Generate a 30-character password with symbols
password, err := pwd.GeneratePassword(30, true)

// Format the password with spaces
formattedPassword, err := pwd.FormatPassword(password)
```

## Security Notes

- Uses crypto/rand for cryptographically secure random number generation
- Enforces a minimum length of 24 characters for better security
- Provides option to include special symbols for increased entropy

## Error Handling

The package includes comprehensive error handling for:
- Invalid password lengths
- Random number generation failures
- String building operations
