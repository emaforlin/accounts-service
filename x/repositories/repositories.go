package repositories

import "github.com/emaforlin/accounts-service/x/entities"

type AccountsRepository interface {
	GetUser(in *entities.GetUserDto) (*entities.User, error)
	AddPerson(in *entities.InsertPersonDto) error
}
