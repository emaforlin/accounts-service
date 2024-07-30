package auth

import (
	"github.com/emaforlin/accounts-service/config"
	"github.com/golang-jwt/jwt/v5"
)

type JwtCustomClaims struct {
	Role   string `json:"role"`
	UserId uint32 `json:"user_id"`
	jwt.RegisteredClaims
}

type JwtFactory struct {
	Secret []byte
	TTL    uint
}

func NewJWTFactory(cfg config.Jwt) *JwtFactory {
	return &JwtFactory{
		Secret: cfg.Secret,
		TTL:    cfg.TTL,
	}
}
