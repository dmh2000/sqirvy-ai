package ai

import (
	"context"
	"fmt"

	"github.com/anthropics/anthropic-sdk-go"
)

const MAX_TOKENS = 2048

type AnthropicClient struct {
	client *anthropic.Client
}

func (c *AnthropicClient) QueryText(prompt string, model string, options Options) (string, error) {
	if c.client == nil {
		c.client = anthropic.NewClient()
	}

	answer := ""
	message, err := c.client.Messages.New(context.TODO(), anthropic.MessageNewParams{
		Model:     anthropic.F(model),
		MaxTokens: anthropic.F(int64(1024)),
		Messages: anthropic.F([]anthropic.MessageParam{
			anthropic.NewUserMessage(anthropic.NewTextBlock(prompt)),
		}),
	})
	if err != nil {
		return "", err
	}
	if len(message.Content) == 0 {
		return "", fmt.Errorf("no content in response")
	}
	for _, content := range message.Content {
		answer += fmt.Sprintf("%v", content.Text)
	}
	return answer, nil
}

func (c *AnthropicClient) QueryJSON(prompt string, model string, options Options) (string, error) {
	if c.client == nil {
		c.client = anthropic.NewClient()
	}

	message, err := c.client.Messages.New(context.TODO(), anthropic.MessageNewParams{
		Model:     anthropic.F(model),
		MaxTokens: anthropic.F(int64(MAX_TOKENS)),
		Messages: anthropic.F([]anthropic.MessageParam{
			anthropic.NewUserMessage(anthropic.NewTextBlock(prompt)),
		}),
	})
	if err != nil {
		return "", err
	}
	if len(message.Content) == 0 {
		return "", fmt.Errorf("no content in response")
	}
	if len(message.Content[0].Text) == 0 {
		return "", fmt.Errorf("no JSON content in response")
	}
	if message.Content[0].JSON.Text.IsInvalid() {
		return "", fmt.Errorf("invalid JSON content in response")
	}
	r := message.Content[0].JSON.Text.Raw()
	return r, nil
}
