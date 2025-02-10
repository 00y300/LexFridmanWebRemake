package web

import (
	"fmt"
	"net/http"
)

// NewRouter sets up the HTTP routes and returns an http.Handler.
func NewRouter() http.Handler {
	mux := http.NewServeMux()

	// Define routes.
	mux.HandleFunc("/", HomeHandler)
	mux.HandleFunc("/about", AboutHandler)

	// Add more routes as needed.
	return mux
}

// HomeHandler handles requests to the root URL.
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the Home Page!")
}

// AboutHandler handles requests to the /about URL.
func AboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is the About Page!")
}
