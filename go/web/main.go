// Package main implements the web interface for the password generator.
package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	pwd "sqirvy.xyz/sqirvy_ai/internal/pwd"
)

//go:embed static/*
var static embed.FS

// PasswordResponse is the struct for the JSON response containing generated passwords.
type PasswordResponse struct {
	Passwords []Password `json:"passwords"`
}

// Password is the struct for a single password, including raw and formatted versions.
type Password struct {
	Raw       string `json:"raw"`
	Formatted string `json:"formatted"`
}

func main() {
	// Parse template
	tmpl, err := template.ParseFS(static, "static/index.html")
	if err != nil {
		log.Fatal(err)
	}

	// Serve static files
	http.Handle("/static/", http.FileServer(http.FS(static)))

	// Serve index page
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, nil)
	})

	// Handle password generation
	http.HandleFunc("/generate", func(w http.ResponseWriter, r *http.Request) {
		lengthStr := r.URL.Query().Get("length")
		countStr := r.URL.Query().Get("count")
		symbolsStr := r.URL.Query().Get("symbols")

		length := 24
		if lengthStr != "" {
			parsedLength, err := strconv.Atoi(lengthStr)
			if err != nil {
				http.Error(w, "Invalid length parameter", http.StatusBadRequest)
				return
			}
			if parsedLength >= 24 {
				length = parsedLength
			}
		}

		count := 1
		if countStr != "" {
			parsedCount, err := strconv.Atoi(countStr)
			if err != nil {
				http.Error(w, "Invalid count parameter", http.StatusBadRequest)
				return
			}
			if parsedCount >= 1 {
				count = parsedCount
			}
		}

		symbols := symbolsStr == "true"

		passwords := make([]Password, count)
		for i := 0; i < count; i++ {
			raw, err := pwd.GeneratePassword(length, symbols)
			if err != nil {
				http.Error(w, fmt.Sprintf("Error generating password: %v", err), http.StatusInternalServerError)
				return
			}
			formatted, err := pwd.FormatPassword(raw)
			if err != nil {
				http.Error(w, fmt.Sprintf("Error formatting password: %v", err), http.StatusInternalServerError)
				return
			}
			passwords[i] = Password{
				Raw:       raw,
				Formatted: formatted,
			}
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(PasswordResponse{Passwords: passwords})
	})

	log.Println("Server starting on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
