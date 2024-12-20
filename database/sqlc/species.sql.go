// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: species.sql

package database

import (
	"context"
)

const createSpecies = `-- name: CreateSpecies :one
INSERT INTO species (name)
VALUES ($1)
RETURNING id, name, created_at, updated_at
`

func (q *Queries) CreateSpecies(ctx context.Context, name string) (Species, error) {
	row := q.db.QueryRow(ctx, createSpecies, name)
	var i Species
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getSpecieByID = `-- name: GetSpecieByID :one
SELECT id, name, created_at, updated_at
FROM species
WHERE id = $1
`

func (q *Queries) GetSpecieByID(ctx context.Context, id int32) (Species, error) {
	row := q.db.QueryRow(ctx, getSpecieByID, id)
	var i Species
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getSpeciesPaginated = `-- name: GetSpeciesPaginated :many
SELECT id, name, created_at, updated_at
FROM species
ORDER BY name
LIMIT $1 OFFSET $2
`

type GetSpeciesPaginatedParams struct {
	Limit  int32 `db:"limit" json:"limit"`
	Offset int32 `db:"offset" json:"offset"`
}

func (q *Queries) GetSpeciesPaginated(ctx context.Context, arg GetSpeciesPaginatedParams) ([]Species, error) {
	rows, err := q.db.Query(ctx, getSpeciesPaginated, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Species
	for rows.Next() {
		var i Species
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateSpeciesByID = `-- name: UpdateSpeciesByID :one
UPDATE species
SET name = $1,
    updated_at = NOW()
WHERE id = $2
RETURNING id, name, created_at, updated_at
`

type UpdateSpeciesByIDParams struct {
	Name string `db:"name" json:"name"`
	ID   int32  `db:"id" json:"id"`
}

func (q *Queries) UpdateSpeciesByID(ctx context.Context, arg UpdateSpeciesByIDParams) (Species, error) {
	row := q.db.QueryRow(ctx, updateSpeciesByID, arg.Name, arg.ID)
	var i Species
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
