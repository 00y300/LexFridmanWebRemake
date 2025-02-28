// This will code setup the inital code for calling onto the Youtube API.
// The purpose of the this code is give a search quuery it return a results that match the query.

// TODO: The goal will soon to be get to receive a custom search query as input and have more control of the result counter

package operations

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

// Command-line flags allow the user to override the default search term
// and the maximum number of results returned from YouTube.
var (
	searchQuery = flag.String("query", "Google", "Search term to use when querying YouTube")
	resultLimit = flag.Int64("max-results", 25, "Maximum number of results to return from YouTube")
)

// SearchResults holds the results of a YouTube search.
// Videos and Channels maps have IDs as keys and their corresponding titles as values.
type SearchResults struct {
	Videos   map[string]string `json:"videos"`
	Channels map[string]string `json:"channels"`
}

// SearchYouTubeAPI initializes the YouTube API client, performs a search
// using the command-line provided query, and returns the results as a JSON string.
func SearchYouTubeAPI() (string, error) {
	// Parse command-line flags. (Call once in the application lifecycle.)
	flag.Parse()

	// Load environment variables from a .env file.
	if err := godotenv.Load(".env"); err != nil {
		return "", fmt.Errorf("error loading .env file: %v", err)
	}

	// Retrieve the YouTube API key from the environment.
	apiKey := os.Getenv("GOOGLE_API")
	if apiKey == "" {
		return "", fmt.Errorf("GOOGLE_API key is not set in the .env file")
	}

	// Create a new YouTube service client with the provided API key.
	ctx := context.Background()
	service, err := youtube.NewService(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return "", fmt.Errorf("error creating YouTube client: %v", err)
	}

	// Fetch videos and channels using the helper function.
	videos, channels, err := fetchSearchResults(service, *searchQuery, *resultLimit)
	if err != nil {
		return "", fmt.Errorf("error searching YouTube: %v", err)
	}

	// Package the search results into a struct.
	results := SearchResults{
		Videos:   videos,
		Channels: channels,
	}

	// Marshal the results struct into a prettified JSON string.
	jsonData, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		return "", fmt.Errorf("error marshaling JSON: %v", err)
	}

	return string(jsonData), nil
}

// fetchSearchResults sends a search request to the YouTube API using the given query and limit.
// It returns two maps: one for videos and one for channels (with their IDs and titles).
func fetchSearchResults(service *youtube.Service, query string, limit int64) (map[string]string, map[string]string, error) {
	// Build the API call with required parts: "id" and "snippet".
	call := service.Search.List([]string{"id", "snippet"}).
		Q(query).         // Set the search query.
		MaxResults(limit) // Set the maximum number of results.

	// Execute the API call.
	response, err := call.Do()
	if err != nil {
		return nil, nil, err
	}

	// Initialize maps to store video and channel results.
	videos := make(map[string]string)
	channels := make(map[string]string)

	// Loop over each returned item and add it to the correct map.
	for _, item := range response.Items {
		switch item.Id.Kind {
		case "youtube#video":
			videos[item.Id.VideoId] = item.Snippet.Title
		case "youtube#channel":
			channels[item.Id.ChannelId] = item.Snippet.Title
		}
	}

	return videos, channels, nil
}

// FetchChannelPlaylists searches for a channel by name and retrieves all its playlists.
// It returns the channel ID along with a map of playlist IDs to their titles.
func FetchChannelPlaylists(service *youtube.Service, query string, limit int64) (string, map[string]string, error) {
	// Search for the channel that matches the provided query.
	channelSearchCall := service.Search.List([]string{"id", "snippet"}).
		Q(query).        // Set the channel search query.
		Type("channel"). // Restrict the search to channels.
		MaxResults(1)    // We only need one matching channel.

	channelSearchResponse, err := channelSearchCall.Do()
	if err != nil {
		return "", nil, fmt.Errorf("error searching for channel: %v", err)
	}
	if len(channelSearchResponse.Items) == 0 {
		return "", nil, fmt.Errorf("no channel found for query: %s", query)
	}

	// Extract the channel ID from the first search result.
	channelID := channelSearchResponse.Items[0].Id.ChannelId

	// Retrieve the playlists for the identified channel using its channel ID.
	playlistsCall := service.Playlists.List([]string{"id", "snippet"}).
		ChannelId(channelID). // Set the channel ID to retrieve playlists.
		MaxResults(limit)     // Limit the number of playlists returned.
	playlistsResponse, err := playlistsCall.Do()
	if err != nil {
		return channelID, nil, fmt.Errorf("error fetching playlists: %v", err)
	}

	// Create a map to hold playlist IDs and their corresponding titles.
	playlists := make(map[string]string)
	for _, item := range playlistsResponse.Items {
		playlists[item.Id] = item.Snippet.Title
	}

	return channelID, playlists, nil
}

// ChannelPlaylistsResponse represents the JSON structure for channel playlists responses.
type ChannelPlaylistsResponse struct {
	ChannelID string            `json:"channel_id"`
	Playlists map[string]string `json:"playlists"`
}

// GetChannelPlaylistsJSON is a helper function that:
// 1. Loads environment variables and creates a YouTube service client.
// 2. Calls FetchChannelPlaylists to get the channel's playlists.
// 3. Returns the channel ID and playlists as a formatted JSON string.
func GetChannelPlaylistsJSON(channelName string, max int64) (string, error) {
	// Load the .env file to get the API key.
	if err := godotenv.Load(".env"); err != nil {
		return "", fmt.Errorf("error loading .env file: %v", err)
	}

	// Retrieve the API key from the environment.
	apiKey := os.Getenv("GOOGLE_API")
	if apiKey == "" {
		return "", fmt.Errorf("GOOGLE_API key is not set in the .env file")
	}

	// Create a new YouTube service client.
	ctx := context.Background()
	service, err := youtube.NewService(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return "", fmt.Errorf("error creating YouTube client: %v", err)
	}

	// Fetch the channel ID and its playlists using the provided channel name.
	channelID, playlists, err := FetchChannelPlaylists(service, channelName, max)
	if err != nil {
		return "", err
	}

	// Package the results into a response struct.
	resp := ChannelPlaylistsResponse{
		ChannelID: channelID,
		Playlists: playlists,
	}

	// Marshal the response struct into an indented JSON string.
	jsonData, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		return "", fmt.Errorf("error marshaling JSON: %v", err)
	}

	return string(jsonData), nil
}
