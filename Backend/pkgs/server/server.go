package server

import (
	"Backend/pkgs/youtubeAPI/operations"
	"flag"
	"log"
	"net/http"
)

// podcastHandler handles requests to the /podcast route.
func podcastHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the Podcast page!"))
}

// YoutubeSearchHandle handles GET requests to the /yt route.
func YoutubeSearchHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// Call the operations.Main() to perform the search.
		jsonData, err := operations.Main()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Set content type and write the JSON response.
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(jsonData))
		return
	}

	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

// defaultHandler is a fallback handler for other routes.
func defaultHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hi there, I love " + r.URL.Path[1:] + "!"))
}

// corsMiddleware sets the CORS headers for each request.
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Replace "*" with specific domain(s) in production.
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Handle preflight OPTIONS request quickly and return.
		if r.Method == http.MethodOptions {
			return
		}

		// Proceed to the next handler if not OPTIONS.
		next.ServeHTTP(w, r)
	})
}

func StartServer() {
	// Define a command-line flag "port" with a default value of "8000".
	port := flag.String("port", "8000", "Port to run the HTTP server on")
	flag.Parse()

	// Create a new ServeMux.
	mux := http.NewServeMux()

	// Register the /podcast route.
	mux.HandleFunc("/podcast", podcastHandler)

	// Register the /yt route.
	mux.HandleFunc("/yt", YoutubeSearchHandle)

	// Optionally, register a default handler for all other routes.
	mux.HandleFunc("/", defaultHandler)

	// Wrap the mux with the CORS middleware.
	handlerWithCORS := corsMiddleware(mux)

	addr := ":" + *port
	log.Printf("Starting server on %s", addr)

	// Listen and serve with the CORS-enabled handler.
	if err := http.ListenAndServe(addr, handlerWithCORS); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
