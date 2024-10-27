package api

import (
	"encoding/json"
	"fmt"
	"tracker/models"

	"io"
	"net/http"
)

// FetchLocations retrieves location data from the API.
func FetchLocations() ([]models.Location, error) {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/locations")
	if err != nil {
		return nil, fmt.Errorf("failed to fetch locations: %w", err)
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

	var wrapper models.LocationsResponse
	if jsonErr := json.Unmarshal(body, &wrapper); jsonErr != nil {
		return nil, fmt.Errorf("failed to decode locations: %w", jsonErr)
	}

	return wrapper.Index, nil
}
