package main

// ModelsResponse represents the response for the /models endpoint
type ModelsResponse struct {
	Models []ModelInfo `json:"models"`
}

// ModelInfo contains information about a single model
type ModelInfo struct {
	Model    string `json:"name"`
	Provider string `json:"provider"`
}

// QueryRequest represents the request body for the /query endpoint
type QueryRequest struct {
	Model       string  `json:"model"`
	Prompt      string  `json:"prompt"`
	Temperature float32 `json:"temperature"`
}

// QueryResponse represents the response from the /query endpoint
type QueryResponse struct {
	Result string `json:"result"`
}
