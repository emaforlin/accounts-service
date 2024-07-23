package usecases

import (
	"github.com/emaforlin/accounts-service/x/entities"
	"github.com/emaforlin/accounts-service/x/models"
)

type AccountUsecase interface {
	GetUserId(in *models.GetUserId) (int32, error)
	AddFoodPlaceAccount(in *models.AddFoodPlaceAccountData) error
	AddPersonAccount(in *models.AddPersonAccountData) error
	GetAccountDetails(in *models.GetAccountData) (*entities.User, error)
}
