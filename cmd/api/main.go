package main

import (
	"encoding/json"
	"log"
	"net/http"
)

const port = ":8080"

type HealthResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	response := HealthResponse{
		Status:  "ok",
		Message: "API is working properly",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/", healthHandler)

	log.Printf("Starting API server on port %s\n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
