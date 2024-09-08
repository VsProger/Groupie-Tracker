package handlers

import (
	"log"
	"net/http"
	"strings"
)

type ErrorData struct {
	ErrorCode    int
	ErrorTitle   string
	ErrorMessage string
}

func displayError(w http.ResponseWriter, errTitle string, errMsg string, errCode int) {
	data := ErrorData{
		ErrorTitle:   errTitle,
		ErrorMessage: errMsg,
		ErrorCode:    errCode,
	}
	w.WriteHeader(errCode)
	err := tpl.ExecuteTemplate(w, "error.html", data)
	if err != nil {
		log.Printf("template execution error: %v", err)
	}
}

func formatLocation(location string) string {

	parts := strings.Split(location, "-")

	for i, part := range parts {
		part = strings.ReplaceAll(part, "_", " ")

		if strings.ToLower(part) == "usa" || strings.ToLower(part) == "uk" {
			parts[i] = strings.ToUpper(part)
		} else {
			words := strings.Fields(part)
			for j, word := range words {
				words[j] = strings.Title(word)
			}
			parts[i] = strings.Join(words, " ")
		}
	}
	return strings.Join(parts, "-")
}
