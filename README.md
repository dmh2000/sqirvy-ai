# Sqirvy AI - Code Review and Documentation Tool

Sqirvy AI is a command-line tool that leverages Anthropic's Claude AI model to automatically review and document code files. It provides two main functionalities: code review and documentation generation.

## Features

- **Code Review** (`sqirvy-review`): Analyzes code for bugs, security issues, style problems, and suggests improvements
- **Documentation** (`sqirvy-doc`): Generates README-style documentation for code files

## Installation

```bash
pip install -e .
```

## Usage

### Code Review

```bash
sqirvy-review <filename>
```
This will analyze the given file and output a markdown-formatted review including:
- Bug detection
- Security analysis 
- Style recommendations
- Idiomatic improvements

The tool will also display:
- API call timing
- Input/output token usage
- Cache status

### Documentation Generation 

```bash
sqirvy-doc <filename>
```
This generates README-style documentation in markdown format, along with performance metrics.

## Configuration

Configuration is managed through `.sqirvy_ai.config.yml`:

```yaml
anthropic:
  TOKEN_LIMIT: 2048  # Maximum tokens per request
  MODEL_NAME: claude-3-5-sonnet-20241022  # Anthropic model to use
```

## Project Structure

### Core Components

- `src/common/anthropic.py`: Handles Claude AI integration:
  - API interaction and response processing
  - Message formatting for reviews/docs
  - Token management and caching
  - Performance metrics

- `src/common/fetch.py`: File processing utility:
  - Code file reading
  - Content preprocessing
  - Error handling

### CLI Tools

- `src/sqirvy_review/main.py`: Code review entry point
- `src/sqirvy_doc/main.py`: Documentation generator entry point

## Dependencies

- `anthropic`: Official Python client for Claude AI
- `pyyaml`: YAML configuration parsing
