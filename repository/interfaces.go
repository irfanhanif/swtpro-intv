// This file contains the interfaces for the repository layer.
// The repository layer is responsible for interacting with the database.
// For testing purpose we will generate mock implementations of these
// interfaces using mockgen. See the Makefile for more information.
package repository

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/irfanhanif/swtpro-intv/entity"
)

//go:generate mockgen -source=interfaces.go -destination=interfaces.mock.gen.go -package=repository

type RepositoryInterface interface {
	GetTestById(ctx context.Context, input GetTestByIdInput) (output GetTestByIdOutput, err error)
}

var ErrPhoneNumberConflict = errors.New("given phone number already exists")

type ICreateNewUser interface {
	CreateNewUser(ctx context.Context, user entity.IUser) error
}

var ErrNoRows = errors.New("no data found")

// Return ErrNoRows when no data found
type IGetUserByPhoneNumber interface {
	GetUserByPhoneNumber(ctx context.Context, phoneNumber string) (entity.IUser, error)
}

type IIncrementLoginCount interface {
	IncrementLoginCount(ctx context.Context, userID uuid.UUID) error
}
