package server

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

// podcastHandler handles requests to the /podcast route.
func podcastHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Podcast page!")
}

// defaultHandler is a fallback handler for other routes.
func defaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func StartServer() {
	// Define a command-line flag "port" with a default value of "8000".
	port := flag.String("port", "8000", "Port to run the HTTP server on")
	flag.Parse()

	// Create a new ServeMux.
	mux := http.NewServeMux()

	// Register the /podcast route.
	mux.HandleFunc("/podcast", podcastHandler)
	// Optionally, register a default handler for all other routes.
	mux.HandleFunc("/", defaultHandler)

	// Construct the address string (e.g., ":8000").
	addr := ":" + *port
	log.Printf("Starting server on %s", addr)
}
