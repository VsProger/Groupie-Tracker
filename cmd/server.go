package cmd

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func StartServer() {
	staticHandler := http.StripPrefix("/static/", http.FileServer(http.Dir("./static")))
	http.Handle("/static/", staticHandler)
	http.HandleFunc("/", Home)
	// http.HandleFunc("/artist", Artist)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
