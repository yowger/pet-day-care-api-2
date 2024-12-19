package controller

type PetController interface {
}

type petController struct{}

func NewPetController() PetController {
	return &petController{}
}
