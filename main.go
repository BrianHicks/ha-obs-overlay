package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		html, err := os.ReadFile("index.html")
		if err != nil {
			http.Error(w, "Error reading index.html", http.StatusInternalServerError)
			return
		}
		w.Write(html)
	})

	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		fmt.Printf("Error starting server: %s", err)
	}
}
