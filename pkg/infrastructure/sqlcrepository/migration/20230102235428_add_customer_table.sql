-- +goose Up
CREATE TABLE customers (
  id   BIGSERIAL PRIMARY KEY,
  name text      NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS customers;
