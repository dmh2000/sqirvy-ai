create a Go command line application in directory cmd/sqirvy-scrape with the following features:

# Command Line Processing

- place the command line processing in a separate file named cli.go and function named "processCommandLine"
- return the prompt, model name if any and any errors
- includes these features:
  - has a command line argment -h that prints a help message and returns
  - has a command line argument -model that specifies the model to use.
  - any other arguments are considered to be web urls to scrape
  - construct a string variable named "prompt" that will concatenate the string values from the following sources in this order.
    - system.md : a system prompt if it is in the local directory
    - stdin : a prompt if it is available
    - each command line argument in the order they were specified
      - if the argument is a valid url, use Colly to scrape the url and add it to the prompt
      - if the argument is a file name that exists, read the file and add it to the prompt
  - if any of the files cannot be read, return with an error message
  - if any of the urls cannot be downloaded, return with an error message
  - if the prompt is empty, return with an error message
- use pkg/util/files.go to read the files and stdin
- use the pkg/util/scraper.go library to scrape the urls

# Model Selection

- use pkg/api/models.go to get the provider name for the model

- create a main function in main.go that performs the following steps

  - call the processCommandLine function
    - if an error occurs, print the error message and return
  - get the model name from the -model argument or default to anthropic and claude-3.5-sonnet if no model was specified
  - return with an error message is there is a model name and it is not recognized
  - create a new api client for the provider specified by the model name based on the code in pkg/api/client.go
  - call the QueryText method on the client with the prompt and the model name
  - print the response to stdout. no comments or heading, just the response text
  - handle any errors

- use sqirvy-query and sqirvy-review as models
