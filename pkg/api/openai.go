package ai

import (
	"context"
	"fmt"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/ai/azopenai"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
)

type OpenAIClient struct {
	client *azopenai.Client
	ctx    context.Context
}

func (c *OpenAIClient) initClient() error {
	if c.client != nil {
		return nil
	}

	var err error
	c.ctx = context.Background()

	key := os.Getenv("AZURE_OPENAI_KEY")
	endpoint := os.Getenv("AZURE_OPENAI_ENDPOINT")

	keyCredential := azcore.NewKeyCredential(key)
	c.client, err = azopenai.NewClientForOpenAI(endpoint, keyCredential, nil)
	if err != nil {
		return fmt.Errorf("failed to create client: %v", err)
	}
	return nil
}

func (c *OpenAIClient) QueryText(prompt string) (string, error) {
	if err := c.initClient(); err != nil {
		return "", err
	}

	deploymentID := os.Getenv("AZURE_OPENAI_DEPLOYMENT_ID")

	resp, err := c.client.GetChatCompletions(
		c.ctx,
		azopenai.ChatCompletionsOptions{
			Messages: []azopenai.ChatMessage{
				{
					Role:    azopenai.ChatRoleUser,
					Content: prompt,
				},
			},
			MaxTokens: pointer(int32(1024)),
		},
		deploymentID,
		nil,
	)
	if err != nil {
		return "", fmt.Errorf("failed to generate completion: %v", err)
	}

	if len(resp.Choices) == 0 {
		return "", fmt.Errorf("no content in response")
	}

	return resp.Choices[0].Message.Content, nil
}

func (c *OpenAIClient) QueryJSON(prompt string) (string, error) {
	if err := c.initClient(); err != nil {
		return "", err
	}

	deploymentID := os.Getenv("AZURE_OPENAI_DEPLOYMENT_ID")

	resp, err := c.client.GetChatCompletions(
		c.ctx,
		azopenai.ChatCompletionsOptions{
			Messages: []azopenai.ChatMessage{
				{
					Role:    azopenai.ChatRoleUser,
					Content: prompt,
				},
			},
			MaxTokens: pointer(int32(1024)),
			ResponseFormat: &azopenai.ChatCompletionsResponseFormat{
				Type: azopenai.ChatCompletionsResponseFormatTypeJSONObject,
			},
		},
		deploymentID,
		nil,
	)
	if err != nil {
		return "", fmt.Errorf("failed to generate completion: %v", err)
	}

	if len(resp.Choices) == 0 {
		return "", fmt.Errorf("no content in response")
	}

	return resp.Choices[0].Message.Content, nil
}

// Helper function to create pointer to int32
func pointer(i int32) *int32 {
	return &i
}
