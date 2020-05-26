-- name: GetOriginalItemByID :one
SELECT * FROM original_items
WHERE id = $1
LIMIT 1;

-- name: GetOriginalItemsList :many
SELECT * FROM original_items
LIMIT $1 OFFSET $2;

-- name: CreateOriginalItem :one
INSERT INTO original_items (id, name, path, url)
VALUES ($1, $2, $3, $4)
RETURNING *;