package dto

type CreateUser struct {
	FirstName   string `json:"first_name" validate:"required,min=2,max=25"`
	LastName    string `json:"last_name" validate:"required,min=2,max=25"`
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required,min=4,max=25"`
	PhoneNumber string `json:"phone_number" validate:"required"`
	RoleID      int32  `json:"role_id" validate:"required"`
}

type UpdateUserAdmin struct {
	FirstName   string `json:"first_name,omitempty" validate:"omitempty,min=2,max=25"`
	LastName    string `json:"last_name,omitempty" validate:"omitempty,min=2,max=25"`
	Email       string `json:"email,omitempty" validate:"omitempty,email"`
	PhoneNumber string `json:"phone_number,omitempty" validate:"omitempty"`
	RoleID      int32  `json:"role_id,omitempty" validate:"omitempty,min=1"`
}

type UpdateUserSelf struct {
	FirstName   string `json:"first_name,omitempty" validate:"omitempty,min=2,max=25"`
	LastName    string `json:"last_name,omitempty" validate:"omitempty,min=2,max=25"`
	Email       string `json:"email,omitempty" validate:"omitempty,email"`
	PhoneNumber string `json:"phone_number,omitempty" validate:"omitempty"`
}
