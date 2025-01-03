package service

import (
	"context"

	"github.com/yowger/pet-day-care-api-2/dto"
	"github.com/yowger/pet-day-care-api-2/model"
	"github.com/yowger/pet-day-care-api-2/repository"
)

type UserService interface {
	CreateUser(userDto dto.CreateUser) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
	GetUserByID(id int32) (*model.User, error)
	UpdateUser(userDto dto.UpdateUserSelf) (*model.User, error)
	DeleteUserByID(id int32) error
}

type userService struct {
	ur  repository.UserRepo
	ctx context.Context
}

func NewUserService(ur repository.UserRepo, ctx context.Context) UserService {
	return &userService{ur: ur, ctx: ctx}
}

func (us *userService) CreateUser(userDto dto.CreateUser) (*model.User, error) {
	// validate dto

	// hash password

	userParams := model.User{
		FirstName:   userDto.FirstName,
		LastName:    userDto.LastName,
		Email:       userDto.Email,
		Password:    userDto.Password,
		PhoneNumber: userDto.PhoneNumber,
		RoleID:      userDto.RoleID,
	}

	user, err := us.ur.CreateUser(us.ctx, &userParams)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (us *userService) GetUserByEmail(email string) (*model.User, error) {
	return us.ur.GetUserByEmail(us.ctx, email)
}

func (us *userService) GetUserByID(id int32) (*model.User, error) {
	return us.ur.GetUserByID(us.ctx, id)
}

func (us *userService) UpdateUser(userDto dto.UpdateUserSelf) (*model.User, error) {
	userParams := model.User{
		FirstName:   userDto.FirstName,
		LastName:    userDto.LastName,
		Email:       userDto.Email,
		PhoneNumber: userDto.PhoneNumber,
	}

	return us.ur.UpdateUser(us.ctx, &userParams)
}

func (us *userService) DeleteUserByID(id int32) error {
	return us.ur.DeleteUserByID(us.ctx, id)
}
