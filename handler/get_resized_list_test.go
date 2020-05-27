package handler

import (
	"context"
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/dmitrymomot/media-srv/repository"
	"github.com/dmitrymomot/media-srv/resizer"
	"github.com/dmitrymomot/media-srv/storage"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestHandler_GetResizedItemsList(t *testing.T) {
	createdAt, _ := time.Parse("Mon Jan 2 15:04:05 -0700 MST 2006", "Mon Jan 2 15:04:05 -0700 MST 2006")
	oid := uuid.New()
	items := []repository.ResizedItem{
		{
			ID:        uuid.New(),
			OID:       oid,
			Name:      "image.png",
			Path:      "uploads/image.png",
			URL:       "http://test/uploads/image.png",
			Width:     100,
			Height:    100,
			CreatedAt: createdAt,
		},
		{
			ID:        uuid.New(),
			OID:       oid,
			Name:      "image.png",
			Path:      "uploads/image.png",
			URL:       "http://test/uploads/image.png",
			Width:     100,
			Height:    100,
			CreatedAt: createdAt,
		},
	}
	db, mock, err := repository.NewSQLMock()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	repository.GetResizedItemsListMock(mock, oid, items, nil)

	repo := repository.New(db)

	s3mock := &storage.S3Mock{}
	opt := storage.Options{
		Bucket:         "test",
		URL:            "http://test.storage",
		ForcePathStyle: false,
	}
	stor := storage.New(s3mock, opt)

	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "/origin/{oid}/resized", nil)
	if err != nil {
		t.Fatal(err)
	}
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("oid", oid.String())
	r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))

	t.Run("success", func(t *testing.T) {
		h := &Handler{
			db:      db,
			query:   repo,
			storage: stor,
			resize:  resizer.Resize,
		}
		if err := h.GetResizedItemsList(w, r); err != nil {
			t.Errorf("Handler.GetResizedItemsList() error = %v", err)
		}

		resp := w.Result()
		body, _ := ioutil.ReadAll(resp.Body)

		assert.Equal(t, resp.StatusCode, http.StatusOK)

		expected, _ := json.Marshal(items)
		assert.JSONEqf(t, string(expected), string(body), "response does not match to expected jsonn string")
	})
}

func TestHandler_GetResizedItemsList_WrongID(t *testing.T) {
	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "/origin/{oid}/resized", nil)
	if err != nil {
		t.Fatal(err)
	}
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("oid", "123")
	r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))

	t.Run("wrong origin id", func(t *testing.T) {
		Wrap((&Handler{}).GetResizedItemsList).ServeHTTP(w, r)
		assert.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
	})
}

func TestHandler_GetResizedItemsList_NotFound(t *testing.T) {
	oid := uuid.New()
	db, mock, err := repository.NewSQLMock()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	repository.GetResizedItemsListMock(mock, oid, nil, sql.ErrNoRows)

	repo := repository.New(db)

	s3mock := &storage.S3Mock{}
	opt := storage.Options{
		Bucket:         "test",
		URL:            "http://test.storage",
		ForcePathStyle: false,
	}
	stor := storage.New(s3mock, opt)

	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "/origin/{oid}/resized", nil)
	if err != nil {
		t.Fatal(err)
	}
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("oid", oid.String())
	r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))

	t.Run("not found", func(t *testing.T) {
		h := &Handler{
			db:      db,
			query:   repo,
			storage: stor,
			resize:  resizer.Resize,
		}
		Wrap(h.GetResizedItemsList).ServeHTTP(w, r)
		assert.Equal(t, http.StatusNotFound, w.Result().StatusCode)
	})
}
