/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
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
	"github.com/sashabaranov/go-openai"
	"github.com/spf13/cobra"
)

var chatCmdDescShort = `Chat-based completion`
var chatCmdDescLong = `Chat-based completion, you can use it in Q&A or Conversation mode.

It's different from text completion:

- text completion, simple but powerful interface to any of OpenAI models,
  desc: https://platform.openai.com/docs/guides/completion/introduction
  model: https://platform.openai.com/docs/models
- chat completion, conversation streaming API, based on gpt-3.5-turbo or gpt-4,
  desc: https://platform.openai.com/docs/guides/chat
  
Because gpt-3.5-turbo performs at a similar capability to text-davinci-003,
but at 10% the price per token, we recommend gpt-3.5-turbo for most use cases.`

// chatCmd represents the chat command
var chatCmd = &cobra.Command{
	Use:   "chat",
	Short: chatCmdDescShort,
	Long:  chatCmdDescLong,
	RunE: func(cmd *cobra.Command, args []string) error {
		model, _ := cmd.Flags().GetString("model")
		if model != openai.GPT3Dot5Turbo && model != openai.GPT4 {
			return errors.New("chat completions: only gpt-3.5-turbo and gpt-4 supported")
		}

		interact, _ := cmd.Flags().GetBool("interactive")
		if !interact && len(args) == 0 {
			return errors.New("missing prompt")
		}

		return handleChatPrompt(model, interact, args)
	},
}

func init() {
	rootCmd.AddCommand(chatCmd)
}

func handleChatPrompt(model string, interfact bool, args []string) error {
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
