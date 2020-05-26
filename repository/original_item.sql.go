// Code generated by sqlc. DO NOT EDIT.
// source: original_item.sql

package repository

import (
	"context"

	"github.com/google/uuid"
)

const createOriginalItem = `-- name: CreateOriginalItem :one
INSERT INTO original_items (id, name, path, url)
VALUES ($1, $2, $3, $4)
RETURNING id, name, path, url, created_at
`

type CreateOriginalItemParams struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
	Path string    `json:"path"`
	URL  string    `json:"url"`
}

func (q *Queries) CreateOriginalItem(ctx context.Context, arg CreateOriginalItemParams) (OriginalItem, error) {
	row := q.db.QueryRowContext(ctx, createOriginalItem,
		arg.ID,
		arg.Name,
		arg.Path,
		arg.URL,
	)
	var i OriginalItem
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Path,
		&i.URL,
		&i.CreatedAt,
	)
	return i, err
}

const getOriginalItemByID = `-- name: GetOriginalItemByID :one
SELECT id, name, path, url, created_at FROM original_items
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetOriginalItemByID(ctx context.Context, id uuid.UUID) (OriginalItem, error) {
	row := q.db.QueryRowContext(ctx, getOriginalItemByID, id)
	var i OriginalItem
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Path,
		&i.URL,
		&i.CreatedAt,
	)
	return i, err
}

const getOriginalItemsList = `-- name: GetOriginalItemsList :many
SELECT id, name, path, url, created_at FROM original_items
LIMIT $1 OFFSET $2
`

type GetOriginalItemsListParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetOriginalItemsList(ctx context.Context, arg GetOriginalItemsListParams) ([]OriginalItem, error) {
	rows, err := q.db.QueryContext(ctx, getOriginalItemsList, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []OriginalItem
	for rows.Next() {
		var i OriginalItem
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Path,
			&i.URL,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}