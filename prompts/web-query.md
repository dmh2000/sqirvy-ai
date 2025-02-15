It took three prompts to get the web app working The code was generate using Aider and claude-3-sonnet-20240229.

- in directory web/query, create a web app and an associated web server that allows users to query a language model using a prompt. there whould be three endpoints: "anthropic", "openai", and "gemini". the anthropic endpoint should use the claude-3-sonnet-20240229 model, the openai endpoint should use the gpt-4-turbo-preview model, and the gemini endpoint should use the gemini-2.0-flash-exp model. the web app should have a form that allows users to enter a prompt and submit it to the appropriate endpoint. the web app should display the response from the language model. the web server should serve static files from the static directory and handle api requests for the anthropic, openai, and gemini endpoints. the server should listen on port 8080 and log messages to the console.

- the code in web/query/main.go in web/query/main.go, instead of having a menu of options, the program should h
  ave a single text box to enter a prompt, and should show the results of each of the three providers so they can
  be compared side by side.

- in web/query, the html page was not modified to support the new layout. can you fix that


- in directory examples/web/compare, create a web server that allows users to query a language model using a prompt. 
- the web app server will written in Go. 
- the web app should have two API endpoints. 
  - One endpoint to request a list of available models. It can use the functions in pkg/sqirvy/models.go to get the list of models and providers. 
  - The second endpoint should receive a prompt and a model name, and return the result of the query to the user. This endpoint should use the sqirvy.ModelToProvider function to get the appropriate Client, then use sqirvy.QueryText function for the appropriate provider to get the result.
- the web app server will also server static files from the static directory.
