package dto

type UserDTO struct {
	FirstName   string `json:"first_name" validate:"required,min=2,max=25"`
	LastName    string `json:"last_name" validate:"required,min=2,max=25"`
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required,min=4,max=25"`
	PhoneNumber string `json:"phone_number" validate:"required"`
}

func NewUserDTO() UserDTO {
	return UserDTO{}
}

func (u UserDTO) Create() *UserDTO {
	return &UserDTO{
		FirstName:   u.FirstName,
		LastName:    u.LastName,
		Email:       u.Email,
		Password:    u.Password,
		PhoneNumber: u.PhoneNumber,
	}
}

func (u UserDTO) UpdateSelf() *UserDTO {
	return &UserDTO{
		FirstName:   u.FirstName,
		LastName:    u.LastName,
		Email:       u.Email,
		PhoneNumber: u.PhoneNumber,
	}
}

func (u UserDTO) UpdatePassword() *UserDTO {
	return &UserDTO{
		Password: u.Password,
	}
}
