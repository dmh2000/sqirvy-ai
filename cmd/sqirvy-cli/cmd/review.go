/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// reviewCmd represents the review command
var reviewCmd = &cobra.Command{
	Use:   "review",
	Short: "Request the LLM to generate a code review .",
	Long: `sqiryv-cli review
It will to review input code and will output the results to stdout.
The prompt is constructed in this order:
    An internal system prompt for code review
    Input from stdin
    Any number of filename or url arguments
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(cmd.Flags().Args())
		fmt.Println(cmd.Flag("model").Value)
		fmt.Println(cmd.Flag("temperature").Value)
		prompt, err := ReadPrompt(reviewPrompt, args)
		if err != nil {
			fmt.Println(fmt.Errorf("error reading prompt: \n%v", err))
			return
		}
		fmt.Println(prompt)
	},
}

func reviewUsage(cmd *cobra.Command) error {
	fmt.Println("Usage: stdin | sqirvy-cli review [flags] [files| urls]")
	return nil
}

func init() {
	rootCmd.AddCommand(reviewCmd)
	reviewCmd.SetUsageFunc(reviewUsage)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// reviewCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// reviewCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
