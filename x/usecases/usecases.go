package usecases

import (
	"github.com/emaforlin/accounts-service/x/entities"
	protos "github.com/emaforlin/accounts-service/x/handlers/grpc/protos"
	"github.com/emaforlin/accounts-service/x/models"
)

type AccountUsecase interface {
	CheckLoginData(in *protos.CheckUserPassRequest) bool
	GetUserId(in *models.GetUserId) (int32, error)
	AddFoodPlaceAccount(in *models.AddFoodPlaceAccountData) error
	AddPersonAccount(in *models.AddPersonAccountData) error
	GetAccountDetails(in *models.GetAccountData) (*entities.User, error)
	GetFoodPlaceDetails(in *models.GetAccountData) (*entities.FoodPlace, error)
	GetPersonDetails(in *models.GetAccountData) (*entities.Person, error)
}

type AuthUsecase interface {
	Authorize(authorization []string) *entities.UserData
}
