package utils

import (
	"errors"
	jwtLib "github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"time"
)

//go:generate mockgen -source=jwt.go -destination=mock/jwt.go -package=mock

type CustomClaims struct {
	UserID uuid.UUID `json:"userID"`
	jwtLib.RegisteredClaims
}

type IGenerateJWT interface {
	GenerateJWT(userID uuid.UUID) (string, error)
}

type IValidateJWT interface {
	ValidateJWT(token string) (uuid.UUID, error)
}

type jwt struct {
	secretKey []byte
}

func NewJWT(secretKey []byte) *jwt {
	return &jwt{secretKey: secretKey}
}

func (j *jwt) GenerateJWT(userID uuid.UUID) (string, error) {
	key, err := jwtLib.ParseRSAPrivateKeyFromPEM(j.secretKey)
	if err != nil {
		return "", err
	}

	token := jwtLib.NewWithClaims(jwtLib.SigningMethodRS256, &CustomClaims{
		UserID: userID,
		RegisteredClaims: jwtLib.RegisteredClaims{
			ExpiresAt: jwtLib.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	})

	return token.SignedString(key)
}

func (j *jwt) ValidateJWT(token string) (uuid.UUID, error) {
	var claims *CustomClaims

	tkn, err := jwtLib.ParseWithClaims(token, claims, func(token *jwtLib.Token) (interface{}, error) {
		return j.secretKey, nil
	})
	if err != nil {
		return uuid.Nil, err
	}

	if !tkn.Valid {
		return uuid.Nil, errors.New("token invalid")
	}

	exp, err := tkn.Claims.GetExpirationTime()
	if err != nil {
		return uuid.Nil, err
	}

	if time.Now().After(exp.Time) {
		return uuid.Nil, errors.New("token expired")
	}

	return claims.UserID, nil
}
