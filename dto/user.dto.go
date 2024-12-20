package dto

import (
	"github.com/go-playground/validator/v10"
)

type CreateUserReq struct {
	FirstName   string `json:"first_name" validate:"required,min=2,max=25"`
	LastName    string `json:"last_name" validate:"required,min=2,max=25"`
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required,min=4,max=25"`
	PhoneNumber string `json:"phone_number" validate:"required"`
	RoleID      int32  `json:"role_id" validate:"required"`
}

type UserRes struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	RoleID      int32  `json:"role_id"`
}

func (req *CreateUserReq) Validate() map[string]string {
	v := validator.New()

	err := v.Struct(req)
	if err != nil {
		return nil
	}

	errors := err.(validator.ValidationErrors)

	return createErrorMessages(errors)
}

func createErrorMessages(errors validator.ValidationErrors) map[string]string {
	result := make(map[string]string)
	for _, err := range errors {
		switch err.Tag() {
		case "required":
			result[err.Field()] = err.Field() + " is required."
		case "min":
			result[err.Field()] = err.Field() + " must be at least " + err.Param() + " characters long."
		case "max":
			result[err.Field()] = err.Field() + " must be at most " + err.Param() + " characters long."
		case "email":
			result[err.Field()] = "Invalid email format."
		default:
			result[err.Field()] = "Invalid value for " + err.Field() + "."
		}
	}

	return result
}
