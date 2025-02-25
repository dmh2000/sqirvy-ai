/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var defaultPrompt string = "Hello"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "sqirvy-cli [command] [flags] [files| urls]",
	Short: "A command line tool to interact with Large Language Models",
	Long: `Sqirvy-cli is a command line tool to interact with Large Language Models (LLMs).
   - It provides a simple interface to send prompts to the LLM and receive responses.
   - Sqirvy-cli commands receive prompt input from stdin, filenames and URL args. Output is sent to stdout.
   - The output is determined by the command and the input prompt.
   - The "query" command is used to send an arbitrary query to the LLM.
   - The "plan" command is used to send a prompt to the LLM and receive a plan in response.
   - The "code" command is used to send a prompt to the LLM and receive source code in response.
   - The "review" command is used to send a prompt to the LLM and receive a code review in response.
   - Sqirvy-cli is designed to support terminal command pipelines. 
	`,

	Run: func(cmd *cobra.Command, args []string) {
		// if no command is specified, use 'query'
		cmd.SetArgs(append([]string{"query"}, args...))
		cmd.Execute()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	fmt.Fprintln(os.Stderr, "Using model :", viper.GetString("model"))

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// flags
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config/sqirvy-cli/config.yaml)")
	rootCmd.PersistentFlags().StringVar(&defaultPrompt, "default-prompt", "Hello", "default prompt to use")
	viper.BindPFlag("default-prompt", rootCmd.PersistentFlags().Lookup("default-prompt"))
	rootCmd.PersistentFlags().StringP("model", "m", "gpt-4-turbo", "LLM model to use")
	viper.BindPFlag("model", rootCmd.PersistentFlags().Lookup("model"))
	rootCmd.PersistentFlags().IntP("temperature", "t", 50, "LLM temperature to use (0..100)")
}

// print config filename only once
var configPrinted bool

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".config/sqirvy-cli" (without extension).
		viper.AddConfigPath(home + "/.config/sqirvy-cli")
		viper.SetConfigType("yaml")
		viper.SetConfigName("config")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		if !configPrinted {
			configPrinted = true
			fmt.Fprintln(os.Stderr, "Config file :", viper.ConfigFileUsed())
		}
	}
}
