/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	_ "embed"
	"fmt"
	"sort"

	sqirvy "sqirvy-ai/pkg/sqirvy"

	"github.com/spf13/cobra"
)

// codeCmd represents the code command
var modelsCmd = &cobra.Command{
	Use:   "models",
	Short: "list the supported models and providers",
	Long:  `sqirvy-cli models will list the supported models and providers`,
	Run: func(cmd *cobra.Command, args []string) {
		var models []string
		var length int
		for _, v := range sqirvy.ModelToProvider {
			length = max(length, len(v))
		}
		fmt.Println("Supported Providers and Models:")
		for k, v := range sqirvy.ModelToProvider {
			models = append(models, fmt.Sprintf("   %-*s: %s\n", length, v, k))
		}
		sort.Strings(models)
		for _, m := range models {
			fmt.Print(m)
		}
		fmt.Println()
	},
}

func modelsUsage(cmd *cobra.Command) error {
	fmt.Println("Usage: sqirvy-cli models")
	return nil
}

func init() {
	rootCmd.AddCommand(modelsCmd)
	codeCmd.SetUsageFunc(modelsUsage)
}
