package routes

import (
	"html/template"
	"log"
	"net/http"
)

type Home struct {
	Template *template.Template
}

func (h *Home) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := h.Template.ExecuteTemplate(w, "layout", nil)
	if err != nil {
		log.Printf("home: could not render template: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
