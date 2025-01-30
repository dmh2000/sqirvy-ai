# Sqirvy: A Multi-Provider AI Interaction Toolkit

Sqirvy is a versatile toolkit designed to facilitate interactions with various AI language models from different providers. It provides a unified interface for querying models, handling authentication, and managing responses. The toolkit is structured into several key components: command-line tools, a core API library, utility functions, and a web interface.

## Overview of Components

### `cmd` Directory

The `cmd` directory contains the source code for command-line tools that leverage the core API. Each subdirectory within `cmd` represents a specific tool or provider integration.

-   **`cmd/sqirvy-query`**: A general-purpose tool for querying AI models. It accepts input from files and stdin, concatenates them, and sends the combined text as a prompt to the specified AI model.
-   **`cmd/sqirvy-review`**: An automated code review tool that uses AI language models to analyze source code and provide feedback. It supports multiple AI providers and uses embedded system prompts for consistent reviews.
-   **`cmd/sqirvy-scrape`**: A web scraping tool that extracts content from URLs and sends it to an AI model for analysis. It supports multiple AI providers and models.
-   **`cmd/anthropic`**: A simple example of using the Anthropic client directly.
-   **`cmd/gemini`**: A simple example of using the Gemini client directly.
-   **`cmd/openai`**: A simple example of using the OpenAI client directly.
-   **`cmd/deepseek`**: A simple example of using the DeepSeek client directly.

### `pkg/sqirvy` Directory

The `pkg/sqirvy` directory contains the core API library for interacting with AI models. It provides a unified interface for different AI providers.

-   **`pkg/sqirvy/client.go`**: Defines the `Client` interface and the `NewClient` function for creating clients for different providers.
-   **`pkg/sqirvy/models.go`**: Defines the mapping between model names and their respective providers.
-   **`pkg/sqirvy/anthropic.go`**: Implements the `Client` interface for Anthropic's Claude models.
-   **`pkg/sqirvy/gemini.go`**: Implements the `Client` interface for Google's Gemini models.
-   **`pkg/sqirvy/openai.go`**: Implements the `Client` interface for OpenAI's GPT models.
-   **`pkg/sqirvy/deepseek.go`**: Implements the `Client` interface for DeepSeek's LLM models.
-   **`pkg/sqirvy/meta-llama.go`**: Implements the `Client` interface for Meta's Llama models.

### `pkg/util` Directory

The `pkg/util` directory contains utility functions for file operations, web scraping, and PDF processing.

-   **`pkg/util/files.go`**: Provides functions for reading files, checking for piped input, and validating file paths.
-   **`pkg/util/scraper.go`**: Implements web scraping functionality using the `colly` library.
-   **`pkg/util/pdf.go`**: Provides functions for extracting text from PDF files and summarizing PDF content using AI.

### `web/sqirvy-web` Directory

The `web/sqirvy-web` directory contains the source code for a web interface that allows users to interact with the AI models.

-   **`web/sqirvy-web/main.go`**: Implements the HTTP server and API endpoints for the web interface.
-   **`web/sqirvy-web/static/`**: Contains static files (HTML, CSS, JavaScript) for the web interface.

## Detailed Component Descriptions

### `cmd` Directory Details

#### `cmd/sqirvy-query`

This tool is designed for general-purpose querying of AI models.

**Usage:**

```bash
sqirvy-query [options] files...
```

**Options:**

-   `-h`: Prints the help message.
-   `-m`: Specifies the AI model to use (default: `claude-3-5-sonnet-latest`).

The tool reads input from stdin, if available, and concatenates any specified files. A system prompt can be provided via `system.md` in the current directory.

#### `cmd/sqirvy-review`

This tool automates code reviews using AI language models.

**Usage:**

```bash
sqirvy-review [options] files...
```

**Options:**

-   `-h`: Prints the help message.
-   `-m`: Specifies the AI model to use (default: `gemini-1.5-flash`).

The tool analyzes source code for potential bugs, security vulnerabilities, performance optimizations, style issues, and documentation quality. It uses embedded system prompts (`system.md` and `review.md`) to ensure consistent and thorough reviews.

#### `cmd/sqirvy-scrape`

This tool scrapes content from URLs and sends it to an AI model for analysis.

**Usage:**

```bash
sqirvy-scrape [options] urls...
```

**Options:**

-   `-h`: Prints the help message.
-   `-m`: Specifies the AI model to use (default: `claude-3-5-haiku-latest`).

The tool reads input from stdin, if available, and scrapes content from the specified URLs. It also supports reading files as input. A system prompt can be provided via `system.md` in the current directory.

#### `cmd/anthropic`, `cmd/gemini`, `cmd/openai`, `cmd/deepseek`

These subdirectories contain simple examples of using the respective provider clients directly. They are primarily for testing and demonstration purposes.

### `pkg/sqirvy` Directory Details

#### `pkg/sqirvy/client.go`

This file defines the core `Client` interface and the `NewClient` function.

**`Client` Interface:**

```go
type Client interface {
	QueryText(prompt string, model string, options Options) (string, error)
}
```

The `Client` interface provides a unified way to interact with different AI models. The `QueryText` method sends a text prompt to the specified model and returns the response.

**`NewClient` Function:**

