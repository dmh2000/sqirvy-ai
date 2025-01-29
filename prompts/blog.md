write a blog post, in markdown format, that :
- describes how to create a sqirvy Client
- describe the three different approaches to interacting with LLM's to make a simple request
  - in openai.go, it uses a standard HTTP REST request as described by the openai REST API spec
  - in anthropic.go, it uses a native library supported directly by Anthropic.com
  - in meta-llama.go, it uses a LangChain compatible library, tmc/langchaingo
- write it as it applies to Go programmers
- store it in file doc/golang-ai.md
