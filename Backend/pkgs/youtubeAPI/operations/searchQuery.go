// This will code setup the inital code for calling onto the Youtube API.
// The purpose of the this code is give a search quuery it return a results that match the query.

// TODO: The goal will soon to be get to receive a custom search query as input and have more control of the result counter

package operations

import (
	"Backend/pkgs/youtubeAPI/settings"
	"context"
	"encoding/json"
	"flag"
	"fmt"

	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

// Command-line flags allow the user to override the default search query and result limit.
var (
	searchQuery = flag.String("query", "Google", "Search term to use when querying YouTube")
	resultLimit = flag.Int64("max-results", 25, "Maximum number of results to return from YouTube")
)

// SearchResults holds the results of a YouTube search.
// 'Videos' and 'Channels' are maps where the keys are IDs and the values are titles.
type SearchResults struct {
	Videos   map[string]string `json:"videos"`
	Channels map[string]string `json:"channels"`
}

// SearchYouTubeAPI initializes the YouTube client, performs a search using the specified query,
// and returns the results (videos and channels) as a prettified JSON string.
func SearchYouTubeAPI() (string, error) {
	// Parse command-line flags (if not already parsed in main).
	flag.Parse()

	// Retrieve the YouTube API key from the environment using the config helper.
	apiKey, err := settings.GetAPIKey()
	if err != nil {
		return "", err
	}

	// Create a new YouTube service client.
	ctx := context.Background()
	service, err := youtube.NewService(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return "", fmt.Errorf("error creating YouTube client: %v", err)
	}

	// Perform the search to fetch videos and channels.
	videos, channels, err := fetchSearchResults(service, *searchQuery, *resultLimit)
	if err != nil {
		return "", fmt.Errorf("error searching YouTube: %v", err)
	}

	// Package the results into a SearchResults struct.
	results := SearchResults{
		Videos:   videos,
		Channels: channels,
	}

	// Marshal the results struct into an indented JSON string.
	jsonData, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		return "", fmt.Errorf("error marshaling JSON: %v", err)
	}

	return string(jsonData), nil
}

// fetchSearchResults sends a search request to the YouTube API using the given query and limit,
// and returns two maps: one for videos and one for channels (with their IDs and titles).
func fetchSearchResults(service *youtube.Service, query string, limit int64) (map[string]string, map[string]string, error) {
	// Build the API request with "id" and "snippet" parts.
	call := service.Search.List([]string{"id", "snippet"}).
		Q(query).         // Set the search query.
		MaxResults(limit) // Set the maximum number of results.

	// Execute the API call.
	response, err := call.Do()
	if err != nil {
		return nil, nil, err
	}

	// Prepare maps to store video and channel results.
	videos := make(map[string]string)
	channels := make(map[string]string)

	// Loop through each item in the response and store it based on its type.
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

// FetchChannelPlaylists searches for a channel using the provided query,
// then retrieves and returns the channel's ID along with its playlists (ID and title).
func FetchChannelPlaylists(service *youtube.Service, query string, limit int64) (string, map[string]string, error) {
	// Build a search request to find the channel (type "channel").
	channelSearchCall := service.Search.List([]string{"id", "snippet"}).
		Q(query).        // Set the channel search query.
		Type("channel"). // Restrict results to channels.
		MaxResults(1)    // We only need the first match.

	// Execute the channel search request.
	channelSearchResponse, err := channelSearchCall.Do()
	if err != nil {
		return "", nil, fmt.Errorf("error searching for channel: %v", err)
	}
	if len(channelSearchResponse.Items) == 0 {
		return "", nil, fmt.Errorf("no channel found for query: %s", query)
	}

	// Extract the channel ID from the first result.
	channelID := channelSearchResponse.Items[0].Id.ChannelId

	// Build a request to fetch playlists for the identified channel.
	playlistsCall := service.Playlists.List([]string{"id", "snippet"}).
		ChannelId(channelID). // Set the channel ID.
		MaxResults(limit)     // Limit the number of playlists returned.

	// Execute the playlists request.
	playlistsResponse, err := playlistsCall.Do()
	if err != nil {
		return channelID, nil, fmt.Errorf("error fetching playlists: %v", err)
	}

	// Prepare a map to store playlist IDs and their titles.
	playlists := make(map[string]string)
	for _, item := range playlistsResponse.Items {
		playlists[item.Id] = item.Snippet.Title
	}

	return channelID, playlists, nil
}

// ChannelPlaylistsResponse defines the JSON structure for the playlists response.
type ChannelPlaylistsResponse struct {
	ChannelID string            `json:"channel_id"`
	Playlists map[string]string `json:"playlists"`
}

// GetChannelPlaylistsJSON is a convenience function that:
//  1. Retrieves the API key using the config helper.
//  2. Creates a YouTube service client.
//  3. Fetches the channel's playlists using the provided channel name and result limit.
//  4. Returns a JSON-encoded string containing the channel ID and playlists.
func GetChannelPlaylistsJSON(channelName string, max int64) (string, error) {
	// Retrieve the API key using the config helper.
	apiKey, err := settings.GetAPIKey()
	if err != nil {
		return "", err
	}

	// Create a new YouTube service client.
	ctx := context.Background()
	service, err := youtube.NewService(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return "", fmt.Errorf("error creating YouTube client: %v", err)
	}

	// Fetch the channel ID and its playlists.
	channelID, playlists, err := FetchChannelPlaylists(service, channelName, max)
	if err != nil {
		return "", err
	}

	// Package the response into a ChannelPlaylistsResponse struct.
	resp := ChannelPlaylistsResponse{
		ChannelID: channelID,
		Playlists: playlists,
	}

	// Marshal the struct into an indented JSON string.
	jsonData, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		return "", fmt.Errorf("error marshaling JSON: %v", err)
	}

	return string(jsonData), nil
}
