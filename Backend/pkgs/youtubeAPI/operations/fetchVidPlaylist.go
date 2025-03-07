package operations

import (
	"Backend/pkgs/youtubeAPI/settings"
	"context"
	"encoding/json"
	"fmt"

	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

// PlaylistVideos holds channel ID, playlist ID, the playlist name,
// and a map of all video IDs to their titles.
type PlaylistVideos struct {
	ChannelID     string            `json:"channelID"`
	PlaylistID    string            `json:"playlistID"`
	PlaylistName  string            `json:"playlistName"`
	YoutubeVideos map[string]string `json:"youtubeVideos"`
}

// FetchAllVideosFromPlaylist retrieves *all* videos from a given playlist,
// ensuring the playlist belongs to the given channel. It returns a struct
// containing the channel ID, playlist ID, playlist name, and a map of videos.
func FetchAllVideosFromPlaylist(channelID, playlistID string) (*PlaylistVideos, error) {
	// 1. Retrieve the YouTube API key
	apiKey, err := settings.GetAPIKey()
	if err != nil {
		return nil, fmt.Errorf("failed to get API key: %v", err)
	}

	// 2. Create a new YouTube service client.
	ctx := context.Background()
	service, err := youtube.NewService(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return nil, fmt.Errorf("error creating YouTube client: %v", err)
	}

	// 3. Validate that the playlist belongs to the given channel.
	playlistCall := service.Playlists.List([]string{"snippet"}).Id(playlistID)
	playlistResponse, err := playlistCall.Do()
	if err != nil {
		return nil, fmt.Errorf("error fetching playlist details: %v", err)
	}
	if len(playlistResponse.Items) == 0 {
		return nil, fmt.Errorf("no playlist found with ID %s", playlistID)
	}
	if playlistResponse.Items[0].Snippet.ChannelId != channelID {
		return nil, fmt.Errorf("playlist %s does not belong to channel %s", playlistID, channelID)
	}

	// Extract the name of the playlist.
	playlistName := playlistResponse.Items[0].Snippet.Title

	// 4. Paginate through all items in the playlist.
	allVideos := make(map[string]string)
	pageToken := ""

	for {
		// Each call can retrieve up to 50 items per page.
		call := service.PlaylistItems.List([]string{"snippet"}).
			PlaylistId(playlistID).
			MaxResults(50).
			PageToken(pageToken)

		response, err := call.Do()
		if err != nil {
			return nil, fmt.Errorf("error fetching playlist items: %v", err)
		}

		for _, item := range response.Items {
			videoID := item.Snippet.ResourceId.VideoId
			title := item.Snippet.Title
			allVideos[videoID] = title
		}

		// If there's no next page token, we've retrieved all videos.
		if response.NextPageToken == "" {
			break
		}
		pageToken = response.NextPageToken
	}

	// 5. Build and return our result struct.
	result := &PlaylistVideos{
		ChannelID:     channelID,
		PlaylistID:    playlistID,
		PlaylistName:  playlistName,
		YoutubeVideos: allVideos,
	}
	return result, nil
}

// FetchAllVideosAsJSON is a helper function that returns the entire
// result as a JSON string.
func FetchAllVideosAsJSON(channelID, playlistID string) (string, error) {
	data, err := FetchAllVideosFromPlaylist(channelID, playlistID)
	if err != nil {
		return "", err
	}

	// Convert the struct to JSON (pretty print for readability).
	jsonBytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal JSON: %v", err)
	}

	return string(jsonBytes), nil
}
