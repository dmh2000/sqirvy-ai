/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// queryCmd represents the query command
var queryCmd = &cobra.Command{
	Use:   "query",
	Short: "Execute an arbitrary query to the LLM",
	Long: `sqirvy-cli query will send a request to the LLM to execute an arbitrary query.
It will not add internal prompts or context. The prompt to the LLM will consist of 
any input from stdint, and then any filename or url arguments, in the order specified.
`,
	Run: func(cmd *cobra.Command, args []string) {
		response, err := executeQuery(cmd, []string{queryPrompt}, args)
		if err != nil {
			log.Fatal(err)
		}
		// Print response to stdout
		fmt.Print(response)
		fmt.Println()
	},
}

func queryUsage(cmd *cobra.Command) error {
	fmt.Println("Usage: stdin | sqirvy-cli query [flags] [files| urls]")
	return nil
}

func init() {
	rootCmd.AddCommand(queryCmd)
	queryCmd.SetUsageFunc(queryUsage)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// queryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// queryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
