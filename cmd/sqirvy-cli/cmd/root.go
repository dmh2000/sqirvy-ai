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

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "sqirvy-cli",
	Short: "A command line tool to interact with Large Language Models",
	Long: `Sqirvy-cli is a command line tool to interact with Large Language Models (LLMs).
It provides a simple interface to send prompts to the LLM and receive responses.
Sqirvy-cli commands receive prompt input from stdin and any filename args. Output is sent to stdout.
The output is determined by the command and the input prompt.
The "query" command is used to send an arbitrary query to the LLM.
The "plan" command is used to send a prompt to the LLM and receive a plan in response.
The "code" command is used to send a prompt to the LLM and receive source code in response.
The "review" command is used to send a prompt to the LLM and receive a code review in response.
Sqirvy-cli is designed to support terminal command pipelines. 
It can receive input from stdin and/or command line arguments, and writes results to stdout. 
	`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
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
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config/sqirvy-cli.yaml)")
	rootCmd.PersistentFlags().StringP("model", "m", "gpt-4-turbo", "LLM model to use")
	rootCmd.PersistentFlags().IntP("temperature", "t", 50, "LLM temperature to use (0..100)")
}

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
		viper.AddConfigPath(home + "/.config/")
		viper.SetConfigType("yaml")
		viper.SetConfigName("sqirvy-cli")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
