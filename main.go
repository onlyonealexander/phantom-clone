package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"question-system/handlers"
)

func main() {
	// Load .env locally. On Railway, this will fail gracefully
	// and the app will use Railway's environment variables.
	if err := godotenv.Load(); err != nil {
		log.Println(".env not found, using Railway environment variables")
	}

	// CSS
	http.Handle("/css/",
		http.StripPrefix(
			"/css/",
			http.FileServer(http.Dir("css")),
		),
	)

	// JS
	http.Handle("/js/",
		http.StripPrefix(
			"/js/",
			http.FileServer(http.Dir("js")),
		),
	)

	// Pages
	http.Handle("/pages/",
		http.StripPrefix(
			"/pages/",
			http.FileServer(http.Dir("pages")),
		),
	)

	// Root
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "pages/index.html")
	})

	// API
	http.HandleFunc("/api/question", handlers.SubmitQuestion)

	// Railway provides the PORT environment variable.
	// Use 8080 when running locally.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server running on port %s", port)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}