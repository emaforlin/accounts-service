package usecases

import (
	"github.com/emaforlin/accounts-service/x/entities"
	"github.com/emaforlin/accounts-service/x/models"
)

type AccountUsecase interface {
	// Accounts in general related
	GetAccountDetails(in *models.GetAccountData) (*entities.User, error)
	DeleteAccount(in *models.DeleteAccountData) error

	// Accounts of type Person related
	AddPersonAccountDetails(in *models.AddPersonAccountData) error
}
