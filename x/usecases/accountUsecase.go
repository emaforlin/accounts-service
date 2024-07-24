package usecases

import (
	"errors"
	"fmt"

	"github.com/emaforlin/accounts-service/x/entities"
	"github.com/emaforlin/accounts-service/x/models"
	"github.com/emaforlin/accounts-service/x/repositories"
	"github.com/rs/zerolog/log"
)

type accountUsecaseImpl struct {
	repository repositories.AccountsRepository
}

func (u *accountUsecaseImpl) GetPersonDetails(in *models.GetAccountData) (*entities.Person, error) {
	found, err := u.repository.SelectPerson(&entities.GetPersonDto{
		UserId: in.Id,
	})
	if err != nil {
		return nil, err
	}
	return found, nil
}

func (u *accountUsecaseImpl) GetFoodPlaceDetails(in *models.GetAccountData) (*entities.FoodPlace, error) {
	found, err := u.repository.SelectFoodPlace(&entities.GetFoodPlaceDto{
		UserId: in.Id,
	})
	if err != nil {
		return nil, err
	}
	return found, nil
}

func (u *accountUsecaseImpl) GetUserId(in *models.GetUserId) (int32, error) {
	found, err := u.repository.SelectAccount(&entities.GetUserDto{
		Username:    in.Username,
		Email:       in.Email,
		PhoneNumber: in.PhoneNumber,
	})
	if err != nil {
		log.Err(err)
		return 0, errors.New("error user not found")
	}
	return int32(found.ID), nil
}

func (u *accountUsecaseImpl) AddFoodPlaceAccount(in *models.AddFoodPlaceAccountData) error {
	_, err := u.repository.SelectAccount(&entities.GetUserDto{
		Username:    in.Username,
		Email:       in.Email,
		PhoneNumber: in.PhoneNumber,
	})

	if err != nil {
		return err
	}
	var tags string

	for _, t := range in.Tags {
		tags = fmt.Sprintf("%s %s", tags, t)
	}

	dto := entities.InsertFoodPlaceDto{
		BusinessName: in.BusinessName,
		Location:     in.Location,
		User: entities.InsertUserDto{
			Role:        "FoodPlace",
			Username:    in.Username,
			Email:       in.Email,
			PhoneNumber: in.PhoneNumber,
			Password:    in.Password,
		},
		Tags: tags,
	}

	err = u.repository.InsertFoodPlace(&dto)
	if err != nil {
		return fmt.Errorf("error creating account")
	}

	return nil
}

func (u *accountUsecaseImpl) AddPersonAccount(in *models.AddPersonAccountData) error {
	_, err := u.repository.SelectAccount(&entities.GetUserDto{
		Username:    in.Username,
		Email:       in.Email,
		PhoneNumber: in.PhoneNumber,
	})

	if err != nil {
		return err
	}

	dto := entities.InsertPersonDto{
		FirstName: in.FirstName,
		LastName:  in.LastName,
		User: entities.InsertUserDto{
			Role:        "Customer",
			Username:    in.Username,
			Email:       in.Email,
			PhoneNumber: in.PhoneNumber,
			Password:    in.Password,
			Verified:    false,
		},
	}

	err = u.repository.InsertPerson(&dto)
	if err != nil {
		return fmt.Errorf("error creating account")
	}
	return nil
}

func (u *accountUsecaseImpl) GetAccountDetails(in *models.GetAccountData) (*entities.User, error) {
	found, err := u.repository.SelectAccount(&entities.GetUserDto{
		ID: in.Id,
	})

	if err != nil {
		return nil, err
	}

	found.Password = ""
	return found, nil
}

func NewAccountUsecaseImpl(repo repositories.AccountsRepository) AccountUsecase {
	return &accountUsecaseImpl{repository: repo}
}
