package controller

type BookController interface {
}

type bookController struct{}

func NewBookController() PetController {
	return &bookController{}
}
