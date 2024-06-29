package handlers

import (
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

var tpl *template.Template

func initTemplates(basePath string) {
	var err error
	tpl, err = template.ParseGlob(filepath.Join(basePath, "templates", "*"))
	if err != nil {
		panic(err)
	}
}

func StartServer() {
	basePath, _ := os.Getwd()
	initTemplates(basePath)

	staticHandler := http.StripPrefix("/static/", http.FileServer(http.Dir(filepath.Join(basePath, "static"))))
	http.Handle("/static/", staticHandler)
	http.HandleFunc("/", Home)
	http.HandleFunc("/artist/", ArtistHandler)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.ListenAndServe(":8080", nil)
}
