package utils

import (
	"html/template"
	"log"
	"net/http"
)

var (
	tpl *template.Template
)

func init() {
	tpl, _ = template.ParseGlob("templates/*.html")
}

func RenderTemplate(w http.ResponseWriter, templateName string, data interface{}) {
	err := tpl.ExecuteTemplate(w, templateName, data)
	if err != nil {
		http.Error(w, "Failed to render template: "+err.Error(), http.StatusInternalServerError)
		log.Printf("Error rendering template: %s", err)
	}
}
