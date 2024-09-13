package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello, World!")
	})

	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		fmt.Printf("Error starting server: %s", err)
	}
}
