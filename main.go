package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, my name is Alexandr, this is my first Go Website")
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

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/projects", projectHandler)
	http.HandleFunc("/contacts", contactHandler)
	fmt.Println("Server started on http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server error", err)
	}
}
