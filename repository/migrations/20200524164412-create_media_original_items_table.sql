-- +migrate Up
CREATE TABLE IF NOT EXISTS original_items (
    id uuid PRIMARY KEY,
    name varchar NOT NULL,
    path varchar NOT NULL,
    url varchar NOT NULL,
    created_at timestamp NOT NULL DEFAULT now()
);
CREATE INDEX original_items_created_at ON original_items USING BTREE (created_at);


-- +migrate Down
DROP TABLE IF EXISTS original_items;