package repositories

import "github.com/emaforlin/accounts-service/x/entities"

type AccountsRepository interface {
	SelectUser(in *entities.GetUserDto) (*entities.User, error)

	SelectPerson(in *entities.GetPersonDto) (*entities.Person, error)
	InsertPerson(in *entities.InsertPersonDto) error

	SelectFoodPlace(in *entities.GetFoodPlaceDto) (*entities.FoodPlace, error)
	InsertFoodPlace(in *entities.InsertFoodPlaceDto) error
	UpdateFoodPlace(userId uint32, in *entities.InsertFoodPlaceDto) error
}
