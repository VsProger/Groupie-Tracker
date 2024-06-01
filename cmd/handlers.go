package cmd

import (
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		displayError(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path != "/" {
		displayError(w, "Page Not Found", http.StatusNotFound)
	}
	artists, err := FetchArtists()
	if err != nil {
		displayError(w, "Internal Server Error", http.StatusInternalServerError)
	}
	tpl.ExecuteTemplate(w, "home.html", artists)
}

func displayError(w http.ResponseWriter, errMsg string, errCode int) {
	data := map[string]interface{}{
		"ErrorMessage": errMsg,
		"ErrorCode":    errCode,
	}
	w.WriteHeader(errCode)
	err := tpl.ExecuteTemplate(w, "error.html", data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
