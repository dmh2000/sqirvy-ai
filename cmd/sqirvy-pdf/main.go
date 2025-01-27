package main

import (
	"fmt"
	"log"

	sqirvy "sqirvy-llm/pkg/sqirvy"
	util "sqirvy-llm/pkg/util"
)

const maxFileSize int64 = 10 * 1024 * 1024 // 10MB

func main() {
	// process command line arguments
	model, stdinData, files, err := processCommandLine()
	if err != nil {
		log.Fatal(err)
	}

	// set default model if not specified
	if model == "" {
		model = "claude-3-5-sonnet-latest"
	}

	// Get the provider for the model
	provider, err := sqirvy.GetProviderName(model)
	if err != nil {
		log.Fatal(err)
	}

	// create API client based on provider
	client, err := sqirvy.NewClient(sqirvy.Provider(provider))
	if err != nil {
		log.Fatal(err)
	}

	// process stdin and files separately and output the response to stdout in markdown format
	context := `You are are an experienced editor and researcher whose job is to review pdf files and generate 
	a consise summary in markdown format while maintaining the general format of the document.`
	context += "\n\n"

	// process stdin data separately
	if stdinData != "" {
		text, err := util.ExtractPDF([]byte(stdinData))
		if err != nil {
			log.Fatal(err)
		}
		context += fmt.Sprintf("```pdf\n%s\n```\n\n", text)
		context += "\n\nSummarize the document and output markdown text.\n"

		// query the model with the context
		response, err := client.QueryText(context, model, sqirvy.Options{})
		if err != nil {
			log.Fatal(err)
		}

		fmt.Print("<markdown>\n")
		fmt.Print(response)
		fmt.Print("</markdown\n")
	}

	// process each file separately
	for _, filename := range files {
		data, _, err := util.ReadFile(filename, maxFileSize)
		if err != nil {
			log.Fatal(err)
		}

		text, err := util.ExtractPDF(data)
		if err != nil {
			log.Fatal(err)
		}
		context += fmt.Sprintf("```pdf\n%s\n```\n\n", text)
		context += "\n\nSummarize the document and output markdown text.\n"

		// query the model with the context
		response, err := client.QueryText(context, model, sqirvy.Options{})
		if err != nil {
			log.Fatal(err)
		}

		fmt.Print("<markdown>\n")
		fmt.Print(response)
		fmt.Print("</markdown>\n")
	}
}
