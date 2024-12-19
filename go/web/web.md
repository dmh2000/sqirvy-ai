# Password Generator Web Service

A Go-based web service that generates secure passwords with customizable parameters. The service provides both a web interface and JSON API endpoints for password generation.

## Features

- Generate one or multiple passwords
- Customizable password length (minimum 24 characters)
- Optional symbol inclusion
- Returns both raw and formatted password versions
- Web interface and API support
- Static file serving with embedded resources

## API Endpoints

### GET /

Serves the main web interface

### GET /generate

Generates passwords based on query parameters:

- `length`: Password length (≥24, defaults to 24)
- `count`: Number of passwords to generate (≥1, defaults to 1)
- `symbols`: Include symbols (true/false)

Response format:

```json
{
  "passwords": [
    {
      "raw": "raw_password_string",
      "formatted": "formatted_password_string"
    }
  ]
}
```

## Usage

The server runs on `http://localhost:8080` by default. Access the web interface by opening this URL in your browser, or make API calls to the `/generate` endpoint.

## Dependencies

- Standard Go libraries
- Custom password generation package (`sqirvy.xyz/sqirvy_ai/internal/pwd`)

## Installation

1. Ensure Go is installed on your system
2. Clone the repository
3. Run `go run main.go`
4. Access the service at `http://localhost:8080`
