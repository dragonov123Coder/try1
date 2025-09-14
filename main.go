package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	http.HandleFunc("/api/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"message": "Hello from Go backend!"}`)
	})

	port := os.Getenv("PORT") // Railway provides this
	if port == "" {
		port = "8080" // default for local testing
	}

	fmt.Println("Server running on port " + port)
	http.ListenAndServe(":"+port, nil)
}
