package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yowger/pet-day-care-api-2/dto"
	"github.com/yowger/pet-day-care-api-2/service"
)

type UserController interface {
	CreateUser(c echo.Context) error
	GetUser(c echo.Context) error
	UpdateUser(c echo.Context) error
	DeleteUser(c echo.Context) error
}

type userController struct {
	service service.UserService
}

// todo add auth to get id

func NewUserController(service service.UserService) UserController {
	return &userController{
		service: service,
	}
}

func (u *userController) CreateUser(c echo.Context) error {
	var createUserDTO dto.CreateUser

	if err := c.Bind(createUserDTO); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request"})
	}

	user, err := u.service.CreateUser(createUserDTO)
	if err != nil {
		// todo - add logger
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to create user"})
	}

	return c.JSON(http.StatusCreated, user)
}

func (u *userController) GetUser(c echo.Context) error {
	return nil
}

func (u *userController) UpdateUser(c echo.Context) error {
	var createUserDTO dto.UpdateUserSelf

	if err := c.Bind(createUserDTO); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request"})
	}

	user, err := u.service.UpdateUser(createUserDTO)
	if err != nil {
		// todo - add logger
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to update user"})
	}

	return c.JSON(http.StatusOK, user)
}

func (u *userController) DeleteUser(c echo.Context) error {
	var id int32 = 1
	if err := u.service.DeleteUserByID(id); err != nil {
		// todo - add logger
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to delete user"})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "User deleted successfully"})
}
