/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	_ "embed"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// codeCmd represents the code command
var codeCmd = &cobra.Command{
	Use:   "code",
	Short: "Request the LLM to generate",
	Long: `sqiryv-cli code will send a request to generate code 
and will output the results to stdout.
The prompt is constructed in this order:
	An internal system prompt for code generation
	Input from stdin
	Any number of filename or url arguments	
	`,
	Run: func(cmd *cobra.Command, args []string) {
		response, err := executeQuery(cmd, codePrompt, args)
		if err != nil {
			log.Fatal(err)
		}
		// Print response to stdout
		fmt.Print(response)
		fmt.Println()
	},
}

func codeUsage(cmd *cobra.Command) error {
	fmt.Println("Usage: stdin | sqirvy-cli code [flags] [files| urls]")
	return nil
}

func init() {
	rootCmd.AddCommand(codeCmd)
	codeCmd.SetUsageFunc(codeUsage)
}
