package mrg

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

var searchResults Artists

func Search(w http.ResponseWriter, r *http.Request) {
	input := r.FormValue("input")
	if len(input) == 0 || len(input) >= 41 {
		w.WriteHeader(http.StatusBadRequest)
		http.ServeFile(w, r, "templates/400.html")
		return
	}
	search := strings.ToLower(input)

	apiURL := "https://groupietrackers.herokuapp.com/api/artists"
	artists, _ := fetchArtists(apiURL)
	artist, _ = fetchRelations(artist)
	artist, _ = fetchLocation(artist)
	artist, _ = fetchDates(artist)

	for _, artist := range artists {
		if strings.Contains(strings.ToLower(artist.Name), search) ||
			contains(artist.Members, search) ||
			strings.Contains(strings.ToLower(artist.FirstAlbum), search) ||
			fmt.Sprintf("%d", artist.CreationDate) == search ||
			contains(artist.Location.Locations, search) {
			searchResults = append(searchResults, artist)
		}
	}

	println(len(searchResults))
	if len(searchResults) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		http.ServeFile(w, r, "templates/400.html")
		return
	}

	tmpl, err := template.ParseFiles("templates/Search.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.ServeFile(w, r, "templates/500.html")
		return
	}

	err = tmpl.Execute(w, searchResults)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.ServeFile(w, r, "templates/500.html")
		return
	}
}

func contains(slice []string, search string) bool {
	for _, item := range slice {
		if strings.Contains(strings.ToLower(item), strings.ToLower(search)) {
			return true
		}
	}
	return false
}
