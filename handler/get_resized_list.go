package handler

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

// GetResizedItemsList http handler
func (h *Handler) GetResizedItemsList(w http.ResponseWriter, r *http.Request) error {
	oid, err := uuid.Parse(chi.URLParam(r, "oid"))
	if err != nil {
		return NewHTTPError(http.StatusBadRequest, "wrong id format")
	}

	items, err := h.query.GetResizedItemsList(r.Context(), oid)
	if err != nil {
		return err
	}

	return jsonResponse(w, http.StatusOK, items)
}
