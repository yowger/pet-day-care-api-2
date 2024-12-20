package controller

type BookController interface {
}

type bookController struct{}

func NewBookController() BookController {
	return &bookController{}
}
