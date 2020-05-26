package handler

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/go-chi/render"
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
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, render.M{"error": err.Error()})
		}
	}
}
