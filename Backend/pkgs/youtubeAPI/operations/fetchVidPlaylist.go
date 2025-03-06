package operations

import (
	"Backend/pkgs/youtubeAPI/settings"
	"context"
	"fmt"

	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

// FetchResultsVideos is a struct made to store the video IDs of the given playlist.
type FetchResultsVideos struct {
	Videos map[string]string `json:"videos"`
	// Channels map[string]string `json:"channels"`
}

// FindVideosFromSourcePlaylist retrieves videos from a specific playlist that belongs to the given channel.
// It validates that the playlist belongs to the channel, then returns maps of video IDs to video titles,
// and video IDs to channel titles.
func FindVideosFromSourcePlaylist(channelID, playlistID string, limit int64) (map[string]string, map[string]string, error) {
	// Retrieve the YouTube API key from the environment using the config helper.
	apiKey, err := settings.GetAPIKey()
	if err != nil {
		return nil, nil, err
	}

	// Create a new YouTube service client.
	ctx := context.Background()
	service, err := youtube.NewService(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return nil, nil, fmt.Errorf("error creating YouTube client: %v", err)
	}

	// Validate that the provided playlist belongs to the given channel.
	playlistCall := service.Playlists.List([]string{"snippet"}).Id(playlistID)
	playlistResponse, err := playlistCall.Do()
	if err != nil {
		return nil, nil, fmt.Errorf("error fetching playlist details: %v", err)
	}
	if len(playlistResponse.Items) == 0 {
		return nil, nil, fmt.Errorf("no playlist found with ID %s", playlistID)
	}
	if playlistResponse.Items[0].Snippet.ChannelId != channelID {
		return nil, nil, fmt.Errorf("playlist %s does not belong to channel %s", playlistID, channelID)
	}

	// Fetch videos from the specified playlist.
	call := service.PlaylistItems.List([]string{"snippet"}).
		PlaylistId(playlistID).
		MaxResults(limit)

	response, err := call.Do()
	if err != nil {
		return nil, nil, fmt.Errorf("error fetching playlist items: %v", err)
	}

	// Initialize the maps to store video details.
	videos := make(map[string]string)
	channels := make(map[string]string)

	// Iterate over the playlist items and collect video information.
	for _, item := range response.Items {
		videoID := item.Snippet.ResourceId.VideoId
		title := item.Snippet.Title
		videos[videoID] = title
		// Optionally, record the channel title for each video.
		channels[videoID] = item.Snippet.ChannelTitle
	}

	return videos, channels, nil
}
