// Code generated by sqlc. DO NOT EDIT.
// source: original_item.sql

package repository

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestQueries_CreateOriginalItem(t *testing.T) {
	createdAt, _ := time.Parse("Mon Jan 2 15:04:05 -0700 MST 2006", "Mon Jan 2 15:04:05 -0700 MST 2006")
	arg := CreateOriginalItemParams{
		ID:   uuid.New(),
		Name: "image.png",
		Path: "uploads/image.png",
		URL:  "http://test/uploads/image.png",
	}
	item := OriginalItem{
		ID:        arg.ID,
		Name:      arg.Name,
		Path:      arg.Path,
		URL:       arg.URL,
		CreatedAt: createdAt,
	}

	db, mock, err := NewSQLMock()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	CreateOriginalItemMock(mock, arg, nil)

	t.Run("success: create original", func(t *testing.T) {
		q := &Queries{
			db: db,
		}
		got, err := q.CreateOriginalItem(context.TODO(), arg)
		if err != nil {
			t.Errorf("Queries.CreateOriginalItem() error = %v", err)
			return
		}
		if !reflect.DeepEqual(got, item) {
			t.Errorf("Queries.CreateOriginalItem() = %v, want %v", got, item)
		}
	})

	wantErr := errors.New("CreateOriginalItem")
	CreateOriginalItemMock(mock, arg, wantErr)
	t.Run("error: create original", func(t *testing.T) {
		q := &Queries{
			db: db,
		}
		_, err := q.CreateOriginalItem(context.TODO(), arg)
		if err == nil {
			t.Errorf("Queries.CreateOriginalItem() error = %v, wanted = %v", err, wantErr)
			return
		}
	})
}

func TestQueries_GetOriginalItemByID(t *testing.T) {
	createdAt, _ := time.Parse("Mon Jan 2 15:04:05 -0700 MST 2006", "Mon Jan 2 15:04:05 -0700 MST 2006")
	item := OriginalItem{
		ID:        uuid.New(),
		Name:      "image.png",
		Path:      "uploads/image.png",
		URL:       "http://test/uploads/image.png",
		CreatedAt: createdAt,
	}
	db, mock, err := NewSQLMock()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	GetOriginalItemByIDMock(mock, item, nil)

	t.Run("success: get origin by id", func(t *testing.T) {
		q := &Queries{
			db: db,
		}
		got, err := q.GetOriginalItemByID(context.TODO(), item.ID)
		if err != nil {
			t.Errorf("Queries.GetOriginalItemByID() error = %v", err)
			return
		}
		if !reflect.DeepEqual(got, item) {
			t.Errorf("Queries.GetOriginalItemByID() = %v, want %v", got, item)
		}
	})

	expErr := errors.New("GetOriginalItemByID")
	GetOriginalItemByIDMock(mock, item, expErr)

	t.Run("error: get origin by id", func(t *testing.T) {
		q := &Queries{
			db: db,
		}
		_, err := q.GetOriginalItemByID(context.TODO(), item.ID)
		if err == nil {
			t.Errorf("Queries.GetOriginalItemByID() error = %v, wantErr %v", err, expErr)
			return
		}
	})
}

func TestQueries_GetOriginalItemsList(t *testing.T) {
	createdAt, _ := time.Parse("Mon Jan 2 15:04:05 -0700 MST 2006", "Mon Jan 2 15:04:05 -0700 MST 2006")
	arg := GetOriginalItemsListParams{
		Limit:  10,
		Offset: 0,
	}
	items := []OriginalItem{
		{
			ID:        uuid.New(),
			Name:      "image.png",
			Path:      "uploads/image.png",
			URL:       "http://test/uploads/image.png",
			CreatedAt: createdAt,
		},
		{
			ID:        uuid.New(),
			Name:      "image.png",
			Path:      "uploads/image.png",
			URL:       "http://test/uploads/image.png",
			CreatedAt: createdAt,
		},
	}
	db, mock, err := NewSQLMock()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	GetOriginalItemsListMock(mock, arg, items, nil)

	type fields struct {
		db DBTX
	}
	type args struct {
		ctx context.Context
		arg GetOriginalItemsListParams
	}
	tt := struct {
		name    string
		fields  fields
		args    args
		want    []OriginalItem
		wantErr bool
	}{"success: get origin list", fields{db}, args{context.TODO(), arg}, items, false}

	t.Run(tt.name, func(t *testing.T) {
		q := &Queries{
			db: tt.fields.db,
		}
		got, err := q.GetOriginalItemsList(tt.args.ctx, tt.args.arg)
		if (err != nil) != tt.wantErr {
			t.Errorf("Queries.GetOriginalItemsList() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("Queries.GetOriginalItemsList() = %v, want %v", got, tt.want)
		}
	})

	expErr := errors.New("GetOriginalItemsList")
	GetOriginalItemsListMock(mock, arg, items, expErr)

	tt = struct {
		name    string
		fields  fields
		args    args
		want    []OriginalItem
		wantErr bool
	}{"error: get origin list", fields{db}, args{context.TODO(), arg}, nil, true}

	t.Run(tt.name, func(t *testing.T) {
		q := &Queries{
			db: tt.fields.db,
		}
		got, err := q.GetOriginalItemsList(tt.args.ctx, tt.args.arg)
		if (err != nil) != tt.wantErr {
			t.Errorf("Queries.GetOriginalItemsList() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if got != nil {
			t.Errorf("Queries.GetOriginalItemsList() = %v, want %v", got, tt.want)
		}
	})
}
