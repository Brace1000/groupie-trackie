package handlers

import (
	"net/http"
	"strconv"
	"tracker/api"
	"tracker/models"
)

// RelationHandler handles requests for relations associated with a specific artist.
func RelationHandler(w http.ResponseWriter, r *http.Request) {
	artistIdStr := r.URL.Query().Get("artistId")
	artistId, err := strconv.Atoi(artistIdStr)
	if err != nil {
		HandleError(w, err, http.StatusBadRequest ,"400.html")
		return
	}

	relations, err := api.FetchRelations()
	if err != nil {
		HandleError(w, err, http.StatusInternalServerError, "500.html")
		return
	}

	filteredRelations := filterRelationsByArtistID(relations.Index, artistId)

	if err := tmplt.ExecuteTemplate(w, "relations.html", filteredRelations); err != nil {
		HandleError(w, err, http.StatusInternalServerError, "500.html")
		return
	}
}

// filterRelationsByArtistID filters relations based on the artist ID.
func filterRelationsByArtistID(relations []models.Relation, artistId int) []models.Relation {
	var filtered []models.Relation
	for _, relation := range relations {
		if relation.ID == artistId {
			filtered = append(filtered, relation)
		}
	}
	return filtered
}
