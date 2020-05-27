package handler

import (
	"fmt"
	"net/http"
	"path"
	"strconv"

	"github.com/dmitrymomot/media-srv/repository"
	"github.com/dmitrymomot/media-srv/storage"
	"github.com/google/uuid"
	"github.com/pkg/errors"
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
		return errors.Wrap(err, "begin db transaction")
	}
	query := h.query.WithTx(tx)

	oid := uuid.New()
	fname := fmt.Sprintf("%s%s", oid.String(), path.Ext(header.Filename))
	ct := header.Header.Get("Content-Type")
	originalItem, err := query.CreateOriginalItem(r.Context(), repository.CreateOriginalItemParams{
		ID:   oid,
		Name: header.Filename,
		Path: h.storage.FilePath(fname),
		URL:  h.storage.FileURL(h.storage.FilePath(fname)),
	})
	if err != nil {
		tx.Rollback()
		return errors.Wrap(err, "store original item to db")
	}
	if err := h.storage.Upload(file, originalItem.Path, storage.Public, ct); err != nil {
		tx.Rollback()
		return errors.Wrap(err, "upload original image")
	}

	if err := tx.Commit(); err != nil {
		return errors.Wrap(err, "commit original item")
	}

	// Store resized image
	tx, err = h.db.Begin()
	if err != nil {
		return errors.Wrap(err, "begin db transaction for resized item")
	}

	query = h.query.WithTx(tx)

	rid := uuid.New()
	rfname := fmt.Sprintf("%s%s", rid.String(), path.Ext(header.Filename))
	resizedItem, err := query.CreateResizedItem(r.Context(), repository.CreateResizedItemParams{
		ID:     rid,
		OID:    oid,
		Name:   header.Filename,
		Path:   h.storage.FilePath(rfname),
		URL:    h.storage.FileURL(h.storage.FilePath(rfname)),
		Width:  int32(width),
		Height: int32(height),
	})
	if err != nil {
		tx.Rollback()
		return errors.Wrap(err, "store resized item to db")
	}

	if _, err := file.Seek(0, 0); err != nil {
		tx.Rollback()
		return errors.Wrap(err, "return coursor to the begin of file")
	}

	resizedFile, err := h.resize(file, width, height)
	if err != nil {
		tx.Rollback()
		return errors.Wrap(err, "resize image")
	}

	if err := h.storage.Upload(resizedFile, resizedItem.Path, storage.Public, ct); err != nil {
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
