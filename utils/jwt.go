package utils

import "github.com/google/uuid"

//go:generate mockgen -source=jwt.go -destination=mock/jwt.go -package=mock

type IGenerateJWT interface {
	GenerateJWT(userID uuid.UUID) (string, error)
}

type jwt struct {
	secretKey string
}

func NewJWT(secretKey string) *jwt {
	return &jwt{secretKey: secretKey}
}

func (j *jwt) GenerateJWT(userID uuid.UUID) (string, error) {
	return "", nil
}
