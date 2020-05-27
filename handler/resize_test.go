package handler

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/dmitrymomot/media-srv/repository"
	"github.com/dmitrymomot/media-srv/resizer"
	"github.com/dmitrymomot/media-srv/storage"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestHandler_Resize(t *testing.T) {
	createdAt, _ := time.Parse("Mon Jan 2 15:04:05 -0700 MST 2006", "Mon Jan 2 15:04:05 -0700 MST 2006")
	oid := uuid.New()
	item := repository.OriginalItem{
		ID:        oid,
		Name:      "image.png",
		Path:      "uploads/image.png",
		URL:       "http://test/uploads/image.png",
		CreatedAt: createdAt,
	}
	rid := uuid.New()
	ritem := repository.ResizedItem{
		ID:        rid,
		OID:       oid,
		Name:      "image.png",
		Path:      "uploads/image-100x100.png",
		URL:       "http://test/uploads/image-100x100.png",
		Width:     100,
		Height:    100,
		CreatedAt: createdAt,
	}
	db, mock, err := repository.NewSQLMock()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	mock.ExpectBegin()
	repository.GetOriginalItemByIDMock(mock, item, nil)
	repository.CreateResizedItemMock(mock, repository.CreateResizedItemParams{
		ID:     rid,
		OID:    oid,
		Name:   ritem.Name,
		Path:   ritem.Path,
		URL:    ritem.URL,
		Width:  ritem.Width,
		Height: ritem.Height,
	}, nil)
	mock.ExpectCommit()

	repo := repository.New(db)

	s3mock := &storage.S3Mock{Filepath: "./testdata/image.png", ContentType: "image/png"}
	opt := storage.Options{
		Bucket:         "test",
		URL:            "http://test.storage",
		ForcePathStyle: false,
	}
	stor := storage.New(s3mock, opt)

	data := url.Values{}
	data.Add("width", "100")
	data.Add("height", "100")
	r, err := http.NewRequest("POST", "/origin/{oid}/resized", strings.NewReader(data.Encode()))
	if err != nil {
		t.Fatal(err)
	}
	r.Form = data
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("oid", item.ID.String())
	r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))
	w := httptest.NewRecorder()

	t.Run("success", func(t *testing.T) {
		h := &Handler{
			db:      db,
			query:   repo,
			storage: stor,
			resize:  resizer.MockResize(nil),
		}
		if err := h.Resize(w, r); err != nil {
			t.Errorf("Handler.Upload() error = %v", err)
		}

		result := w.Result()
		body, _ := ioutil.ReadAll(result.Body)

		assert.Equal(t, http.StatusOK, result.StatusCode)

		expected, _ := json.Marshal(map[string]interface{}{
			"original": item,
			"resized":  ritem,
		})
		assert.JSONEqf(t, string(expected), string(body), "response does not match to expected jsonn string")
	})
}
