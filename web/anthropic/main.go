package main

import (
	"encoding/json"
	"log"
	"net/http"
	api "sqirvyllm/pkg/api"
)

type QueryRequest struct {
	Prompt string `json:"prompt"`
}

type QueryResponse struct {
	Result string `json:"result"`
	Error  string `json:"error,omitempty"`
}

func main() {
	// Serve static files from the static directory
	fs := http.FileServer(http.Dir("web/anthropic/static"))
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

	client, err := api.NewClient(api.Anthropic)
	if err != nil {
		sendJSONResponse(w, QueryResponse{Error: err.Error()}, http.StatusInternalServerError)
		return
	}

	result, err := client.QueryText(prompt, "claude-3-sonnet-20240229", api.Options{})
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
