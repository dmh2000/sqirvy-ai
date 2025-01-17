package main

import (
	"encoding/json"
	"log"
	"net/http"
	sqirvy "sqirvyllm/pkg/sqirvy"
)

type ProviderResponse struct {
	Result string `json:"result"`
	Error  string `json:"error,omitempty"`
}

type QueryResponse struct {
	Anthropic ProviderResponse `json:"anthropic"`
	OpenAI    ProviderResponse `json:"openai"`
	Gemini    ProviderResponse `json:"gemini"`
}

func main() {
	// Serve static files from the static directory
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	// Handle API requests
	http.HandleFunc("/api/query", handleQuery)

	log.Println("Server starting on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func handleQuery(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	prompt := r.URL.Query().Get("prompt")
	if prompt == "" {
		http.Error(w, "Prompt parameter is required", http.StatusBadRequest)
		return
	}

	response := QueryResponse{}

	// Query Anthropic
	if client, err := sqirvy.NewClient(sqirvy.Anthropic); err == nil {
		if result, err := client.QueryText(prompt, "claude-3-sonnet-20240229", sqirvy.Options{}); err == nil {
			response.Anthropic.Result = result
		} else {
			response.Anthropic.Error = err.Error()
		}
	} else {
		response.Anthropic.Error = err.Error()
	}

	// Query OpenAI
	if client, err := sqirvy.NewClient(sqirvy.OpenAI); err == nil {
		if result, err := client.QueryText(prompt, "gpt-4-turbo-preview", sqirvy.Options{}); err == nil {
			response.OpenAI.Result = result
		} else {
			response.OpenAI.Error = err.Error()
		}
	} else {
		response.OpenAI.Error = err.Error()
	}

	// Query Gemini
	if client, err := sqirvy.NewClient(sqirvy.Gemini); err == nil {
		if result, err := client.QueryText(prompt, "gemini-2.0-flash-exp", sqirvy.Options{}); err == nil {
			response.Gemini.Result = result
		} else {
			response.Gemini.Error = err.Error()
		}
	} else {
		response.Gemini.Error = err.Error()
	}

	sendJSONResponse(w, response, http.StatusOK)
}

func sendJSONResponse(w http.ResponseWriter, response QueryResponse, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
}
