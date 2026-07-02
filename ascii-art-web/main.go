package main

import (
	"fmt"
	"net/http"
	"ps/handlers"
)

func main() {
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/ascii-art", handlers.ASCIIArtHandler)

	fmt.Println("Server started at http://localhost:8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server error:", err)
	}
}
