create a Go command line application in directory cmd/sqirvy-review with the following features:

# Command Line Processing

- has a command line argment -h that prints a help message and returns
- has a command line argument -model that specifies the model to use.
- any other arguments are considered to be filenames to read from
- construct a string variable named "prompt" that will concatenate the string values from the following sources in this order.
  - each filename argument in the order they were specified
  - do not read from stdin
- if any of the files cannot be read, return with an error message
- if the prompt is empty, return with an error message
- place the command line processing in a separate file named cli.go and function named "processCommandLine"
- return the prompt, model name if any and any errors
- place the command line processing code in a separate file named "cli.go"
- use pkg/util/files.go to read the files and stdin
- use sqirvy-query as a model

# Model Selection

- use pkg/sqirvy/models.go to get the provider name for the model

- create a main function that performs the following steps
  - call the processCommandLine function
    - if an error occurs, print the error message and return
  - get the model name from the -model argument or default to anthropic and claude-3.5-sonnet if no model was specified
  - return with an error message is there is a model name and it is not recognized
  - create a new api client for the provider specified by the model name based on the code in pkg/sqirvy/client.go
  - call the QueryText method on the client with the prompt and the model name
  - print the response to stdout. no comments or heading, just the response text
  - handle any errors
