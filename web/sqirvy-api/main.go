package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"

	"sqirvy-ai/pkg/sqirvy"
)

func main() {
	// Parse command line flags
	addr := flag.String("addr", ":8080", "HTTP server address")
	flag.Parse()

	// Create handlers
	http.HandleFunc("/models", handleModels)
	http.HandleFunc("/query", handleQuery)

	// Start server
	log.Printf("Starting server on %s", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

func handleModels(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	log.Printf("Handling models request from %s", r.RemoteAddr)
	if r.Method != http.MethodGet {
		log.Printf("Method not allowed: %s", r.Method)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Create response with models and their providers
	response := ModelsResponse{
		Models: make([]ModelInfo, 0, len(sqirvy.ModelToProvider)),
	}

	for model, provider := range sqirvy.ModelToProvider {
		response.Models = append(response.Models, ModelInfo{
			Name:     model,
			Provider: provider,
		})
	}

	// Send JSON response
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func handleQuery(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	log.Printf("Handling query request from %s", r.RemoteAddr)
	if r.Method != http.MethodPost {
		log.Printf("Method not allowed: %s", r.Method)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse request body
	var req QueryRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate request
	if req.Prompt == "" {
		http.Error(w, "Prompt cannot be empty", http.StatusBadRequest)
		return
	}

	// Get provider for the model
	provider, err := sqirvy.GetProviderName(req.Model)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid model: %v", err), http.StatusBadRequest)
		return
	}

	// Create client for the provider
	client, err := sqirvy.NewClient(sqirvy.Provider(provider))
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to create client: %v", err), http.StatusInternalServerError)
		return
	}
	defer client.Close()

	// Query the model
	result, err := client.QueryText(r.Context(), req.Prompt, req.Model, sqirvy.Options{
		Temperature: req.Temperature,
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("Query failed: %v", err), http.StatusInternalServerError)
		return
	}

	// Send response
	response := QueryResponse{
		Result: result,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
