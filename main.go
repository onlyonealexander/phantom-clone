package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"

	"question-system/handlers"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env")
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

	log.Println("Server running at http://localhost:8080/pages")

	log.Fatal(http.ListenAndServe(":8080", nil))
}