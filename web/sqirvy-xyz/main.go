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

	log.Printf("Starting template parsing...")
	// Parse templates
	var err error
	templates, err = template.ParseGlob("./templates/*.html")
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
	log.Printf("Handling home request from %s", r.RemoteAddr)
	if r.URL.Path != "/" {
		log.Printf("404 for path: %s", r.URL.Path)
		http.NotFound(w, r)
		return
	}
	if err := templates.ExecuteTemplate(w, "layout.html", map[string]interface{}{
		"Template": "home.html",
	}); err != nil {
		log.Printf("Error executing home template: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	log.Printf("Home page rendered successfully")
}

func handleAbout(w http.ResponseWriter, r *http.Request) {
	log.Printf("Handling about request from %s", r.RemoteAddr)
	if err := templates.ExecuteTemplate(w, "layout.html", map[string]interface{}{
		"Content": "about.html",
	}); err != nil {
		log.Printf("Error executing about template: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	log.Printf("About page rendered successfully")
}
