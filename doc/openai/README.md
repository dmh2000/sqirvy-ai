# Documentation for Sqirvy Codebase

## Introduction

Sqirvy is an AI-powered tool designed to assist with web scraping, data analysis, and code review. The codebase is organized into several key directories:

- `./cmd`: Contains command-line interface (CLI) tools for interacting with various AI models and performing tasks such as querying, scraping, and reviewing code.
- `./pkg/sqirvy`: Houses the core library implementations, including integrations with different AI providers (Anthropic, OpenAI, Google Gemini, DeepSeek, Meta Llama), model management, and client interfaces.
- `./pkg/util`: Provides utility functions for file operations, web scraping, PDF processing, and handling input/output operations.
- `./web/sqirvy-web`: Implements a web-based interface for interacting with the Sqirvy AI models, allowing users to perform queries through a web browser.

This documentation aims to provide an overview of the code structure, key functionalities, and guidance on how to use and extend the codebase. It is intended for experienced software engineers looking to understand and leverage Sqirvy's capabilities in their projects.

## Directory Structure

### ./cmd

The `./cmd` directory contains various CLI tools that serve different purposes within the Sqirvy ecosystem. Each subdirectory under `./cmd` corresponds to a specific tool or functionality:

- `./cmd/deepseek`: Tool for interacting with DeepSeek's AI models.
- `./cmd/anthropic`: CLI for Anthropic's Claude models.
- `./cmd/gemini`: Interface for Google's Gemini AI models.
- `./cmd/openai`: CLI for OpenAI's GPT models.
- `./cmd/sqirvy-query`: Tool for querying AI models with combined input from files and stdin.
- `./cmd/sqirvy-review`: Automated code review tool leveraging AI models.
- `./cmd/sqirvy-scrape`: Command-line tool for web scraping and AI analysis.

Each CLI tool follows a similar structure with `main.go` handling the entry point and command-line processing.

### ./pkg/sqirvy

The `./pkg/sqirvy` directory contains the core library responsible for managing AI model integrations and providing a unified interface for querying different AI providers. Key components include:

- **Clients**: Implementations for each AI provider (`AnthropicClient`, `OpenAIClient`, `GeminiClient`, `DeepSeekClient`, `MetaLlamaClient`) adhering to the `Client` interface.
- **Model Management**: Mapping of model names to providers (`ModelToProvider`) and utility functions like `GetProviderName`.
- **Options**: Configuration structures for provider-specific options.

This library abstracts the complexities of interacting with various AI APIs, allowing CLI tools and other components to utilize AI functionalities seamlessly.

### ./pkg/util

The `./pkg/util` directory provides utility functions that support different aspects of the codebase, including:

- **File Operations**: Reading and validating file content, handling input from stdin and files, size limit enforcement.
- **Web Scraping**: Functions for scraping content from URLs using the `colly` library.
- **PDF Processing**: Extracting text from PDFs using external tools like `pdftotext`.

These utilities are designed to be reusable across different parts of the codebase, ensuring consistency and reliability in common operations.

### ./web/sqirvy-web

The `./web/sqirvy-web` directory implements the web-based interface for Sqirvy. Key features include:

- **Static Files**: Serving static assets such as HTML, CSS, and JavaScript from the `static` directory.
- **API Endpoints**: Handling API requests to query different AI providers and returning their responses in a structured JSON format.

This web interface allows users to interact with Sqirvy's AI functionalities through a user-friendly web application, enabling real-time queries and responses.

## Getting Started

To get started with Sqirvy, follow these steps:

1. **Clone the Repository**
   ```
   git clone https://github.com/yourusername/sqirvy.git
   cd sqirvy
   ```

2. **Install Dependencies**
   Ensure you have the necessary dependencies installed, such as Go and any required tools like `pdftotext`.

3. **Set Up Environment Variables**
   Configure the required API keys for the AI providers:
   ```
   export OPENAI_API_KEY=your_openai_api_key
   export ANTHROPIC_API_KEY=your_anthropic_api_key
   export GEMINI_API_KEY=your_gemini_api_key
   export DEEPSEEK_API_KEY=your_deepseek_api_key
   export TOGETHER_API_KEY=your_meta_llama_api_key
   ```

4. **Build and Run CLI Tools**
   Navigate to the desired CLI tool directory and build the executable:
   ```
   cd cmd/sqirvy-query
   go build -o sqirvy-query
   ./sqirvy-query -m openai "Your query here"
   ```

5. **Run the Web Interface**
   Navigate to the web interface directory, build, and start the server:
   ```
   cd ../web/sqirvy-web
   go build -o sqirvy-web
   ./sqirvy-web
   ```
   Access the web application at `http://localhost:8080`.

## Usage Examples

### Querying AI Models via CLI

Use the `sqirvy-query` tool to send prompts to AI models:
```
./sqirvy-query -m gpt-4-turbo "Explain the significance of the Turing test."
```

### Automated Code Review

Run the `sqirvy-review` tool to perform a code review:
```
./sqirvy-review -m gemini-1.5-flash ./cmd ./pkg
```

### Web Scraping and AI Analysis

Use the `sqirvy-scrape` tool to scrape web content and analyze it with an AI model:
```
./sqirvy-scrape -m openai https://example.com
```

## Extending the Codebase

To add support for a new AI provider:

1. **Implement the Client Interface**
   Create a new client struct in `./pkg/sqirvy` that implements the `Client` interface.

2. **Update Model Mappings**
   Add the new model and provider to the `ModelToProvider` map in `models.go`.

3. **Create a CLI Tool (Optional)**
   Add a new CLI tool under `./cmd` if you want a dedicated command-line interface for the new provider.

4. **Update Documentation**
   Document the new provider and its usage in the `doc/openai/README.md`.

## Contribution

Contributions are welcome! Please open issues and submit pull requests for enhancements and bug fixes.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
