It took three prompts to get the web app working The code was generate using Aider and claude-3-sonnet-20240229.

- in directory web/query, create a web app and an associated web server that allows users to query a language model using a prompt. there whould be three endpoints: "anthropic", "openai", and "gemini". the anthropic endpoint should use the claude-3-sonnet-20240229 model, the openai endpoint should use the gpt-4-turbo-preview model, and the gemini endpoint should use the gemini-2.0-flash-exp model. the web app should have a form that allows users to enter a prompt and submit it to the appropriate endpoint. the web app should display the response from the language model. the web server should serve static files from the static directory and handle api requests for the anthropic, openai, and gemini endpoints. the server should listen on port 8080 and log messages to the console.

- the code in web/query/main.go in web/query/main.go, instead of having a menu of options, the program should h
  ave a single text box to enter a prompt, and should show the results of each of the three providers so they can
  be compared side by side.

- in web/query, the html page was not modified to support the new layout. can you fix that


- in directory web/sqirvy-api, create a web server that allows users to query a language model using a prompt. 
- the web app server will written in Go using the net/http package. 
- the web app should have two API endpoints. 
  - One endpoint to request a list of available models. It can use the functions in pkg/sqirvy/models.go to get the list of models and providers. 
  - The second endpoint should receive a prompt and a model name from the caller, and return the result of the query to the user. This endpoint should use the sqirvy.ModelToProvider function to find the Provider, then create the appropriate Client, then use sqirvy.QueryText function from the Client to get the result.
- endpoints should use JSON for requests and responses
- the server does not need authentication
- the server should take a command line argument for the listen address and port. The default should be ":8080"
- the server should import the local package ""sqirvy-ai/pkg/sqirvy" for the required functions

/add
 pkg/sqirvy/models.go
 pkg/sqirvy/anthropic.go
 pkg/sqirvy/openai.go
 pkg/sqirvy/deepseek.go
 pkg/sqirvy/gemini.go
 pkg/sqirvy/llama.go
 pkg/sqirvy/client.go
 web/sqirvy-api/main.go
 web/sqirvy-xyz/main.go
 web/sqirvy-api/types.go


in web/sqirvy-api, add tests using the go test framework. the tests should assume the api 
endpoint is a localhost:8080. the tests should cover requesting the list of models and
performing queries using the Anthropic, Gemini and Openai modeles. the tests should be table driven.

in directory web/sqirvy-xyz, create a web application that allows users to query three different ai models and compare the results. there should be a single text box to enter a prompt and a button to submit requests. the results should be displayed in separate areas. each of these areas should have a title, a drop down list to select a model, and a text box to display the results. 
When the page is loaded, the home page should populate the drop down lists with the available models using the web/sqirvy-api/models endpoint. the user should be able to select any available model in any of the three areas.
when the user selects three models and clicks the submit button, the app should use the web/sqirvy-api request endpoing to execute the query for the selected model and display the results. the web app should use the go net/http package. The web page should be clean, simple and modern. It should have a label at the top of the page called "sqiryv.xyz". the web app should also have an 'about' page that has a description of the app. the home page should have a button to access the 'about' page. it should have a web server that serves the static. files. 


                

