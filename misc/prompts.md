in examples/web/sqirvy-web, the user interface is fixed to use anthropic,openai and google gemini in each
 of the three output display sections. can you change it so that the user can select any provider and model for each section. there should be a dropdown list in each section to selct a model. the title of the section should be updated to show the name of the provider. the list of models and providers is in pkg/sqirvy/models.go. when the web logic performs a query to the backend, it should specify the model name. the back end server should then return the query data for the selected provider and model. it should work the same for all three sections 


 /add
 ./examples/web/sqirvy-web/main.go
 ./examples/web/sqirvy-web/static/index.html
 ./examples/web/sqirvy-web/static/style.css
 ./examples/web/sqirvy-web/static/script.js
 ./pkg/sqirvy/models.go
 ./pkg/sqirvy/anthropic.go
 ./pkg/sqirvy/deepseek.go
 ./pkg/sqirvy/gemini.go
 ./pkg/sqirvy/llama.go
 ./pkg/sqirvy/client.go