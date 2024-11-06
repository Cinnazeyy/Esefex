package routes

import (
	"html/template"
	"log"
	"net/http"
)

type LinkDeferData struct {
	LinkToken string
}

// link?<guild_id>
func (h *RouteHandlers) GetLinkDefer() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		linkToken := r.URL.Query().Get("t")

		if linkToken == "" {
			http.Error(w, "Missing link token", http.StatusBadRequest)
			return
		}

		data := LinkDeferData{
			LinkToken: linkToken,
		}

		tmpl, err := template.ParseFiles("./ressource/templates/link.html")
		if err != nil {
			log.Println(err)
			http.Error(w, "Error parsing template", http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Error executing template", http.StatusInternalServerError)
			return
		}
	})
}
