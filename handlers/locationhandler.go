package handlers

import (
	"net/http"
	"strconv"
	"tracker/api"
	"tracker/models"
)

// LocationHandler handles requests for locations associated with a specific artist.
func LocationHandler(w http.ResponseWriter, r *http.Request) {
	artistIdStr := r.URL.Query().Get("artistId")
	artistId, err := strconv.Atoi(artistIdStr)
	if err != nil {
		HandleError(w, err, http.StatusBadRequest ,"400.html")
		return
	}

	locations, err := api.FetchLocations()
	if err != nil {
		HandleError(w, err, http.StatusInternalServerError, "500.html")
		return
	}

	filteredLocations := filterLocationsByArtistID(locations, artistId)

	if err := tmplt.ExecuteTemplate(w, "locations.html", filteredLocations); err != nil {
		HandleError(w, err, http.StatusInternalServerError, "500.html")
		return
	}
}

// filterLocationsByArtistID filters locations based on the artist ID.
func filterLocationsByArtistID(locations []models.Location, artistId int) []models.Location {
	var filtered []models.Location
	for _, location := range locations {
		if location.ID == artistId {
			filtered = append(filtered, location)
		}
	}
	return filtered
}
