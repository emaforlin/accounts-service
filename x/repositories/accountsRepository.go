package repositories

import (
	"github.com/emaforlin/accounts-service/database"
	"github.com/emaforlin/accounts-service/x/entities"
)

type accountsMysqlRepository struct {
	db database.Database
}

// GetUser implements AccountsRepository.
func (u *accountsMysqlRepository) GetUser(in *entities.GetUserDto) (*entities.User, error) {
	result := entities.User{}
	response := u.db.GetDb().Table("users").Take(&result, in) //Model(&in).Find(&result)
	if response.Error != nil {
		return nil, response.Error
	}
	return &result, nil
}

// AddPerson implements AccountsRepository
func (u *accountsMysqlRepository) AddPerson(in *entities.InsertPersonDto) error {
	response := u.db.GetDb().Create(in)
	if response.Error != nil {
		return response.Error
	}
	return nil
}

func NewUserMysqlRepository(d database.Database) AccountsRepository {
	return &accountsMysqlRepository{
		db: d,
	}
}
