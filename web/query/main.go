package main

import (
	"encoding/json"
	"log"
	"net/http"
	api "sqirvyllm/pkg/api"
)

type QueryResponse struct {
	Result string `json:"result"`
	Error  string `json:"error,omitempty"`
}

func main() {
	// Serve static files from the static directory
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	// Handle API requests for each provider
	http.HandleFunc("/api/anthropic", handleAnthropicQuery)
	http.HandleFunc("/api/openai", handleOpenAIQuery)
	http.HandleFunc("/api/gemini", handleGeminiQuery)

	log.Println("Server starting on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func handleAnthropicQuery(w http.ResponseWriter, r *http.Request) {
	handleQuery(w, r, api.Anthropic, "claude-3-sonnet-20240229")
}

func handleOpenAIQuery(w http.ResponseWriter, r *http.Request) {
	handleQuery(w, r, api.OpenAI, "gpt-4-turbo-preview")
}

func handleGeminiQuery(w http.ResponseWriter, r *http.Request) {
	handleQuery(w, r, api.Gemini, "gemini-2.0-flash-exp")
}

func handleQuery(w http.ResponseWriter, r *http.Request, provider api.Provider, model string) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	prompt := r.URL.Query().Get("prompt")
	if prompt == "" {
		http.Error(w, "Prompt parameter is required", http.StatusBadRequest)
		return
	}

	client, err := api.NewClient(provider)
	if err != nil {
		sendJSONResponse(w, QueryResponse{Error: err.Error()}, http.StatusInternalServerError)
		return
	}

	result, err := client.QueryText(prompt, model, api.Options{})
	if err != nil {
		sendJSONResponse(w, QueryResponse{Error: err.Error()}, http.StatusInternalServerError)
		return
	}

	sendJSONResponse(w, QueryResponse{Result: result}, http.StatusOK)
}

func sendJSONResponse(w http.ResponseWriter, response QueryResponse, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
}
