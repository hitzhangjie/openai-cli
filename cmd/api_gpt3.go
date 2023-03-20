package cmd

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/fatih/color"
	openai "github.com/sashabaranov/go-openai"
)

func handleChatGPT3Prompt(model string, interfact bool, args []string) error {
	token, err := getOpenAIToken()
	if err != nil {
		return err
	}

	ctx := context.Background()
	client := openai.NewClient(token)

	if !interfact {
		req := buildChatCompletionRequest(model, args, interfact)
		resp, err := client.CreateChatCompletion(ctx, req)
		if err != nil {
			return fmt.Errorf("ChatCompletion error: %v", err)
		}
		fmt.Println(resp.Choices[0].Message.Content)
		return nil
	}

	messages := make([]openai.ChatCompletionMessage, 0)
	reader := bufio.NewReader(os.Stdin)

	for {
		color.Set(color.FgWhite)
		fmt.Printf("\nprompt $ ")

		var text string
		for {
			v, err := reader.ReadString('\n')
			if err == io.EOF {
				fmt.Println()
				break
			}
			text += v
		}
		messages = append(messages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleUser,
			Content: text,
		})

		color.Set(color.FgGreen)
		fmt.Printf("\ncomplete $ ")

		req := buildChatCompletionRequestWithMessages(model, messages)
		stream, err := client.CreateChatCompletionStream(ctx, req)
		if err != nil {
			return fmt.Errorf("ChatCompletionStream error: %v", err)
		}
		for {
			resp, err := stream.Recv()
			if errors.Is(err, io.EOF) {
				break
			}
			if err != nil {
				return fmt.Errorf("Stream error: %v", err)
			}
			fmt.Printf(resp.Choices[0].Delta.Content)
		}
		fmt.Println()
	}
}

func buildChatCompletionRequest(model string, prompt []string, interfact bool) openai.ChatCompletionRequest {
	return openai.ChatCompletionRequest{
		Model: model,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: strings.Join(prompt, " "),
			},
		},
		Stream: interfact,
	}
}

func buildChatCompletionRequestWithMessages(model string, messages []openai.ChatCompletionMessage) openai.ChatCompletionRequest {
	return openai.ChatCompletionRequest{
		Model:    model,
		Messages: messages,
	}
}
