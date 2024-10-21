package mrg

import (
	"html/template"
	"net/http"
)

func Search(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/Search.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.ServeFile(w, r, "templates/500.html")
		return
	}
	tmpl.Execute(w, nil)
}
