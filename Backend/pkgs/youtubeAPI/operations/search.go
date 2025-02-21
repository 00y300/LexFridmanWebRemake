// This will code setup the inital code for calling onto the Youtube API.
// The purpose of the this code is give a search quuery it return a results that match the query.

// TODO: The goal will soon to be get to receive a custom search query as input and have more control of the result counter

package operations

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

// Flags for command-line arguments.
var (
	query      = flag.String("query", "Google", "Search term")
	maxResults = flag.Int64("max-results", 25, "Max YouTube results")
)

// SearchResults holds the videos and channels maps.
type SearchResults struct {
	Videos   map[string]string `json:"videos"`
	Channels map[string]string `json:"channels"`
}

func main() {
	// Parse command-line flags.
	flag.Parse()

	// Load environment variables from .env.
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	developerKey := os.Getenv("GOOGLE_API")
	if developerKey == "" {
		log.Fatal("GOOGLE_API key is not set in the .env file")
	}

	// Create a context.
	ctx := context.Background()

	// Create the YouTube service using the context and API key option.
	service, err := youtube.NewService(ctx, option.WithAPIKey(developerKey))
	if err != nil {
		log.Fatalf("Error creating new YouTube client: %v", err)
	}

	// Call the search function.
	videos, channels, err := searchYouTube(service, *query, *maxResults)
	if err != nil {
		log.Fatalf("Error searching YouTube: %v", err)
	}

	// Combine the results.
	results := SearchResults{
		Videos:   videos,
		Channels: channels,
	}

	// Marshal results to JSON.
	jsonData, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		log.Fatalf("Error marshaling JSON: %v", err)
	}

	// Print the JSON output.
	fmt.Println(string(jsonData))
}

// searchYouTube takes a YouTube service, a query string, and a max result count,
// and returns two maps containing video and channel IDs mapped to their titles.
func searchYouTube(service *youtube.Service, q string, max int64) (map[string]string, map[string]string, error) {
	call := service.Search.List([]string{"id", "snippet"}).
		Q(q).
		MaxResults(max)

	response, err := call.Do()
	if err != nil {
		return nil, nil, err
	}

	// Group video and channel results in separate maps.
	videos := make(map[string]string)
	channels := make(map[string]string)

	// Iterate through each item and add it to the correct map.
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
