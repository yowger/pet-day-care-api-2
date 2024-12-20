package controller

import "github.com/labstack/echo/v4"

type UserController interface {
	CreateUser(c echo.Context) error
	GetUser(c echo.Context) error
	UpdateUser(c echo.Context) error
	DeleteUser(c echo.Context) error
}

type userController struct{}

func NewUserController() UserController {
	return &userController{}
}

func (u *userController) CreateUser(c echo.Context) error {
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
