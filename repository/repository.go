package repository

import (
	"context"

	db "github.com/yowger/pet-day-care-api-2/database/sqlc"
	"github.com/yowger/pet-day-care-api-2/model"
)

type UserRepo interface {
	CreateUser(ctx context.Context, userParams *model.User) (*model.User, error)
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
	GetUserByID(ctx context.Context, id int32) (*model.User, error)
	UpdateUser(ctx context.Context, userParams *model.User) (*model.User, error)
	DeleteUser(ctx context.Context, id int32) error
}

type userRepo struct {
	queries *db.Queries
}

func NewUserRepo(queries *db.Queries) UserRepo {
	return &userRepo{
		queries: queries}
}

func (ur *userRepo) CreateUser(ctx context.Context, userParams *model.User) (*model.User, error) {
	dbParams := db.CreateUserParams{
		FirstName:   userParams.FirstName,
		LastName:    userParams.LastName,
		Email:       userParams.Email,
		Password:    userParams.Password,
		PhoneNumber: userParams.PhoneNumber,
		RoleID:      userParams.RoleID,
	}

	user, err := ur.queries.CreateUser(ctx, dbParams)

	return toDomainUser(&user), err
}

func (ur *userRepo) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	user, err := ur.queries.GetUserByEmail(ctx, email)

	return toDomainUser(&user), err
}

func (ur *userRepo) GetUserByID(ctx context.Context, id int32) (*model.User, error) {
	user, err := ur.queries.GetUserByID(ctx, id)

	return toDomainUser(&user), err
}

func (ur *userRepo) UpdateUser(ctx context.Context, userParams *model.User) (*model.User, error) {
	dbParams := db.UpdateUserParams{
		ID:          userParams.ID,
		FirstName:   userParams.FirstName,
		LastName:    userParams.LastName,
		Email:       userParams.Email,
		PhoneNumber: userParams.PhoneNumber,
		RoleID:      userParams.RoleID,
	}

	user, err := ur.queries.UpdateUser(ctx, dbParams)

	return toDomainUser(&user), err
}

func (ur *userRepo) DeleteUser(ctx context.Context, id int32) error {
	return ur.queries.DeleteUser(ctx, id)
}

func toDomainUser(dbUser *db.User) *model.User {
	return &model.User{
		ID:          dbUser.ID,
		FirstName:   dbUser.FirstName,
		LastName:    dbUser.LastName,
		Email:       dbUser.Email,
		Password:    dbUser.Password,
		PhoneNumber: dbUser.PhoneNumber,
		RoleID:      dbUser.RoleID,
	}
}
