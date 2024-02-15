package service

import (
	"errors"
	"github.com/google/uuid"
	"strings"
)

//go:generate mockgen -source=interfaces.go -destination=mock/interfaces.go -package=mock

var ErrPhoneNumberConflict = errors.New("given phone number already exists")

type NewUser struct {
	PhoneNumber string
	Password    string
	FullName    string
}

type ErrFields struct {
	Errs []string
}

func (fe *ErrFields) Error() string {
	return strings.Join(fe.Errs, ", ")
}

// Returns ErrFields when input doesn't meed the required conditions
// Returns ErrPhoneNumberConflict when phone number already exists
type IRegisterNewUser interface {
	RegisterNewUser(newUser NewUser) (uuid.UUID, error)
}
