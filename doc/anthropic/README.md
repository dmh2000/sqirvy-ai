# Sqirvy Code Documentation

## Introduction
Sqirvy is a multi-provider AI interaction toolkit that provides command-line tools, libraries, and web interfaces for interacting with various AI models including Anthropic's Claude, OpenAI's GPT models, Google's Gemini, and DeepSeek. The codebase provides unified interfaces for querying these AI services while handling provider-specific implementations transparently.

## Command Line Tools (`./cmd`)

### sqirvy-query
A command-line tool for sending queries to AI models.
- Usage: `sqirvy-query [options] files...`
- Options:
  - `-h`: Display help message
  - `-m`: Specify AI model (default: claude-3-5-sonnet-latest)
- Accepts input from stdin or files
- Concatenates multiple input files in order

### sqirvy-review
Code review tool using AI models.
- Usage: `sqirvy-review [options] files...`
- Options:
  - `-h`: Display help message
  - `-m`: Specify AI model (default: gemini-1.5-flash)
- Designed for code review workflows

### sqirvy-scrape
Web content scraping tool with AI processing capabilities.
- Usage: `sqirvy-scrape [options] urls...`
- Options:
  - `-h`: Display help message
  - `-m`: Specify AI model (default: claude-3-5-sonnet-latest)
- Scrapes content from provided URLs and processes it through specified AI model

## Core Library (`./pkg/sqirvy`)

### Client Interface
The core interface for AI provider interactions:
```go
type Client interface {
    QueryText(prompt string, model string, options Options) (string, error)
}
```

### Supported Providers
- Anthropic (Claude models)
- OpenAI (GPT models)
- Google Gemini
- DeepSeek
- Meta Llama (via Together API)

### Provider-Specific Clients
Each provider has a dedicated client implementation:
- `AnthropicClient`: Anthropic API integration
- `OpenAIClient`: OpenAI API integration
- `GeminiClient`: Google Gemini API integration
- `DeepSeekClient`: DeepSeek API integration
- `MetaLlamaClient`: Meta Llama models via Together API

### Configuration
Each provider requires specific environment variables:
- Anthropic: `ANTHROPIC_API_KEY`
- OpenAI: `OPENAI_API_KEY`, optional `OPENAI_API_BASE`
- Gemini: `GEMINI_API_KEY`
- DeepSeek: `DEEPSEEK_API_KEY`, `DEEPSEEK_API_BASE`
- Together: `TOGETHER_API_KEY`

## Utility Package (`./pkg/util`)

### File Operations
- `InputIsFromPipe()`: Checks if input is coming from a pipe
- `ReadStdin()`: Reads input from stdin with size limits
- `ReadFile()`: Safely reads file content with path validation
- `ReadFiles()`: Processes multiple input files

### PDF Processing
- `ExtractPDF()`: Extracts text content from PDF data
- `SummarizePDF()`: Creates summarized version of PDF content

### Web Scraping
- `ScrapeURL()`: Extracts content from a single URL
- `ScrapeAll()`: Processes multiple URLs and combines their content

## Web Interface (`./web/sqirvy-web`)

### Overview
Web-based interface for interacting with multiple AI providers simultaneously.

### API Response Structure
```go
type QueryResponse struct {
    Anthropic ProviderResponse
    OpenAI    ProviderResponse
    Gemini    ProviderResponse
}

type ProviderResponse struct {
    Result string
    Error  string
}
```

### Features
- Concurrent queries to multiple AI providers
- JSON-based API responses
- Web interface for easy interaction
- Support for comparing responses across providers

## Usage Examples

### Basic Query
```go
client, err := sqirvy.NewClient(sqirvy.Anthropic)
if err != nil {
    log.Fatal(err)
}

response, err := client.QueryText("Your prompt here", "claude-3-5-sonnet-latest", sqirvy.Options{})
```

### File Processing
```go
content, size, err := util.ReadFiles([]string{"file1.txt", "file2.txt"}, maxBytes)
```

### Web Scraping
```go
content, err := util.ScrapeURL("https://example.com")
```

## Error Handling
All components return explicit errors that should be checked. Common errors include:
- API authentication failures
- Rate limiting
- Invalid input
- File access issues
- Network connectivity problems

## Best Practices
1. Always check returned errors
2. Use appropriate model for the task
3. Respect API rate limits
4. Validate input before processing
5. Handle provider-specific limitations
6. Use environment variables for API keys
7. Implement proper error handling in production

## Security Considerations
1. Never commit API keys
2. Validate file paths
3. Sanitize user input
4. Use HTTPS for API calls
5. Implement proper access controls in web interface
6. Handle sensitive data appropriately
