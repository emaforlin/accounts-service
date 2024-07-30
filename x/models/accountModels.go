package models

type (
	GetUserId struct {
		Username    string `validate:""`
		Email       string `validate:""`
		PhoneNumber string `validate:""`
	}

	AddPersonAccountData struct {
		Username    string `validate:"required,min=4,max=30"`
		FirstName   string `validate:"required,max=40"`
		LastName    string `validate:"required,max=80"`
		PhoneNumber string `validate:"required,e164"`
		Email       string `validate:"required,email"`
		Password    string `validate:"min=8,max=64"`
	}

	GetAccountData struct {
		Id uint32 `validate:"required"`
	}

	AddFoodPlaceAccountData struct {
		Username     string   `validate:"required,min=4,max=20"`
		PhoneNumber  string   `validate:"required,e164"`
		Email        string   `validate:"required,email"`
		Password     string   `validate:"required,min=8,max=64"`
		BusinessName string   `validate:"required,max=256"`
		Location     string   `validate:"required,max=256"`
		Tags         []string `validate:"required"`
	}

	LoginAccount struct {
		Email       string `validate:""`
		Username    string `validate:""`
		PhoneNumber string `validate:""`
		Password    string `validate:""`
		Role        string `validate:""`
	}
)
