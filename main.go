package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		notFoundHandler(w, r)
		return
	}
	fmt.Fprint(w, "Hello! My name is Alexander and this is my first project in Golang")
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "OK")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "About me: I'm Technical Support Engineer and now I'm learning Go and building my first personal website")
}

func projectHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Projects: Go Visit Site, Tuxback and Campus bot")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Contacts: Telegram, WhatsApp, GitHub, LinkedIn, Email")
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "404 page not found")
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/projects", projectHandler)
	http.HandleFunc("/contacts", contactHandler)
	http.HandleFunc("/health", healthHandler)
	fmt.Println("Server started on http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server error", err)
	}
}
