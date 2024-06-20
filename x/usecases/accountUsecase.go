package usecases

import (
	"fmt"

	"github.com/emaforlin/accounts-service/x/entities"
	"github.com/emaforlin/accounts-service/x/models"
	"github.com/emaforlin/accounts-service/x/repositories"
	"github.com/rs/zerolog/log"
)

type accountUsecaseImpl struct {
	repository repositories.AccountsRepository
}

func (u *accountUsecaseImpl) DeleteAccount(in *models.DeleteAccountData) error {
	return u.repository.DeleteUser(&entities.GetUserDto{
		ID: in.Id,
	})
}

// AddPersonAccountDetails implements AccountsUsecase.
func (u *accountUsecaseImpl) AddPersonAccount(in *models.AddPersonAccountData) error {
	found, err := u.repository.SelectUser(&entities.GetUserDto{
		Username:    in.Username,
		Email:       in.Email,
		PhoneNumber: in.PhoneNumber,
	})
	if err != nil {
		return err
	}

	if found != nil {
		return fmt.Errorf("user already exists")
	}

	var dto = entities.InsertPersonDto{
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
		return err
	}
	return nil
}

// GetUserDetails implements UserUsecase.
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
