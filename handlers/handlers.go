package handlers

import (
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		displayError(w, "Method Not Allowed", "The method is not allowed for the requested URL.", http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path != "/" {
		displayError(w, "Page Not Found", "The page you are looking for might have been removed had its name changed or is temporarily unavailable.", http.StatusNotFound)
		return
	}
	artists, err := FetchArtists()
	if err != nil {
		displayError(w, "Internal Server Error", "There was an error fetching the artists.", http.StatusInternalServerError)
		return
	}
	err = tpl.ExecuteTemplate(w, "home.html", artists)
	// if err != nil {
	// 	displayError(w, "Internal Server Error", "There was an error rendering the home page.", http.StatusInternalServerError)
	// 	return
	// }
}
