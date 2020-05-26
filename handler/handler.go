package handler

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/dmitrymomot/media-srv/repository"
	"github.com/dmitrymomot/media-srv/storage"
	"github.com/go-chi/render"
	"github.com/pkg/errors"
)

type (
	// Wrap function makes a handler compatible with the default go mux interface
	Wrap func(w http.ResponseWriter, r *http.Request) error

	// ValidationError ...
	ValidationError url.Values

	// HTTPError struct
	HTTPError struct {
		Code    int
		Message interface{}
	}

	resizerFunc func(f io.ReadCloser, w, h int) (io.ReadSeeker, error)

	logger interface {
		Debug(v ...interface{})
		Info(v ...interface{})
		Warn(v ...interface{})
		Error(v ...interface{})
	}

	// Handler struct
	Handler struct {
		db      *sql.DB
		query   *repository.Queries
		storage *storage.Interactor
		resize  resizerFunc
		log     logger
	}
)

// Implementation of the error interface
func (e ValidationError) Error() string {
	return fmt.Sprintf("validationError: %v", e)
}

// Implementation of the error interface
func (e HTTPError) Error() string {
	return fmt.Sprintf("code: %d; message: %v", e.Code, e.Message)
}

// ServeHTTP func is http.Handler interface implementation
func (h Wrap) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h(w, r); err != nil {
		switch err.(type) {
		case *ValidationError:
			render.Status(r, http.StatusUnprocessableEntity)
			render.JSON(w, r, render.M{"validationError": err})
		case *HTTPError:
			er := err.(*HTTPError)
			render.Status(r, er.Code)
			render.JSON(w, r, render.M{"error": er.Message})
		default:
			log.Println(errors.Wrap(err, "undefined http error"))
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, render.M{"error": err.Error()})
		}
	}
}

// New is a factory function,
// returns a new instance of the HTTP handler interactor
func New(db *sql.DB, q *repository.Queries, s *storage.Interactor, rf resizerFunc, l logger) *Handler {
	return &Handler{
		db:      db,
		query:   q,
		storage: s,
		resize:  rf,
		log:     l,
	}
}
