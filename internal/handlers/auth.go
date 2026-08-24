package handlers

import (
	"net/http"

	"rqpt/hadida/internal/templates"
)

func (h *Handler) showRegister() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		component := templates.Register(templates.RegisterProps{})
		component.Render(r.Context(), w)
	}
}
