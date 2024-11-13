package handler

import (
	"fmt"
	"net/http"
	"text/template"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	templatePath := "templates/home.html"

	t, err := template.ParseFiles(templatePath)
	if err != nil {
		fmt.Printf("could not find or load template: %v", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	if err := t.Execute(w, map[string]interface{}{}); err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}
