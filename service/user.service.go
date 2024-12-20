package service

import (
	"context"
	"log"

	database "github.com/yowger/pet-day-care-api-2/database/sqlc"
	"github.com/yowger/pet-day-care-api-2/dto"
)

type UserService interface {
	CreateUser(dto *dto.CreateUserReq) (*dto.UserRes, map[string]string)
	FindUserByID(id int32) (*database.User, error)
	FindAllUsersByPage(page, size string)
	UpdateUser()
	DeleteUser(id int)
}

type userService struct {
	queries database.Queries
}

func NewUserService() UserService {
	return &userService{}
}

func (us *userService) CreateUser(req *dto.CreateUserReq) (*dto.UserRes, map[string]string) {
	if errs := req.Validate(); errs != nil {
		return nil, errs
	}

	params := database.CreateUserParams{
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		Password:    req.Password,
		RoleID:      req.RoleID,
	}
	user, err := us.queries.CreateUser(context.Background(), params)
	if err != nil {
		log.Println(err)

		return nil, map[string]string{"error": "failed to create user"}
	}

	res := dto.UserRes{
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		RoleID:      user.RoleID,
	}

	return &res, nil
}

func (us *userService) FindUserByID(id int32) (*database.User, error) {
	user, err := us.queries.FindUserByID(context.Background(), id)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (us *userService) FindAllUsersByPage(page, size string) {
	// users, err := us.queries.Find
}

func (us *userService) UpdateUser() {
	// userParams := database.UpdateUserParams{}
	// user, err := us.queries.UpdateUser(context.Background(), userParams)
}

func (us *userService) DeleteUser(id int) {
}
