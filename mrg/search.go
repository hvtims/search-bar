package mrg

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

func Search(w http.ResponseWriter, r *http.Request) {
	input := r.FormValue("input")
	if len(input) == 0 || len(input) >= 41 {
		w.WriteHeader(http.StatusBadRequest)
		http.ServeFile(w, r, "templates/400.html")
		return
	}
	search := strings.ToLower(input)

	apiURL := "https://groupietrackers.herokuapp.com/api/artists"
	artists, err := fetchArtists(apiURL)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.ServeFile(w, r, "templates/500.html")
		return
	}

	var searchResults Artists
	for _, artist := range artists {
		if strings.Contains(strings.ToLower(artist.Name), search) ||
			contains(artist.Members, search) ||
			strings.Contains(strings.ToLower(artist.FirstAlbum), search) ||
			fmt.Sprintf("%d", artist.CreationDate) == search {

			artist, err = fetchRelations(artist)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				http.ServeFile(w, r, "templates/500.html")
				return
			}
			artist, err = fetchLocation(artist)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				http.ServeFile(w, r, "templates/500.html")
				return
			}
			artist, err = fetchDates(artist)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				http.ServeFile(w, r, "templates/500.html")
				return
			}
			searchResults = append(searchResults, artist)
		}
	}
	println(len(searchResults))
	if len(searchResults) == 0{
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
		if strings.Contains(strings.ToLower(item), search) {
			return true
		}
	}
	return false
}
