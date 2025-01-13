# sqirvy-query

A command line tool for querying AI language models from multiple providers.

## Description

sqirvy-query is a command line interface that allows you to send prompts to various AI language models and receive their responses. It can read prompts from standard input (stdin) and/or files, concatenate them, and send the combined prompt to the specified AI model.

## Usage

```bash
sqirvy-query [options] files...
```

### Options

- `-h`: Print help message
- `-m`: Specify the AI model to use (default: claude-3-5-sonnet-latest)

### Environment Variables Required

Depending on which model you select, you'll need to set the appropriate API key:

- For Anthropic models: `ANTHROPIC_API_KEY`
- For OpenAI models: `OPENAI_API_KEY`
- For Gemini models: `GEMINI_API_KEY`

### Examples

1. Send a prompt from stdin:
```bash
echo "What is the capital of France?" | sqirvy-query
```

2. Send a prompt from a file:
```bash
sqirvy-query prompt.txt
```

3. Combine stdin and files:
```bash
echo "Please answer this question:" | sqirvy-query prompt.txt
```

4. Use a specific model:
```bash
sqirvy-query -m gpt-4-turbo-preview prompt.txt
```

### Supported Models

The tool supports various models from different providers:

Anthropic:
- claude-3-opus-20240229
- claude-3-sonnet-20240229
- claude-2.1
- claude-instant-1.2

OpenAI:
- gpt-4-turbo-preview
- gpt-4
- gpt-3.5-turbo

Gemini:
- gemini-pro
- gemini-pro-vision

## Exit Status

- 0: Success
- 1: Error (invalid arguments, API errors, etc.)
