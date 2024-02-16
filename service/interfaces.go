package service

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/irfanhanif/swtpro-intv/entity"
	"github.com/irfanhanif/swtpro-intv/valueobj"
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
	RegisterNewUser(ctx context.Context, newUser NewUser) (uuid.UUID, error)
}

var ErrLoginFailed = errors.New("login failed")

// Returns ErrLoginFailed when login attempt is not allowed
type IGenerateToken interface {
	GenerateToken(ctx context.Context, phoneNumber, password string) (string, uuid.UUID, error)
}

var ErrNotFound = errors.New("data not found")

// Returns ErrNotFound when user is not found
type IGetUserByID interface {
	GetUserByID(ctx context.Context, id uuid.UUID) (entity.IUser, error)
}

// Returns ErrNotFound when user is not found
// Returns ErrPhoneNumberConflict when phone number already exists
type IUpdateUserByID interface {
	UpdateUserByID(ctx context.Context, id uuid.UUID, updateData valueobj.UserUpdateData) error
}
