package main

import (
	"fmt"
	"net/http"
	"os"

	"mrg/mrg"
)

func main() {
	if len(os.Args) != 1 {
		fmt.Println("check args")
		return
	}
	fmt.Println("this is your port : http://localhost:8080/ ")
	http.HandleFunc("/", mrg.HandleHome)
	http.HandleFunc("/Artist/{id}", mrg.HandlePage)
	http.HandleFunc("/404", mrg.ErrorHandler)
	http.HandleFunc("/frontend/css/", mrg.CssHandler)
	http.HandleFunc("/Search", mrg.Search)
	http.HandleFunc("/frontend/images/", mrg.ImageHandler)
	http.ListenAndServe(":8080", nil)
}
