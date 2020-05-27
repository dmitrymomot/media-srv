package handler

import (
	"net/http"
	"strconv"

	"github.com/dmitrymomot/media-srv/repository"
	"github.com/dmitrymomot/media-srv/storage"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/thedevsaddam/govalidator"
)

// Resize an uploaded item according to passed parameters
func (h *Handler) Resize(w http.ResponseWriter, r *http.Request) error {
	rules := govalidator.MapData{
		"height": []string{"required", "numeric", "numeric_between:32,512"},
		"width":  []string{"required", "numeric", "numeric_between:32,512"},
	}
	if err := validate(r, rules, nil); err != nil {
		return err
	}

	oid, err := uuid.Parse(chi.URLParam(r, "oid"))
	if err != nil {
		return NewHTTPError(http.StatusBadRequest, "wrong id format")
	}

	tx, err := h.db.Begin()
	if err != nil {
		return errors.Wrap(err, "begin transaction")
	}

	query := h.query.WithTx(tx)

	originalItem, err := query.GetOriginalItemByID(r.Context(), oid)
	if err != nil {
		tx.Rollback()
		return errors.Wrap(err, "get original item from db")
	}

	height, _ := strconv.Atoi(r.FormValue("height"))
	width, _ := strconv.Atoi(r.FormValue("width"))

	rid := uuid.New()
	resizedItem, err := query.CreateResizedItem(r.Context(), repository.CreateResizedItemParams{
		ID:     rid,
		OID:    oid,
		Name:   originalItem.Name,
		Path:   h.storage.FilePath(rid.String()),
		URL:    h.storage.FileURL(h.storage.FilePath(rid.String())),
		Width:  int32(width),
		Height: int32(height),
	})
	if err != nil {
		tx.Rollback()
		return errors.Wrap(err, "store resized item to db")
	}

	file, ct, err := h.storage.Download(originalItem.Path)
	if err != nil {
		tx.Rollback()
		return errors.Wrap(err, "download original image")
	}

	resizedFile, err := h.resize(file, width, height)
	if err != nil {
		tx.Rollback()
		return errors.Wrap(err, "resize image")
	}

	if err := h.storage.Upload(resizedFile, resizedItem.Path, storage.Public, *ct); err != nil {
		tx.Rollback()
		return errors.Wrap(err, "upload resized image")
	}

	if err := tx.Commit(); err != nil {
		return errors.Wrap(err, "commit resized item")
	}

	return jsonResponse(w, http.StatusOK, data{
		"original": originalItem,
		"resized":  resizedItem,
	})
}
