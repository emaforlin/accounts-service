package models

type AddPersonAccountData struct {
	Username    string `json:"username" validate:"required,min=4,max=30"`
	FirstName   string `json:"first_name" validate:"required,max=40"`
	LastName    string `json:"last_name" validate:"required,max=80"`
	PhoneNumber string `json:"phone_number" validate:"required,e164"`
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"min=8,max=64"`
}

type GetAccountData struct {
	Username    string `validate:"required,min=4,max=20"`
	Email       string `validate:"required,email"`
	PhoneNumber string `validate:"required,e164"`
}
