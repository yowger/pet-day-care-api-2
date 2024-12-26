package validation

import (
	"github.com/go-playground/validator/v10"
	"github.com/yowger/pet-day-care-api-2/dto"
)

func ValidateCreateUserRequest(req dto.CreateUser) map[string]string {
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