```go
func NewClient(provider Provider) (Client, error)
```

The `NewClient` function creates a new client for the specified provider. It supports the following providers: `Anthropic`, `DeepSeek`, `Gemini`, `OpenAI`, and `MetaLlama`.

#### `pkg/sqirvy/models.go`

This file defines the mapping between model names and their respective providers.

**`ModelToProvider` Map:**

```go
var ModelToProvider = map[string]string{
	// anthropic models
	"claude-3-5-sonnet-latest": "anthropic",
	"claude-3-5-haiku-latest":  "anthropic",
	"claude-3-opus-latest":     "anthropic",
	// google gemini models
	"gemini-2.0-flash-exp": "gemini",
	"gemini-1.5-flash":     "gemini",
	"gemini-1.5-pro":       "gemini",
	// openai models
	"gpt-4o":      "openai",
	"gpt-4o-mini": "openai",
	"gpt-4-turbo": "openai",
	"o1-mini":     "openai",
	// meta-llama models
	"meta-llama/meta-llama-3.1-8b-instruct-turbo": "meta-llama",
	"meta-llama/Llama-3.3-70B-Instruct-Turbo":     "meta-llama",
}
```

The `ModelToProvider` map is used to determine which client implementation should handle requests for a given model.

**`GetProviderName` Function:**

```go
func GetProviderName(model string) (string, error)
```

The `GetProviderName` function returns the provider name for a given model identifier. It returns an error if the model is not recognized.

#### `pkg/sqirvy/anthropic.go`, `pkg/sqirvy/gemini.go`, `pkg/sqirvy/openai.go`, `pkg/sqirvy/deepseek.go`, `pkg/sqirvy/meta-llama.go`

These files implement the `Client` interface for the respective AI providers. They handle authentication, request formatting, and response parsing specific to each provider's API.

### `pkg/util` Directory Details

#### `pkg/util/files.go`

This file provides utility functions for file operations.

**`InputIsFromPipe` Function:**

```go
func InputIsFromPipe() (bool, error)
```

The `InputIsFromPipe` function determines if the program is receiving piped input on stdin.

**`ReadStdin` Function:**

```go
func ReadStdin(maxTotalBytes int64) (string, int64, error)
```

The `ReadStdin` function reads and concatenates the contents of stdin, up to a maximum size.

**`ReadFile` Function:**

```go
func ReadFile(fname string, maxTotalBytes int64) ([]byte, int64, error)
```

The `ReadFile` function reads the contents of a file, up to a maximum size. It also validates the file path to prevent security issues.

**`ReadFiles` Function:**

```go
func ReadFiles(filenames []string, maxTotalBytes int64) (string, int64, error)
```

The `ReadFiles` function reads and concatenates the contents of multiple files, up to a maximum size.

#### `pkg/util/scraper.go`

This file implements web scraping functionality using the `colly` library.

**`ScrapeURL` Function:**

```go
func ScrapeURL(url string) (string, error)
```

The `ScrapeURL` function scrapes the content from a single URL and returns it as a string.

**`ScrapeAll` Function:**

```go
func ScrapeAll(urls []string) (string, error)
```

The `ScrapeAll` function scrapes content from multiple URLs and concatenates the results.

#### `pkg/util/pdf.go`

This file provides functions for extracting text from PDF files and summarizing PDF content using AI.

**`ExtractPDF` Function:**

```go
func ExtractPDF(pdfData []byte) (string, error)
```

The `ExtractPDF` function extracts text content from PDF binary data using the `pdftotext` command-line tool.

**`SummarizePDF` Function:**

```go
func SummarizePDF(pdfData []byte, maxTokens int) (string, error)
```

The `SummarizePDF` function generates an AI-powered summary of PDF content using the sqirvy library.

### `web/sqirvy-web` Directory Details

#### `web/sqirvy-web/main.go`

This file implements the HTTP server and API endpoints for the web interface.

**API Endpoint:**

-   `/api/query`: Accepts a `prompt` parameter and returns responses from Anthropic, OpenAI, and Gemini models.

The server serves static files from the `static` directory.

#### `web/sqirvy-web/static/`

This directory contains static files for the web interface, including `index.html`, `script.js`, and `styles.css`.

## Usage Examples

### Querying an AI Model

```bash
sqirvy-query -m gpt-4-turbo "What is the capital of France?"
```

This command queries the `gpt-4-turbo` model with the prompt "What is the capital of France?".

### Reviewing Code

```bash
sqirvy-review -m gemini-1.5-pro main.go util.go
```

This command reviews the `main.go` and `util.go` files using the `gemini-1.5-pro` model.

### Scraping a Web Page

```bash
sqirvy-scrape -m claude-3-5-sonnet-latest https://example.com
```

This command scrapes the content from `https://example.com` and sends it to the `claude-3-5-sonnet-latest` model.

### Using the Web Interface

1.  Run the web server:

    ```bash
    go run web/sqirvy-web/main.go
    ```
2.  Open a web browser and navigate to `http://localhost:8080`.
3.  Enter a prompt in the text box and click "Query".

## Conclusion

Sqirvy provides a flexible and extensible framework for interacting with various AI language models. Its modular design allows for easy integration of new providers and models. The command-line tools and web interface make it accessible to both developers and end-users.
