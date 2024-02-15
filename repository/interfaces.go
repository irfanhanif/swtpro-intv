// This file contains the interfaces for the repository layer.
// The repository layer is responsible for interacting with the database.
// For testing purpose we will generate mock implementations of these
// interfaces using mockgen. See the Makefile for more information.
package repository

import (
	"context"
	"github.com/irfanhanif/swtpro-intv/entity"
)

//go:generate mockgen -source=interfaces.go -destination=mock/interfaces.go -package=mock

type RepositoryInterface interface {
	GetTestById(ctx context.Context, input GetTestByIdInput) (output GetTestByIdOutput, err error)
}

type ICreateNewUser interface {
	CreateNewUser(ctx context.Context, user entity.IUser) error
}
