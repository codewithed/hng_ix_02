// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: persons.sql

package db

import (
	"context"
)

const createPerson = `-- name: CreatePerson :one
INSERT INTO persons (
    name
) VALUES ($1)
RETURNING id, name, created_at
`

func (q *Queries) CreatePerson(ctx context.Context, name string) (Person, error) {
	row := q.db.QueryRowContext(ctx, createPerson, name)
	var i Person
	err := row.Scan(&i.ID, &i.Name, &i.CreatedAt)
	return i, err
}

const deletePerson = `-- name: DeletePerson :exec
DELETE FROM persons WHERE name = $1
`

func (q *Queries) DeletePerson(ctx context.Context, name string) error {
	_, err := q.db.ExecContext(ctx, deletePerson, name)
	return err
}

const getPerson = `-- name: GetPerson :one
SELECT id, name, created_at FROM persons
WHERE name = $1 LIMIT 1
`

func (q *Queries) GetPerson(ctx context.Context, name string) (Person, error) {
	row := q.db.QueryRowContext(ctx, getPerson, name)
	var i Person
	err := row.Scan(&i.ID, &i.Name, &i.CreatedAt)
	return i, err
}

const updatePerson = `-- name: UpdatePerson :one
UPDATE persons SET name = $1
WHERE name = $2
RETURNING id, name, created_at
`

type UpdatePersonParams struct {
	Name   string `json:"name"`
	Name_2 string `json:"name_2"`
}

func (q *Queries) UpdatePerson(ctx context.Context, arg UpdatePersonParams) (Person, error) {
	row := q.db.QueryRowContext(ctx, updatePerson, arg.Name, arg.Name_2)
	var i Person
	err := row.Scan(&i.ID, &i.Name, &i.CreatedAt)
	return i, err
}
