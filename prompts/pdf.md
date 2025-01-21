- create a go library in file pkg/util/pdf.go, that does the following:
  - create a function, ExtractPdf, that gets pdf binary data input as []byte and outputs the pdf text. do this by invoking the command line program pdftotext
  - create a function SummarizePdf, that generates a summary of the pdf using  the pkg/sqirvy library
- make sure the file has the following documentation:
  - a package comment
  - a function comment for each function
  - a comment for each parameter
  - a comment for each return value
  - a comment for each error condition
  - an example usage of each function
  - relevant comments inside the functions

- add tests to pkg/util for pdf.go. include tests that download valid urls, invalid urls and non-existent urls. test both ExtractPdf and SummarizePdf

-----------------------------------------------------------------------------------------------------
create a Go command line application in directory cmd/sqirvy-pdf with the following features:

# Command Line Processing

- has a command line argment -h that prints a help message and returns
- has a command line argument -m that specifies the model to use.
- any other arguments are considered to be filenames to read from
- returns 
  - the model name if any
  - data from stdin if it is available
  - a list of filenames argument in the order they were specified
- place the command line processing code in a separate file named "cli.go"
- use pkg/util/files.go for reading input as needed

# Main Program

- create a main function that performs the following steps
  - call the processCommandLine function
    - if an error occurs, print the error message and return
  - get the model name argument or default to anthropic and claude-3.5-sonnet if no model was specified
  - return with an error message is there is a model name and it is not recognized
  - for each of  stdin data and each filename argument
      - call pkg/util/pdf.go to extract the text from the data
      - concatenate the output into a single variablenamed 'context'. wrap each separate data source in tripple backtics named 'pdf'
  - create a new api client for the provider specified by the model name based on the code in pkg/sqirvy/client.go
  - call client.queryText on the concatenated data
  - print the response to stdout. no comments or heading, just the response text
  - handle any errors
