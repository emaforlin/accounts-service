package usecases

import (
	"github.com/emaforlin/accounts-service/x/entities"
	"github.com/emaforlin/accounts-service/x/models"
)

type AccountsUsecase interface {
	// Account related
	GetAccountDetails(in *models.GetAccountData) (*entities.User, error)
	DeleteAccount(in *models.DeleteAccountData) error

	// Person entity related
	AddPersonAccountDetails(in *models.AddPersonAccountData) error
}
