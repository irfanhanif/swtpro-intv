package service

import (
	"github.com/google/uuid"
	"strings"
)

//go:generate mockgen -source=interfaces.go -destination=mock/interfaces.go -package=mock

type NewUser struct {
	PhoneNumber string
	Password    string
	FullName    string
}

type FieldErrors struct {
	Errs []string
}

func (fe *FieldErrors) Error() string {
	return strings.Join(fe.Errs, ", ")
}

// Returns FieldErrors when input doesn't meed the required conditions
type IRegisterNewUser interface {
	RegisterNewUser(newUser NewUser) (uuid.UUID, error)
}
