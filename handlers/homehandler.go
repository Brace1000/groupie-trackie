package handlers

import (
	"errors"
	"html/template"
	"net/http"

	"tracker/api"
)

var tmplt *template.Template

func init() {
	var err error
	tmplt, err = template.ParseGlob("./templates/*.html")
	if err != nil {
		panic("failed to parse templates: " + err.Error())
	}
}

// HomeHandler handles the home page requests.
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if !isValidPath(r.URL.Path) {
		err := errors.New("unsupported URL path") // Define an appropriate error
		HandleError(w, err, http.StatusNotFound, "404.html")
		return
	}

	artists, err := api.FetchArtists()
	if err != nil {
		HandleError(w, err, http.StatusInternalServerError, "500.html")
		return
	}

	// Execute the template with data
	if execErr := tmplt.ExecuteTemplate(w, "index.html", artists); execErr != nil {
		HandleError(w, execErr, http.StatusInternalServerError, "500.html")
		return
	}
}

// ArtistHandler handles requests for artist-specific information.
func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	if !isValidPath(r.URL.Path) {
		err := http.ErrNotSupported // Define an appropriate error
		HandleError(w, err, http.StatusBadRequest ,"400.html")
		return
	}

	artists, err := api.FetchArtists()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Execute the template with data
	if execErr := tmplt.ExecuteTemplate(w, "artist.html", artists); execErr != nil {
		HandleError(w, execErr, http.StatusInternalServerError, "500.html")
		return
	}
}

// isValidPath checks if the requested URL path is valid.
func isValidPath(path string) bool {
	return path == "/" || path == "/dates" || path == "/artist" || path == "/locations" || path == "/relations"
}
