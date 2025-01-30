# Code Deep Dive: sqirvy-query, pkg/sqirvy, and pkg/util

This blog post provides a detailed walkthrough of the code structure and functionality within the `cmd/sqirvy-query`, `pkg/sqirvy`, and `pkg/util` directories of the sqirvy-llm project. We'll explore how these components work together to enable command-line querying of various AI language models.

## cmd/sqirvy-query

The `cmd/sqirvy-query` directory contains the main application logic for the `sqirvy-query` command-line tool. This tool allows users to send prompts to different AI models, reading input from files, standard input, and a system prompt file.

### `main.go`

The `main.go` file is the entry point for the `sqirvy-query` application. It handles:

-   **Command-line argument parsing:** It uses `processCommandLine()` to parse command-line arguments, including the model to use and input files.
-   **Model selection:** It determines the AI model to use, defaulting to `claude-3-5-sonnet-latest` if none is specified.
-   **Provider lookup:** It uses `sqirvy.GetProviderName()` to find the provider associated with the selected model.
-   **Client creation:** It creates an AI client using `sqirvy.NewClient()` for the determined provider.
-   **Query execution:** It sends the combined prompt to the AI model using `client.QueryText()`.
-   **Response output:** It prints the AI model's response to standard output.

### `cli.go`

The `cli.go` file handles command-line interface specifics:

-   **Help message:** The `helpMessage()` function displays usage information, including available options and supported models.
-   **Argument processing:** The `processCommandLine()` function:
    -   Parses command-line flags using the `flag` package.
    -   Reads a system prompt from `system.md` if it exists.
    -   Reads input from standard input (if available).
    -   Reads content from specified files.
    -   Enforces a maximum total size limit (`MaxTotalBytes`) for all input.
    -   Returns the combined prompt string and the selected model.

## pkg/sqirvy

The `pkg/sqirvy` directory contains the core logic for interacting with different AI providers. It defines interfaces and implementations for each supported provider.

### `client.go`

The `client.go` file defines the main `Client` interface and the `NewClient` function:

-   **`Provider` type:** Defines an enumeration of supported AI providers (`Anthropic`, `DeepSeek`, `Gemini`, `OpenAI`, `MetaLlama`).
-   **`Options` struct:** Defines a structure for provider-specific options (currently empty).
-   **`Client` interface:** Defines the `QueryText` method for sending text prompts to AI models.
-   **`NewClient` function:** Creates a new client instance for the specified provider using a switch statement.

### `anthropic.go`

The `anthropic.go` file implements the `Client` interface for Anthropic's Claude models:

-   **`AnthropicClient` struct:** Holds the Anthropic API client.
-   **`QueryText` method:**
    -   Initializes the Anthropic client if not already done.
    -   Creates a new message request with the provided prompt.
    -   Sends the request to the Anthropic API.
    -   Returns the response from the AI model.

### `deepseek.go`

The `deepseek.go` file implements the `Client` interface for DeepSeek's LLM models:

-   **`DeepSeekClient` struct:** Holds the DeepSeek API client and API key.
-   **`QueryText` method:**
    -   Initializes the DeepSeek client and API key if not already done.
    -   Constructs the request body with the prompt.
    -   Sends the request to the DeepSeek API.
    -   Returns the response from the AI model.
-   **`QueryJSON` method:**
    -   Similar to `QueryText` but can be used for JSON requests.
-   **`makeRequest` method:**
    -   Marshals the request body to JSON.
    -   Creates an HTTP request with the JSON body.
    -   Sets required headers, including the API key.
    -   Sends the request and reads the response.
    -   Parses the JSON response and returns the content.

### `gemini.go`

The `gemini.go` file implements the `Client` interface for Google's Gemini models:

-   **`GeminiClient` struct:** Holds the Gemini API client and context.
-   **`initClient` method:** Initializes the Gemini client with the API key from the environment.
-   **`QueryText` method:**
    -   Initializes the Gemini client if not already done.
    -   Creates a generative model instance with the specified model name.
    -   Sends the prompt to the Gemini API.
    -   Returns the response from the AI model.

### `meta-llama.go`

The `meta-llama.go` file implements the `Client` interface for Meta's Llama models:

-   **`MetaLlamaClient` struct:** Holds the OpenAI-compatible LLM client.
-   **`QueryText` method:**
    -   Initializes the LLM client if not already done, using the `TOGETHER_API_KEY` and `TOGETHER_API_BASE` environment variables.
    -   Sends the prompt to the LLM.
    -   Returns the response from the AI model.

### `openai.go`

The `openai.go` file implements the `Client` interface for OpenAI's GPT models:

-   **`OpenAIClient` struct:** Holds the OpenAI API key and HTTP client.
-   **`QueryText` method:**
    -   Initializes the HTTP client and API key if not already done.
    -   Constructs the request body with the prompt.
    -   Sends the request to the OpenAI API.
    -   Returns the response from the AI model.
-   **`makeRequest` method:**
    -   Marshals the request body to JSON.
    -   Creates an HTTP request with the JSON body.
    -   Sets required headers, including the API key.
    -   Sends the request and reads the response.
    -   Parses the JSON response and returns the content.

### `models.go`

The `models.go` file defines the mapping between model names and providers:

-   **`ModelToProvider` map:** Maps model names to their respective providers.
-   **`GetProviderName` function:** Returns the provider name for a given model identifier.

## pkg/util

The `pkg/util` directory contains utility functions for file operations, web scraping, and PDF processing.

### `files.go`

The `files.go` file provides utility functions for file and standard input handling:

-   **`InputIsFromPipe` function:** Checks if the program is receiving piped input on stdin.
-   **`ReadStdin` function:** Reads and concatenates the contents of stdin.
-   **`validateFilePath` function:** Checks if the given file path is safe and returns a cleaned version.
-   **`ReadFile` function:** Reads and concatenates the contents of a single file, enforcing a size limit.
-   **`ReadFiles` function:** Reads and concatenates the contents of multiple files, enforcing a size limit.

### `scraper.go`

The `scraper.go` file provides web scraping functionality:

-   **`ScrapeURL` function:** Scrapes the content from a single URL using the `colly` library.
-   **`ScrapeAll` function:** Scrapes content from multiple URLs and concatenates the results.

### `pdf.go`

The `pdf.go` file provides PDF processing functionality:

-   **`ExtractPDF` function:** Extracts text content from PDF binary data using the `pdftotext` command-line tool.
-   **`SummarizePDF` function:** Generates an AI-powered summary of PDF content using the `ExtractPDF` function and the `sqirvy` library.

## Conclusion

This blog post has provided a comprehensive overview of the code in `cmd/sqirvy-query`, `pkg/sqirvy`, and `pkg/util`. These components work together to provide a flexible and powerful command-line tool for interacting with various AI language models. By understanding the structure and functionality of each part, developers can more easily extend and customize the tool for their specific needs.
