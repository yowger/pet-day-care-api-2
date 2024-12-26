package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/yowger/pet-day-care-api-2/service"
)

type UserController interface {
	CreateUser(c echo.Context) error
	GetUser(c echo.Context) error
	UpdateUser(c echo.Context) error
	DeleteUser(c echo.Context) error
}

type userController struct{
	service service.UserService
}

func NewUserController(service service.UserService) UserController {
	return &userController{
		service: service,
	}
}

func (u *userController) CreateUser(c echo.Context) error {
	user, err := u.

	return nil
}

func (u *userController) GetUser(c echo.Context) error {
	return nil
}

func (u *userController) UpdateUser(c echo.Context) error {
	return nil
}

func (u *userController) DeleteUser(c echo.Context) error {
	return nil
}
