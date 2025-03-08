package server

import (
	"Backend/pkgs/youtubeAPI/operations"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// podcastHandler handles requests to the /podcast route.
func podcastHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the Podcast page!"))
}

// YoutubeSearchHandle handles GET requests to the /yt route.
func YoutubeSearchHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// Call the YouTube search function (make sure this function exists in your operations package)
		jsonData, err := operations.SearchYouTubeAPI()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(jsonData))
		return
	}
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

// YoutubeChannelPlaylistsHandle handles GET requests to /yt/playlists,
// returning a list of playlists for a given channel.
func YoutubeChannelPlaylistsHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract query parameters: ?channel=<channel_name>&max=<number>
	channelName := r.URL.Query().Get("channel")
	if channelName == "" {
		http.Error(w, "Missing 'channel' query parameter", http.StatusBadRequest)
		return
	}

	maxParam := r.URL.Query().Get("max")
	var maxResults int64 = 5 // default
	if maxParam != "" {
		if val, err := strconv.ParseInt(maxParam, 10, 64); err == nil {
			maxResults = val
		}
	}

	// Call our helper to fetch playlists as JSON
	jsonData, err := operations.GetChannelPlaylistsJSON(channelName, maxResults)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching playlists: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(jsonData))
}

// YoutubePlaylistVideosHandle handles GET requests to /yt/playlist,
// returning all videos from a given playlist as JSON.
// It expects query parameters: ?channel=<channelID>&playlist=<playlistID>
func YoutubePlaylistVideosHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract query parameters for channel and playlist IDs.
	channelID := r.URL.Query().Get("channel")
	playlistID := r.URL.Query().Get("playlist")
	if channelID == "" || playlistID == "" {
		http.Error(w, "Missing 'channel' or 'playlist' query parameter", http.StatusBadRequest)
		return
	}

	// Call the helper to fetch all videos in the playlist as JSON.
	jsonData, err := operations.FetchAllVideosAsJSON(channelID, playlistID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching playlist videos: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(jsonData))
}

// defaultHandler is a fallback handler for other routes.
func defaultHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hi there, I love " + r.URL.Path[1:] + "!"))
}

// corsMiddleware sets the CORS headers for each request.
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Allow all origins and specific methods/headers.
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Handle preflight OPTIONS requests.
		if r.Method == http.MethodOptions {
			return
		}
		next.ServeHTTP(w, r)
	})
}

// StartServer initializes and starts the HTTP server.
func StartServer() {
	port := flag.String("port", "8000", "Port to run the HTTP server on")
	flag.Parse()

	mux := http.NewServeMux()

	// Register route handlers.
	mux.HandleFunc("/podcast", podcastHandler)
	mux.HandleFunc("/yt", YoutubeSearchHandle)
	mux.HandleFunc("/yt/playlists", YoutubeChannelPlaylistsHandle)
	mux.HandleFunc("/yt/playlist", YoutubePlaylistVideosHandle)
	mux.HandleFunc("/", defaultHandler)

	// Apply CORS middleware.
	handlerWithCORS := corsMiddleware(mux)
	addr := ":" + *port
	log.Printf("Starting server on %s", addr)

	// Start listening and serving requests.
	if err := http.ListenAndServe(addr, handlerWithCORS); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
