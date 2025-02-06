package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Struct for a response message
type Response struct {
	Message string `json:"message"`
}

// Function to handle HTTP requests
func userHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := Response{Message: "User Service is Running!"}
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/user", userHandler) // Route: localhost:8080/user

	fmt.Println("Starting User Service on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil)) // Start the server
}
