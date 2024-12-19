// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package database

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Booking struct {
	ID        int32            `db:"id" json:"id"`
	UserID    int32            `db:"user_id" json:"user_id"`
	PetID     int32            `db:"pet_id" json:"pet_id"`
	StartTime pgtype.Date      `db:"start_time" json:"start_time"`
	EndTime   pgtype.Date      `db:"end_time" json:"end_time"`
	Comments  pgtype.Text      `db:"comments" json:"comments"`
	CreatedAt pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type Breed struct {
	ID        int32            `db:"id" json:"id"`
	Name      string           `db:"name" json:"name"`
	SpeciesID int32            `db:"species_id" json:"species_id"`
	CreatedAt pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type Pet struct {
	ID        int32            `db:"id" json:"id"`
	BirthDate pgtype.Date      `db:"birth_date" json:"birth_date"`
	Name      string           `db:"name" json:"name"`
	SpeciesID int32            `db:"species_id" json:"species_id"`
	BreedID   int32            `db:"breed_id" json:"breed_id"`
	CreatedAt pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type Role struct {
	ID        int32            `db:"id" json:"id"`
	Name      string           `db:"name" json:"name"`
	CreatedAt pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type Species struct {
	ID        int32            `db:"id" json:"id"`
	Name      string           `db:"name" json:"name"`
	CreatedAt pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type User struct {
	ID          int32            `db:"id" json:"id"`
	FirstName   string           `db:"first_name" json:"first_name"`
	LastName    string           `db:"last_name" json:"last_name"`
	Email       string           `db:"email" json:"email"`
	PhoneNumber string           `db:"phone_number" json:"phone_number"`
	Password    string           `db:"password" json:"password"`
	RoleID      int32            `db:"role_id" json:"role_id"`
	CreatedAt   pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt   pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}
