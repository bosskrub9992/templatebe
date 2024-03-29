// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: query.sql

package sqlc

import (
	"context"
)

const createCustomer = `-- name: CreateCustomer :one
INSERT INTO customers (
  name
) VALUES (
  $1
)
RETURNING id, name
`

func (q *Queries) CreateCustomer(ctx context.Context, name string) (Customer, error) {
	row := q.db.QueryRowContext(ctx, createCustomer, name)
	var i Customer
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}
