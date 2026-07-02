package handlers

import (
	"errors"
	"html/template"
	"net/http"
	"os"

	"ps/ascii"
	"ps/models"
)

func ASCIIArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/ascii-art" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
		return
	}

	text := r.FormValue("text")
	banner := r.FormValue("banner")

	result, err := ascii.GenerateASCII(text, banner)

	if err != nil {
		switch {
		case errors.Is(err, ascii.ErrInvalidBanner):
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return

		case errors.Is(err, os.ErrNotExist):
			http.Error(w, "Banner Not Found", http.StatusNotFound)
			return

		default:
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}

	data := models.PageData{
		Result: result,
		Text:   text,
		Banner: banner,
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Template Not Found", http.StatusNotFound)
		return
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
