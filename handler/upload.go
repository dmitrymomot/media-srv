package handler

import (
	"net/http"
	"strconv"

	"github.com/dmitrymomot/media-srv/repository"
	"github.com/dmitrymomot/media-srv/storage"
	"github.com/google/uuid"
	"github.com/thedevsaddam/govalidator"
)

// Upload new item and resize according to passed parameters
func (h *Handler) Upload(w http.ResponseWriter, r *http.Request) error {
	rules := govalidator.MapData{
		"file:image": []string{"required", "ext:png", "size:2097152", "mime:image/png"},
		"height":     []string{"required", "numeric", "numeric_between:32,512"},
		"width":      []string{"required", "numeric", "numeric_between:32,512"},
	}
	if err := validate(r, rules, nil); err != nil {
		return err
	}

	height, _ := strconv.Atoi(r.FormValue("height"))
	width, _ := strconv.Atoi(r.FormValue("width"))
	file, header, err := r.FormFile("image")
	defer file.Close()

	// Store original image
	tx, err := h.db.Begin()
	if err != nil {
		return err
	}
	query := h.query.WithTx(tx)

	oid := uuid.New()
	originalItem, err := query.CreateOriginalItem(r.Context(), repository.CreateOriginalItemParams{
		ID:   oid,
		Name: header.Filename,
		Path: h.storage.FilePath(oid.String()),
		URL:  h.storage.FileURL(h.storage.FilePath(oid.String())),
	})
	if err != nil {
		tx.Rollback()
		return err
	}
	if err := h.storage.Upload(file, originalItem.Path, storage.Public); err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	// Store resized image
	tx, err = h.db.Begin()
	if err != nil {
		return err
	}

	query = h.query.WithTx(tx)

	rid := uuid.New()
	resizedItem, err := query.CreateResizedItem(r.Context(), repository.CreateResizedItemParams{
		ID:     rid,
		OID:    oid,
		Name:   header.Filename,
		Path:   h.storage.FilePath(rid.String()),
		URL:    h.storage.FileURL(h.storage.FilePath(rid.String())),
		Width:  int32(width),
		Height: int32(height),
	})
	if err != nil {
		tx.Rollback()
		return err
	}

	resizedFile, err := h.resize(file, width, height)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err := h.storage.Upload(resizedFile, resizedItem.Path, storage.Public); err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return jsonResponse(w, http.StatusOK, data{
		"original": originalItem,
		"resized":  resizedItem,
	})
}
