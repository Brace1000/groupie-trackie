package handlers

import (
	"net/http"
	"strconv"
	"tracker/api"
	"tracker/models"
)

func ArtistDetails(w http.ResponseWriter, r *http.Request) {

	artistIdStr := r.URL.Query().Get("artistId")
	artistId, err := strconv.Atoi(artistIdStr)
	if err != nil {
		HandleError(w, err, http.StatusBadRequest, "400.html")
		return
	}

	artists, err := api.FetchArtists()
	if err != nil {
		HandleError(w, err, http.StatusInternalServerError, "500.html")
		return
	}

	var selectedArtist models.Artist

	for _, artist := range artists {
		if artist.ID == artistId {
			selectedArtist = artist
			break
		}
	}

	if selectedArtist.ID == 0 {
		HandleError(w, err, http.StatusBadRequest, "400.html")
		return
	}

	// Fetch location data
	locations, err := api.FetchLocations()
	if err != nil {
		HandleError(w, err, http.StatusInternalServerError, "500.html")
		return
	}

	// Fetch concert dates
	concertDates, err := api.FetchConcertDates()
	if err != nil {
		HandleError(w, err, http.StatusInternalServerError, "500.html")
		return
	}

	// Fetch relations
	relations, err := api.FetchRelations()
	if err != nil {
		HandleError(w, err, http.StatusInternalServerError, "500.html")
		return
	}

	// Filter data by artist ID
	var artistLocations []models.Location
	for _, location := range locations {
		if location.ID == artistId {
			artistLocations = append(artistLocations, location)
		}
	}

	var artistConcertDates []models.Date
	for _, concertDate := range concertDates.Index {
		if concertDate.ID == artistId {
			artistConcertDates = append(artistConcertDates, concertDate)
		}
	}

	var artistRelations []models.Relation
	for _, relation := range relations.Index {
		if relation.ID == artistId {
			artistRelations = append(artistRelations, relation)
		}
	}

	// Combine artist data with related data
	data := struct {
		Artist       models.Artist
		Locations    []models.Location
		ConcertDates []models.Date
		Relations    []models.Relation
	}{
		Artist:       selectedArtist,
		Locations:    artistLocations,
		ConcertDates: artistConcertDates,
		Relations:    artistRelations,
	}

	// Execute the template with data
	err = tmplt.ExecuteTemplate(w, "artistdetails.html", data)
	if err != nil {
		HandleError(w, err, http.StatusInternalServerError, "500.html")
		return
	}
}
