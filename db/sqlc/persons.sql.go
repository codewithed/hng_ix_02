// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0
// source: persons.sql

package db

import (
	"context"
)

const createPerson = `-- name: CreatePerson :one
INSERT INTO persons (
    name, bio
) VALUES ($1, $2)
RETURNING id, name, bio, created_at
`

type CreatePersonParams struct {
	Name string `json:"name"`
	Bio  string `json:"bio"`
}

func (q *Queries) CreatePerson(ctx context.Context, arg CreatePersonParams) (Person, error) {
	row := q.db.QueryRowContext(ctx, createPerson, arg.Name, arg.Bio)
	var i Person
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Bio,
		&i.CreatedAt,
	)
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
SELECT id, name, bio, created_at FROM persons
WHERE name = $1 LIMIT 1
`

func (q *Queries) GetPerson(ctx context.Context, name string) (Person, error) {
	row := q.db.QueryRowContext(ctx, getPerson, name)
	var i Person
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Bio,
		&i.CreatedAt,
	)
	return i, err
}

const updatePerson = `-- name: UpdatePerson :one
UPDATE persons SET bio = $1
WHERE name = $2
RETURNING id, name, bio, created_at
`

type UpdatePersonParams struct {
	Bio  string `json:"bio"`
	Name string `json:"name"`
}

func (q *Queries) UpdatePerson(ctx context.Context, arg UpdatePersonParams) (Person, error) {
	row := q.db.QueryRowContext(ctx, updatePerson, arg.Bio, arg.Name)
	var i Person
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Bio,
		&i.CreatedAt,
	)
	return i, err
}
