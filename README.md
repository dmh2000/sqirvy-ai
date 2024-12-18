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

### Documentation Generation

```bash
sqirvy-doc <filename>
```

## Project Structure

### Main Components

- `src/sqirvy_review/main.py`: Entry point for the code review functionality. Processes command line arguments and displays results including token usage and API call timing.

- `src/sqirvy_doc/main.py`: Entry point for the documentation generation functionality. Handles command line input and outputs generated documentation with performance metrics.

### Common Utilities

- `src/common/anthropic.py`: Core functionality for interacting with Claude AI:
  - Message formatting for reviews and documentation
  - API call handling with token limits
  - Response processing and error handling
  - Performance metrics tracking
- `src/common/fetch.py`: File handling utility that:
  - Reads source code files
  - Preprocesses content for AI analysis
  - Handles file not found errors

## Dependencies

- anthropic: Python client for the Anthropic Claude API

## Configuration

The tool uses Claude 3 Sonnet (version 20241022) with:

- Maximum token limit: 2048
- Prompt caching enabled
