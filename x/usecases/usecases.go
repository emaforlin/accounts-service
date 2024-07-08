package usecases

import (
	"github.com/emaforlin/accounts-service/x/entities"
	"github.com/emaforlin/accounts-service/x/models"
)

type AccountUsecase interface {
	GetAccountDetails(in *models.GetAccountData) (*entities.User, error)
	AddPersonAccount(in *models.AddPersonAccountData) error
}
