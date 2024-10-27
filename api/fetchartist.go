package api

import (
	"encoding/json"
	"fmt"
	"tracker/models"

	"io"
	"net/http"
)

// FetchArtists retrieves the list of artists from the API.
func FetchArtists() ([]models.Artist, error) {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		return nil, fmt.Errorf("failed to fetch artists: %w", err)
	}
	defer func() {
		if closeErr := resp.Body.Close(); closeErr != nil {
			fmt.Println("Error closing response body:", closeErr)
		}
	}()

	body, readErr := io.ReadAll(resp.Body)
	if readErr != nil {
		return nil, fmt.Errorf("failed to read response body: %w", readErr)
	}

	var artists []models.Artist
	if jsonErr := json.Unmarshal(body, &artists); jsonErr != nil {
		return nil, fmt.Errorf("failed to decode artists: %w", jsonErr)
	}
	return artists, nil
}
