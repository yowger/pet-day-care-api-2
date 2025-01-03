// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: pets.sql

package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createPet = `-- name: CreatePet :one
INSERT INTO pets (name, birth_date, species_id, breed_id)
VALUES ($1, $2, $3, $4)
RETURNING id, birth_date, name, species_id, breed_id, created_at, updated_at
`

type CreatePetParams struct {
	Name      string      `db:"name" json:"name"`
	BirthDate pgtype.Date `db:"birth_date" json:"birth_date"`
	SpeciesID int32       `db:"species_id" json:"species_id"`
	BreedID   int32       `db:"breed_id" json:"breed_id"`
}

func (q *Queries) CreatePet(ctx context.Context, arg CreatePetParams) (Pet, error) {
	row := q.db.QueryRow(ctx, createPet,
		arg.Name,
		arg.BirthDate,
		arg.SpeciesID,
		arg.BreedID,
	)
	var i Pet
	err := row.Scan(
		&i.ID,
		&i.BirthDate,
		&i.Name,
		&i.SpeciesID,
		&i.BreedID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getPetByID = `-- name: GetPetByID :one
SELECT p.id, p.birth_date, p.name, p.species_id, p.breed_id, p.created_at, p.updated_at,
    b.name AS breed_name,
    s.name AS species_name
FROM pets p
    LEFT JOIN breeds b ON p.breed_id = b.id
    LEFT JOIN species s ON p.species_id = s.id
WHERE p.id = $1
`

type GetPetByIDRow struct {
	ID          int32            `db:"id" json:"id"`
	BirthDate   pgtype.Date      `db:"birth_date" json:"birth_date"`
	Name        string           `db:"name" json:"name"`
	SpeciesID   int32            `db:"species_id" json:"species_id"`
	BreedID     int32            `db:"breed_id" json:"breed_id"`
	CreatedAt   pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt   pgtype.Timestamp `db:"updated_at" json:"updated_at"`
	BreedName   pgtype.Text      `db:"breed_name" json:"breed_name"`
	SpeciesName pgtype.Text      `db:"species_name" json:"species_name"`
}

func (q *Queries) GetPetByID(ctx context.Context, id int32) (GetPetByIDRow, error) {
	row := q.db.QueryRow(ctx, getPetByID, id)
	var i GetPetByIDRow
	err := row.Scan(
		&i.ID,
		&i.BirthDate,
		&i.Name,
		&i.SpeciesID,
		&i.BreedID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.BreedName,
		&i.SpeciesName,
	)
	return i, err
}

const getPetsPaginated = `-- name: GetPetsPaginated :many
SELECT p.id AS pet_id,
    p.name AS pet_name,
    p.birth_date AS pet_birth_date,
    s.name AS species_name,
    b.name AS breed_name
FROM pets p
    LEFT JOIN species s ON p.species_id = s.id
    LEFT JOIN breeds b ON p.breed_id = b.id
ORDER BY p.created_at DESC
LIMIT $1 OFFSET $2
`

type GetPetsPaginatedParams struct {
	Limit  int32 `db:"limit" json:"limit"`
	Offset int32 `db:"offset" json:"offset"`
}

type GetPetsPaginatedRow struct {
	PetID        int32       `db:"pet_id" json:"pet_id"`
	PetName      string      `db:"pet_name" json:"pet_name"`
	PetBirthDate pgtype.Date `db:"pet_birth_date" json:"pet_birth_date"`
	SpeciesName  pgtype.Text `db:"species_name" json:"species_name"`
	BreedName    pgtype.Text `db:"breed_name" json:"breed_name"`
}

func (q *Queries) GetPetsPaginated(ctx context.Context, arg GetPetsPaginatedParams) ([]GetPetsPaginatedRow, error) {
	rows, err := q.db.Query(ctx, getPetsPaginated, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetPetsPaginatedRow
	for rows.Next() {
		var i GetPetsPaginatedRow
		if err := rows.Scan(
			&i.PetID,
			&i.PetName,
			&i.PetBirthDate,
			&i.SpeciesName,
			&i.BreedName,
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

const getPetsWithOwnersPaginated = `-- name: GetPetsWithOwnersPaginated :many
SELECT p.id AS pet_id,
    p.name AS pet_name,
    p.birth_date AS pet_birth_date,
    s.name AS species_name,
    b.name AS breed_name,
    u.id AS owner_id,
    CONCAT(u.first_name, ' ', u.last_name) AS owner_name,
    u.email AS owner_email,
    u.phone_number AS owner_phone
FROM pets p
    LEFT JOIN species s ON p.species_id = s.id
    LEFT JOIN breeds b ON p.breed_id = b.id
    LEFT JOIN users u ON p.owner_id = u.id
ORDER BY p.created_at DESC
LIMIT $1 OFFSET $2
`

type GetPetsWithOwnersPaginatedParams struct {
	Limit  int32 `db:"limit" json:"limit"`
	Offset int32 `db:"offset" json:"offset"`
}

type GetPetsWithOwnersPaginatedRow struct {
	PetID        int32       `db:"pet_id" json:"pet_id"`
	PetName      string      `db:"pet_name" json:"pet_name"`
	PetBirthDate pgtype.Date `db:"pet_birth_date" json:"pet_birth_date"`
	SpeciesName  pgtype.Text `db:"species_name" json:"species_name"`
	BreedName    pgtype.Text `db:"breed_name" json:"breed_name"`
	OwnerID      pgtype.Int4 `db:"owner_id" json:"owner_id"`
	OwnerName    interface{} `db:"owner_name" json:"owner_name"`
	OwnerEmail   pgtype.Text `db:"owner_email" json:"owner_email"`
	OwnerPhone   pgtype.Text `db:"owner_phone" json:"owner_phone"`
}

func (q *Queries) GetPetsWithOwnersPaginated(ctx context.Context, arg GetPetsWithOwnersPaginatedParams) ([]GetPetsWithOwnersPaginatedRow, error) {
	rows, err := q.db.Query(ctx, getPetsWithOwnersPaginated, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetPetsWithOwnersPaginatedRow
	for rows.Next() {
		var i GetPetsWithOwnersPaginatedRow
		if err := rows.Scan(
			&i.PetID,
			&i.PetName,
			&i.PetBirthDate,
			&i.SpeciesName,
			&i.BreedName,
			&i.OwnerID,
			&i.OwnerName,
			&i.OwnerEmail,
			&i.OwnerPhone,
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

const updatePet = `-- name: UpdatePet :one
UPDATE pets
SET name = $1,
    species_id = $2,
    breed_id = $3,
    updated_at = NOW()
WHERE id = $4
RETURNING id, birth_date, name, species_id, breed_id, created_at, updated_at
`

type UpdatePetParams struct {
	Name      string `db:"name" json:"name"`
	SpeciesID int32  `db:"species_id" json:"species_id"`
	BreedID   int32  `db:"breed_id" json:"breed_id"`
	ID        int32  `db:"id" json:"id"`
}

func (q *Queries) UpdatePet(ctx context.Context, arg UpdatePetParams) (Pet, error) {
	row := q.db.QueryRow(ctx, updatePet,
		arg.Name,
		arg.SpeciesID,
		arg.BreedID,
		arg.ID,
	)
	var i Pet
	err := row.Scan(
		&i.ID,
		&i.BirthDate,
		&i.Name,
		&i.SpeciesID,
		&i.BreedID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
