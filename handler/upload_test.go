package handler

import (
	"database/sql"
	"net/http"
	"testing"

	"github.com/dmitrymomot/media-srv/repository"
	"github.com/dmitrymomot/media-srv/storage"
)

func TestHandler_Upload(t *testing.T) {
	type fields struct {
		db      *sql.DB
		query   *repository.Queries
		storage *storage.Interactor
		resize  resizerFunc
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				db:      tt.fields.db,
				query:   tt.fields.query,
				storage: tt.fields.storage,
				resize:  tt.fields.resize,
			}
			if err := h.Upload(tt.args.w, tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("Handler.Upload() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
