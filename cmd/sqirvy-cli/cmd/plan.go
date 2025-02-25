/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// planCmd represents the plan command
var planCmd = &cobra.Command{
	Use:   "plan",
	Short: "Request the LLM to generate a plan.",
	Long: `sqiryv-cli plan:
It will ask the LLM to generate a plan based on the given prompt. 
It will send a request to the LLM and output the results to stdout.
Typical usage would be to generate a plan for an application and send it 
to [sqirvy-cli code] to generate the actual code. 
The prompt is constructed in this order:
	An internal system prompt for general planning 
	Input from stdin
	Any number of filename or url arguments	`,
	Run: func(cmd *cobra.Command, args []string) {
		response, err := executeQuery(cmd, planPrompt, args)
		if err != nil {
			log.Fatal(err)
		}
		// Print response to stdout
		fmt.Print(response)
		fmt.Println()
	},
}

func planUsage(cmd *cobra.Command) error {
	fmt.Println("Usage: stdin | sqirvy-cli plan [flags] [files| urls]")
	return nil
}

func init() {
	rootCmd.AddCommand(planCmd)
	planCmd.SetUsageFunc(planUsage)
}
