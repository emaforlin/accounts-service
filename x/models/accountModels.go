package models

// person related models
type AddPersonAccountData struct {
	Username    string `json:"username,omitempty" validate:"required,min=4,max=20"`
	FirstName   string `json:"first_name,omitempty" validate:"required,min=1,max=20"`
	LastName    string `json:"last_name,omitempty" validate:"required,min=1,max=20"`
	PhoneNumber string `json:"phone_number" validate:"required,startswith=+,max=15"`
	Email       string `json:"email,omitempty" validate:"required,email"`
	Password    string `json:"password,omitempty" validate:"min=8,max=64"`
}

// account related models
type GetAccountData struct {
	Username    string
	Email       string
	PhoneNumber string
}

type DeleteAccountData struct {
	Id uint32
}
