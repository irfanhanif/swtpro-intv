package utils

import (
	jwtLib "github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"time"
)

//go:generate mockgen -source=jwt.go -destination=mock/jwt.go -package=mock

type IGenerateJWT interface {
	GenerateJWT(userID uuid.UUID) (string, error)
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

	token := jwtLib.NewWithClaims(jwtLib.SigningMethodRS256, jwtLib.MapClaims{
		"userID": userID.String(),
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString(key)
}
