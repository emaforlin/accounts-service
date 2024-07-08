package models

type AddPersonAccountData struct {
	Username    string `validate:"required,min=4,max=30"`
	FirstName   string `validate:"required,max=40"`
	LastName    string `validate:"required,max=80"`
	PhoneNumber string `validate:"required,e164"`
	Email       string `validate:"required,email"`
	Password    string `validate:"min=8,max=64"`
}

type GetAccountData struct {
	Username    string `validate:"required,min=4,max=20"`
	Email       string `validate:"required,email"`
	PhoneNumber string `validate:"required,e164"`
}
