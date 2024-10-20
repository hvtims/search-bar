package mrg

// import "net/http"

// func Search(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != "GET" {
// 		w.WriteHeader(http.StatusBadRequest)
// 		http.ServeFile(w, r, "templates/405.html")
// 	}
// 	apiArtist := "https://groupietrackers.herokuapp.com/api/artists"

// 	artists, err := fetchLocation(apiArtist)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		http.ServeFile(w, r, "templates/500.html")
// 		return
// 	}
// }
