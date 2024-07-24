package repositories

import (
	"strings"

	"github.com/emaforlin/accounts-service/database"
	"github.com/emaforlin/accounts-service/x/entities"
	"golang.org/x/crypto/bcrypt"
)

type accountsMysqlRepositoryImpl struct {
	db database.Database
}

func (u *accountsMysqlRepositoryImpl) InsertFoodPlace(in *entities.InsertFoodPlaceDto) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(in.User.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	// hash password
	in.User.Password = string(hash)
	response := u.db.GetDb().Create(in)
	if response.Error != nil {
		return response.Error
	}
	return nil
}

func (u *accountsMysqlRepositoryImpl) SelectFoodPlace(in *entities.GetFoodPlaceDto) (*entities.FoodPlace, error) {
	found := &entities.InsertFoodPlaceDto{}
	response := u.db.GetDb().Model(entities.GetFoodPlaceDto{}).First(&found, in)
	if response.Error != nil {
		return nil, response.Error
	}
	return &entities.FoodPlace{
		UserId:       found.UserId,
		BusinessName: found.BusinessName,
		Location:     found.Location,
		Tags:         strings.Split(found.Tags, " "),
	}, nil
}

func (u *accountsMysqlRepositoryImpl) UpdateFoodPlace(userId uint32, in *entities.InsertFoodPlaceDto) error {
	response := u.db.GetDb().Model(&entities.InsertFoodPlaceDto{}).Where("user_id = ?", userId).Updates(in)

	if response.Error != nil {
		return response.Error
	}
	return nil
}

func (u *accountsMysqlRepositoryImpl) SelectPerson(in *entities.GetPersonDto) (*entities.Person, error) {
	result := &entities.Person{}
	response := u.db.GetDb().Model(entities.GetPersonDto{}).First(&result, in)
	if response.Error != nil {
		return nil, response.Error
	}
	return result, nil
}

func (u *accountsMysqlRepositoryImpl) InsertPerson(in *entities.InsertPersonDto) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(in.User.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	// hash password
	in.User.Password = string(hash)
	response := u.db.GetDb().Model(&entities.InsertPersonDto{}).Create(in)
	if response.Error != nil {
		return response.Error
	}
	return nil
}

func (u *accountsMysqlRepositoryImpl) DeleteUser(in *entities.GetUserDto) error {
	res := u.db.GetDb().Delete(in)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (u *accountsMysqlRepositoryImpl) SelectAccount(in *entities.GetUserDto) (*entities.User, error) {
	result := &entities.User{}
	response := u.db.GetDb().Model(entities.GetUserDto{}).Find(&result, in)

	if response.Error != nil {
		return nil, response.Error
	}
	return result, nil
}

func NewAccountMysqlRepositoryImpl(d database.Database) AccountsRepository {
	return &accountsMysqlRepositoryImpl{
		db: d,
	}
}
