package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, my name is Alexandr, this is my first Go Website")
}

func main() {
	http.HandleFunc("/", homeHandler)

	fmt.Println("Server started on localhost:http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server error", err)
	}
}
