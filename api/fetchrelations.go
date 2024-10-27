package api

import (
	"encoding/json"
	"fmt"
	"tracker/models"

	"io"
	"net/http"
)

// FetchRelations retrieves relations data from the API.
func FetchRelations() (models.RelationsResponse, error) {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/relation")
	if err != nil {
		return models.RelationsResponse{}, fmt.Errorf("failed to fetch relations: %w", err)
	}
	defer func() {
		if closeErr := resp.Body.Close(); closeErr != nil {
			fmt.Println("Error closing response body:", closeErr)
		}
	}()

	body, readErr := io.ReadAll(resp.Body)
	if readErr != nil {
		return models.RelationsResponse{}, fmt.Errorf("failed to read response body: %w", readErr)
	}

	var relations models.RelationsResponse
	if jsonErr := json.Unmarshal(body, &relations); jsonErr != nil {
		return models.RelationsResponse{}, fmt.Errorf("failed to decode relations: %w", jsonErr)
	}

	return relations, nil
}
