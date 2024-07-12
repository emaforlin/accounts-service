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

func (u *accountUsecaseImpl) VerifyFoodPlaceAccount(in *models.VerifyFoodPlaceAccount) error {
	err := u.repository.UpdateFoodPlace(in.UserId, &entities.InsertFoodPlaceDto{Verified: true})
	if err != nil {
		return errors.New("account could not be verified")
	}
	return nil
}

func (u *accountUsecaseImpl) AddFoodPlaceAccount(in *models.AddFoodPlaceAccountData) error {
	_, err := u.repository.SelectUser(&entities.GetUserDto{
		Username:    in.Username,
		Email:       in.Email,
		PhoneNumber: in.PhoneNumber,
	})

	if err != nil {
		return err
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
	}
	err = u.repository.InsertFoodPlace(&dto)
	if err != nil {
		log.Err(err)
		return fmt.Errorf("error creating account")
	}

	return nil
}

func (u *accountUsecaseImpl) AddPersonAccount(in *models.AddPersonAccountData) error {
	_, err := u.repository.SelectUser(&entities.GetUserDto{
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
		},
	}

	err = u.repository.InsertPerson(&dto)
	if err != nil {
		log.Err(err)
		return fmt.Errorf("error creating account")
	}
	return nil
}

func (u *accountUsecaseImpl) GetAccountDetails(in *models.GetAccountData) (*entities.User, error) {
	found, err := u.repository.SelectUser(&entities.GetUserDto{
		Username:    in.Username,
		Email:       in.Email,
		PhoneNumber: in.PhoneNumber,
	})

	if err != nil {
		log.Err(err)
		return nil, fmt.Errorf("error user not found")
	}

	return found, nil
}

func NewAccountUsecaseImpl(repo repositories.AccountsRepository) AccountUsecase {
	return &accountUsecaseImpl{repository: repo}
}
