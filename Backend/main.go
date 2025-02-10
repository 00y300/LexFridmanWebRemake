package main

import (
	"Backend/cmd/web"
	"log"
	"net/http"
	// Adjust this import path based on your module name.
)

func main() {
	// Initialize the router from the web package.
	router := web.NewRouter()

	// Define the address and port.
	addr := ":8080"
	log.Printf("Server is running on %s", addr)

	// Start the HTTP server.
	if err := http.ListenAndServe(addr, router); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
