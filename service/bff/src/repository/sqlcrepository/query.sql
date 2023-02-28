-- name: CreateCustomer :one
INSERT INTO customers (
  name
) VALUES (
  $1
)
RETURNING *;
