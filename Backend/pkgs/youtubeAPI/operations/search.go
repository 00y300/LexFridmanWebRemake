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

// Flags for command-line arguments
var (
	query      = flag.String("query", "Google", "Search term")
	maxResults = flag.Int64("max-results", 25, "Max YouTube results")
)

// SearchResults holds the videos and channels maps.
type SearchResults struct {
	Videos   map[string]string `json:"videos"`
	Channels map[string]string `json:"channels"`
}

// Main performs the YouTube search and returns the JSON result.
func Main() (string, error) {
	// It’s best to parse flags only once. If your server has already parsed them,
	// you might consider removing this flag.Parse() call.
	flag.Parse()

	if err := godotenv.Load(".env"); err != nil {
		return "", fmt.Errorf("error loading .env file: %v", err)
	}

	developerKey := os.Getenv("GOOGLE_API")
	if developerKey == "" {
		return "", fmt.Errorf("GOOGLE_API key is not set in the .env file")
	}

	ctx := context.Background()
	service, err := youtube.NewService(ctx, option.WithAPIKey(developerKey))
	if err != nil {
		return "", fmt.Errorf("error creating new YouTube client: %v", err)
	}

	videos, channels, err := searchYouTube(service, *query, *maxResults)
	if err != nil {
		return "", fmt.Errorf("error searching YouTube: %v", err)
	}

	results := SearchResults{
		Videos:   videos,
		Channels: channels,
	}

	jsonData, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		return "", fmt.Errorf("error marshaling JSON: %v", err)
	}

	return string(jsonData), nil
}

func searchYouTube(service *youtube.Service, q string, max int64) (map[string]string, map[string]string, error) {
	call := service.Search.List([]string{"id", "snippet"}).
		Q(q).
		MaxResults(max)

	response, err := call.Do()
	if err != nil {
		return nil, nil, err
	}

	videos := make(map[string]string)
	channels := make(map[string]string)

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
