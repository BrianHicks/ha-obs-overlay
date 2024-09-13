package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func proxyHomeAssistantState(state string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		url := fmt.Sprintf("http://homeassistant.local:8123/api/states/%s", state)

		// Create a new HTTP request
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Fatalf("Error creating request: %v", err)
		}

		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("HA_TOKEN")))
		req.Header.Set("Content-Type", "application/json")

		// Create an HTTP client and send the request
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Fatalf("Error making request: %v", err)
		}
		defer resp.Body.Close()

		w.Header().Set("Content-Type", resp.Header.Get("Content-Type"))

		// Read the response body
		_, err = io.Copy(w, resp.Body)
		if err != nil {
			http.Error(w, "Error copying HA response", http.StatusInternalServerError)
		}
	}
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		html, err := os.ReadFile("index.html")
		if err != nil {
			http.Error(w, "Error reading index.html", http.StatusInternalServerError)
			return
		}
		w.Write(html)
	})

	http.HandleFunc("/state/steps", proxyHomeAssistantState("sensor.brians_iphone_steps"))
	http.HandleFunc("/state/co2", proxyHomeAssistantState("sensor.aranet4_1eef2_carbon_dioxide"))

	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		fmt.Printf("Error starting server: %s", err)
	}
}
