// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1

package db

import (
	"context"
)

type Querier interface {
	CreatePerson(ctx context.Context, name string) (Person, error)
	DeletePerson(ctx context.Context, name string) error
	GetPerson(ctx context.Context, name string) (Person, error)
	UpdatePerson(ctx context.Context, arg UpdatePersonParams) (Person, error)
}

var _ Querier = (*Queries)(nil)
