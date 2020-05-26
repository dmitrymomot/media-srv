package handler

import (
	"net/http"

	"github.com/go-chi/chi"
)

// Router returns http.Handler innterface
func Router(h *Handler) http.Handler {
	r := chi.NewRouter()
	return r
}
