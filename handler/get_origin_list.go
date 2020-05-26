package handler

import (
	"net/http"
	"strconv"

	"github.com/dmitrymomot/media-srv/repository"
)

// GetOriginalItemsList http handler
func (h *Handler) GetOriginalItemsList(w http.ResponseWriter, r *http.Request) error {
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	if limit <= 0 {
		limit = 10
	}
	offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))
	items, err := h.query.GetOriginalItemsList(r.Context(), repository.GetOriginalItemsListParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	})
	if err != nil {
		return err
	}

	return jsonResponse(w, http.StatusOK, items)
}
