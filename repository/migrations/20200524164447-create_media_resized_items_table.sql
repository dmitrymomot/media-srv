-- +migrate Up
CREATE TABLE IF NOT EXISTS resized_items (
    id uuid PRIMARY KEY,
    oid uuid NOT NULL,
    name varchar NOT NULL,
    path varchar NOT NULL,
    url varchar NOT NULL,
    height int NOT NULL,
    width int NOT NULL,
    created_at timestamp NOT NULL DEFAULT now()
);
CREATE INDEX resized_items_original ON resized_items USING BTREE (oid);
CREATE INDEX resized_items_created_at ON resized_items USING BTREE (created_at);


-- +migrate Down
DROP TABLE IF EXISTS resized_items;