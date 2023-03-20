package cmd

import (
	"context"
	"fmt"
	"strings"

	openai "github.com/sashabaranov/go-openai"
)

func handleGPT3Prompt(model string, args []string) error {
	token, err := getOpenAIToken()
	if err != nil {
		return err
	}

	client := openai.NewClient(token)

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: model,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: strings.Join(args, " "),
				},
			},
		},
	)
	if err != nil {
		return fmt.Errorf("ChatCompletion error: %v", err)
	}
	fmt.Println(resp.Choices[0].Message.Content)
	return nil
}
