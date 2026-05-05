package main

import (
	"fmt"
	"net/http"

	"github.com/common-nighthawk/go-figure"
)

func handleIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/index.html")
}

func handleGenerate(w http.ResponseWriter, r *http.Request) {
	text := r.URL.Query().Get("text")
	font := r.URL.Query().Get("font")

	if font == "" {
		font = "standard"
	}

	art := figure.NewFigure(text, font, true)

	w.Header().Set("Content-type", "text/plain")
	fmt.Fprint(w, art.String())

}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handleIndex)
	mux.HandleFunc("/generate", handleGenerate)
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", mux)
}
