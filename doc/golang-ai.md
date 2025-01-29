# Three Approaches to LLM Integration in Go

When integrating Large Language Models (LLMs) into your Go applications, there are several approaches you can take. This post explores three different methods using the Sqirvy client library as examples.

## Creating a Sqirvy Client

The Sqirvy library provides a unified interface for interacting with various LLM providers. To create a client:

```go
import "github.com/yourusername/sqirvy/pkg/sqirvy"

// Create a client for your chosen provider
client, err := sqirvy.NewClient(sqirvy.OpenAI) // or Anthropic, MetaLlama, etc.
if err != nil {
    log.Fatal(err)
}

// Make a simple query
response, err := client.QueryText("Tell me a joke", "gpt-4-turbo-preview", sqirvy.Options{})
```

## Three Integration Approaches

### 1. Direct REST API Integration (OpenAI Example)

The most straightforward approach is to directly use the provider's REST API. This gives you complete control but requires more boilerplate code. Here's how OpenAI integration works in Sqirvy:

```go
// OpenAI client using standard HTTP requests
type OpenAIClient struct {
    apiKey string
    client *http.Client
}

// Make a request
jsonBody := openAIRequest{
    Model: model,
    Messages: []openAIMessage{
        {Role: "user", Content: prompt},
    },
}
// Send HTTP POST request to OpenAI endpoint
```

Advantages:
- No external dependencies
- Full control over request/response handling
- Easy to debug and modify
- Direct mapping to API documentation

Disadvantages:
- More boilerplate code
- Need to handle HTTP details manually
- Must implement retry logic yourself

### 2. Official SDK Integration (Anthropic Example)

Using an official SDK provides a more polished experience with built-in types and error handling:

```go
// Anthropic client using official SDK
client := anthropic.NewClient()

message, err := client.Messages.New(context.TODO(), anthropic.MessageNewParams{
    Model: anthropic.F(model),
    Messages: anthropic.F([]anthropic.MessageParam{
        anthropic.NewUserMessage(anthropic.NewTextBlock(prompt)),
    }),
})
```

Advantages:
- Type safety
- Built-in error handling
- Official support
- Automatic retries and best practices

Disadvantages:
- Dependency on external package
- Less flexibility for customization
- May lag behind API updates

### 3. LangChain Integration (Meta Llama Example)

Using LangChain provides a unified interface across multiple providers:

```go
// Meta Llama client using LangChain
llm, err := openai.New(
    openai.WithBaseURL(baseURL),
    openai.WithToken(apiKey),
    openai.WithModel(model),
)

completion, err := llms.GenerateFromSinglePrompt(context.Background(), llm, prompt)
```

Advantages:
- Unified interface across providers
- Rich ecosystem of tools and chains
- Easy to switch between providers
- Community support

Disadvantages:
- Additional abstraction layer
- May not expose provider-specific features
- Larger dependency footprint

## Choosing an Approach

Consider these factors when choosing an approach:

1. **Direct REST API** if you:
   - Need complete control
   - Want minimal dependencies
   - Are only using one provider
   - Need to match API docs exactly

2. **Official SDK** if you:
   - Want type safety and official support
   - Need built-in best practices
   - Are primarily using one provider
   - Value ease of use over flexibility

3. **LangChain** if you:
   - Need to support multiple providers
   - Want access to higher-level abstractions
   - Plan to build complex chains
   - Value provider interchangeability

## Conclusion

Each approach has its merits, and the choice depends on your specific needs. Sqirvy demonstrates all three approaches, allowing you to see how they work in practice and choose the one that best fits your use case.

Remember to always handle your API keys securely and respect rate limits regardless of which approach you choose.
