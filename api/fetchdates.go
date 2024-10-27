package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"tracker/models"
)

// FetchConcertDates retrieves concert dates from the API.
func FetchConcertDates() (models.DatesResponse, error) {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/dates")
	if err != nil {
		return models.DatesResponse{}, fmt.Errorf("failed to fetch concert dates: %w", err)
	}
	defer func() {
		if closeErr := resp.Body.Close(); closeErr != nil {
			fmt.Println("Error closing response body:", closeErr)
		}
	}()

	body, readErr := io.ReadAll(resp.Body)
	if readErr != nil {
		return models.DatesResponse{}, fmt.Errorf("failed to read response body: %w", readErr)
	}

	var concertDates models.DatesResponse
	if jsonErr := json.Unmarshal(body, &concertDates); jsonErr != nil {
		return models.DatesResponse{}, fmt.Errorf("failed to decode concert dates: %w", jsonErr)
	}
	return concertDates, nil
}
