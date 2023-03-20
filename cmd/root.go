/*
j
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/sashabaranov/go-openai"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "openai-cli [flags] <prompt>",
	Short: "openai-cli is an easy-to-use OpenAI client",
	Long: `openai-cli is an easy-to-use OpenAI client,
which supports GPT-3, GPT-4, ChatGPT, DALL.E 2, Whisper APIs.

Use it in interactive mode or in simple Q&A mode ... Enjoy!`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	RunE: func(cmd *cobra.Command, args []string) error {
		interact, _ := cmd.Flags().GetBool("interactive")
		if !interact && len(args) == 0 {
			return errors.New("missing prompt")
		}
		api, _ := cmd.Flags().GetString("api")
		model, _ := cmd.Flags().GetString("model")

		switch api {
		case APIChatGPT:
			return handleChatGPT3Prompt(model, interact, args)
		default:
			return errors.New("not supported")
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.openai-cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.Flags().BoolP("interactive", "i", false, "The mode to use")
	rootCmd.Flags().String("api", APIChatGPT, "The API to use")
	rootCmd.Flags().String("model", openai.GPT3Dot5Turbo, "The model to use")

	// GPT3 flags

	// GPT4 flags

	// ChatGPT flags

	// Dalle.2 flags

	// Whisper flags
}

const (
	APIGPT3    = "gpt-3"
	APIGPT4    = "gpt-4"
	APIChatGPT = "chatgpt"
	APIDalle   = "dalle"
	APIWhisper = "whisper"

	OpenAIToken = "OPENAI_API_TOKEN"
)

func getOpenAIToken() (string, error) {
	v := os.Getenv(OpenAIToken)
	if len(v) == 0 {
		return "", fmt.Errorf("env %s unset", "OPENAI_API_TOKEN")
	}
	return v, nil
}
