package usecases

import (
	"errors"
	"strings"

	"github.com/emaforlin/accounts-service/auth"
	"github.com/emaforlin/accounts-service/config"
	"github.com/emaforlin/accounts-service/x/entities"
	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrEmptyAuthorization = errors.New("received empty authorization")
	ErrInvalidToken       = errors.New("invalid token")
	ErrInvalidUserData    = errors.New("invalid")
)

type authUsecaseImpl struct {
	auth *auth.JwtFactory
}

func (u *authUsecaseImpl) Authorize(authorization []string) *entities.UserData {
	if len(authorization) < 1 {
		return nil
	}
	tokenStr := strings.TrimPrefix(authorization[0], "Bearer ")

	claims := &auth.JwtCustomClaims{}

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
		cfg := config.LoadConfig()
		return cfg.Jwt.Secret, nil
	})

	if err != nil {
		return nil
	}

	if !token.Valid {
		return nil
	}

	if claims.Role == "" || claims.UserId == 0 {
		return nil
	}

	return &entities.UserData{
		Id:   claims.UserId,
		Role: claims.Role,
	}
}

func NewAuthUsecaseImpl(factory *auth.JwtFactory) AuthUsecase {
	return &authUsecaseImpl{
		auth: factory,
	}
}
