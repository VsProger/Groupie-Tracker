package handlers

import (
	"log"
	"net/http"
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
