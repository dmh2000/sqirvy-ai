package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

func main() {
	// Parse command line flags
	addr := flag.String("addr", ":8081", "HTTP server address")
	flag.Parse()

	// Parse templates
	var err error
	templates, err = template.ParseGlob("web/sqirvy-xyz/templates/*.html")
	if err != nil {
		log.Fatalf("Failed to parse templates: %v", err)
	}

	// Serve static files
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Setup routes
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/about", handleAbout)

	// Start server
	log.Printf("Starting server on %s", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	if err := templates.ExecuteTemplate(w, "home.html", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handleAbout(w http.ResponseWriter, r *http.Request) {
	if err := templates.ExecuteTemplate(w, "about.html", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
