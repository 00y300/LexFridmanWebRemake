package settings

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// GetAPIKey loads environment variables from a .env file and returns the GOOGLE_API key.
// It returns an error if the .env file cannot be loaded or if the key is not set.
func GetAPIKey() (string, error) {
	// Load environment variables from the .env file.
	if err := godotenv.Load(".env"); err != nil {
		return "", fmt.Errorf("error loading .env file: %v", err)
	}

	// Retrieve the YouTube API key from the environment.
	apiKey := os.Getenv("GOOGLE_API")
	if apiKey == "" {
		return "", fmt.Errorf("GOOGLE_API key is not set in the .env file")
	}

	return apiKey, nil
}
