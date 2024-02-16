package service

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/irfanhanif/swtpro-intv/entity"
	"github.com/irfanhanif/swtpro-intv/repository"
)

type userGetter struct {
	repo repository.IGetUserByID
}

func NewUserGetter(repo repository.IGetUserByID) *userGetter {
	return &userGetter{repo: repo}
}

func (u *userGetter) GetUserByID(ctx context.Context, id uuid.UUID) (entity.IUser, error) {
	user, err := u.repo.GetUserByID(ctx, id)
	if errors.Is(err, repository.ErrNoRows) {
		return nil, ErrNotFound
	}

	return user, err
}
