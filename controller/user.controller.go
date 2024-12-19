package controller

type UserController interface {
}

type userController struct{}

func NewUserController() PetController {
	return &userController{}
}
