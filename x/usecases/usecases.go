package usecases

import (
	"github.com/emaforlin/accounts-service/x/entities"
	"github.com/emaforlin/accounts-service/x/models"
)

type AccountUsecase interface {
	VerifyFoodPlaceAccount(in *models.VerifyFoodPlaceAccount) error
	AddFoodPlaceAccount(in *models.AddFoodPlaceAccountData) error
	AddPersonAccount(in *models.AddPersonAccountData) error
	GetAccountDetails(in *models.GetAccountData) (*entities.User, error)
}
