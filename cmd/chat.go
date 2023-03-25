/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"

	"github.com/sashabaranov/go-openai"
	"github.com/spf13/cobra"
)

// chatCmd represents the chat command
var chatCmd = &cobra.Command{
	Use:   "chat",
	Short: "Chat-based completions",
	Long: `Chat-based completions, you can use it in Q&A or Conversation mode.
It's different from text completion:
- text completion, simple but powerful interface to any of OpenAI models,
  desc: https://platform.openai.com/docs/guides/completion/introduction
  model: https://platform.openai.com/docs/models
- chat completion, conversation streaming API, based on gpt-3.5-turbo or gpt-4,
  desc: https://platform.openai.com/docs/guides/chat
  
Because gpt-3.5-turbo performs at a similar capability to text-davinci-003,
but at 10% the price per token, we recommend gpt-3.5-turbo for most use cases.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		model, _ := cmd.Flags().GetString("model")
		if model != openai.GPT3Dot5Turbo && model != openai.GPT4 {
			return errors.New("chat completions: only gpt-3.5-turbo and gpt-4 supported")
		}

		interact, _ := cmd.Flags().GetBool("interactive")
		if !interact && len(args) == 0 {
			return errors.New("missing prompt")
		}

		return handleChatGPT3Prompt(model, interact, args)
	},
}

func init() {
	rootCmd.AddCommand(chatCmd)
}
