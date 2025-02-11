create a new file, pkg/sqirvy/deepseek.go to support DeepSeek LLM models, that does the following:
- model the new file after the openai  interface in pkg/sqirvy/openai.go
- deepseek uses the openai  interface
- the API key is in environment variable LLAMA_API_KEY
- the endpoint is in environment variable LLAMA_BASE_URL
- the new file should have the same functionality as the openai.go file, but with the deepseek specs.

update pkg/sqirvy/models.go to include a client for deepseek

add a test file, deepseek_test.go, for the new deepseek client that is similar to the openai_test.go code

----------------------------------------
in directory cmd/deepseek, create a simple go program that is similar to cmd/openai.go but that uses a deepseek client instead of the openai client. 

-----------------------------------------------------------------------------------------------------
i want to create a python program that 
- takes a web url as input on the command line
- scrapes the web url
- generates a summary of the web site, 
- create a new web file that will display the summary
- uses html, css and js in a single index.html file
- creates a python web server for that index.html
- starts the web sever