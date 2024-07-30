package usecases

import (
	"errors"
	"fmt"
	"time"

	"github.com/emaforlin/accounts-service/auth"
	"github.com/emaforlin/accounts-service/config"
	"github.com/emaforlin/accounts-service/x/entities"
	protos "github.com/emaforlin/accounts-service/x/handlers/grpc/protos"
	"github.com/emaforlin/accounts-service/x/models"
	"github.com/emaforlin/accounts-service/x/repositories"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrUserNotFound        = errors.New("user not found")
	ErrFailedCreateAccount = errors.New("create new account failed")
)

type accountUsecaseImpl struct {
	repository repositories.AccountsRepository
	jwtManager *auth.JwtFactory
}

func (u *accountUsecaseImpl) Login(in *protos.LoginUserRequest) (string, error) {
	user, err := u.repository.SelectAccount(&entities.GetUserDto{
		Email:       in.GetEmail(),
		Username:    in.GetUsername(),
		PhoneNumber: in.GetPhoneNumber(),
		Role:        in.GetRole(),
	})
	if err != nil {
		return "", err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(in.Password)); err != nil {
		return "", err
	}
	claims := auth.JwtCustomClaims{
		UserId: user.ID,
		Role:   user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(u.jwtManager.TTL) * time.Second)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(u.jwtManager.Secret)
	if err != nil {
		return "", err
	}
	return ss, nil
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
		return 0, ErrUserNotFound
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
		return ErrFailedCreateAccount
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
		return ErrFailedCreateAccount
	}
	return nil
}

func (u *accountUsecaseImpl) GetAccountDetails(in *models.GetAccountData) (*entities.User, error) {
	found, err := u.repository.SelectAccount(&entities.GetUserDto{
		ID: in.Id,
	})

	if err != nil {
		return nil, ErrUserNotFound
	}

	found.Password = ""
	return found, nil
}

func NewAccountUsecaseImpl(repo repositories.AccountsRepository, cfg config.Config) AccountUsecase {
	return &accountUsecaseImpl{
		repository: repo,
		jwtManager: auth.NewJWTFactory(cfg.Jwt),
	}
}
