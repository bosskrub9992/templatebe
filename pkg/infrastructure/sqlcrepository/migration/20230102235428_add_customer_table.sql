-- +goose Up
-- +goose StatementBegin
CREATE TABLE customers (
  id   BIGSERIAL PRIMARY KEY,
  name text      NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS customers;
-- +goose StatementEnd
