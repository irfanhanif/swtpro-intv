package service

import "github.com/google/uuid"

//go:generate mockgen -source=interfaces.go -destination=mock/interfaces.go -package=mock

type NewUser struct {
	PhoneNumber string
	Password    string
	FullName    string
}

type IRegisterNewUser interface {
	RegisterNewUser(newUser NewUser) (uuid.UUID, error)
}
