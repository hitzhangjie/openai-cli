/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var completeCmdDescShort = `Text completion`
var completeCmdDescLong = `Text completion, it's prompt-based completion built upon models:

  - text-davinci-003
  - text-curie-003
  - text-babbage-001
  - text-ada-001
  - etc.

Text completion API works like single-turn prompt&completion, rather 
than chat-based completion which supports multi-turn conversation.

But if we provide all previous prompt&completion as next prompt,
the text completion API will works like multi-turn mode similarly.

Actually, chat-based API do it similarly.`

// completeCmd represents the complete command
var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: completeCmdDescShort,
	Long:  completeCmdDescLong,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("complete called")
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// completeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// completeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
