-- name: GetResizedItemByID :one
SELECT * FROM resized_items
WHERE id = $1
LIMIT 1;

-- name: GetResizedItemsList :many
SELECT * FROM resized_items
WHERE oid = $1
ORDER BY created_at DESC;

-- name: CreateResizedItem :one
INSERT INTO resized_items (id, oid, name, path, url, height, width)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;