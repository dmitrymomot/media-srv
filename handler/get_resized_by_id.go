package handler

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

// GetResizedItem http handler
func (h *Handler) GetResizedItem(w http.ResponseWriter, r *http.Request) error {
	rid, err := uuid.Parse(chi.URLParam(r, "rid"))
	if err != nil {
		return NewHTTPError(http.StatusBadRequest, "wrong id format")
	}
	item, err := h.query.GetResizedItemByID(r.Context(), rid)
	if err != nil {
		return err
	}

	return jsonResponse(w, http.StatusOK, item)
}
