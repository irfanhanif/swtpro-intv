package service

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/irfanhanif/swtpro-intv/entity"
	"github.com/irfanhanif/swtpro-intv/repository"
)

type userRegistration struct {
	userFactory entity.INewUser
	repo        repository.ICreateNewUser
}

func NewUserRegistration(userFactory entity.INewUser, repo repository.ICreateNewUser) *userRegistration {
	return &userRegistration{
		userFactory: userFactory,
		repo:        repo,
	}
}

func (u *userRegistration) RegisterNewUser(ctx context.Context, newUser NewUser) (uuid.UUID, error) {
	errorsToStrings := func(errs []error) []string {
		result := []string{}
		for _, err := range errs {
			result = append(result, err.Error())
		}
		return result
	}

	user := u.userFactory.NewUser(newUser.PhoneNumber, newUser.Password, newUser.FullName)
	if errs := user.Validate(); errs != nil {
		errFields := &ErrFields{Errs: []string{}}
		errFields.Errs = errorsToStrings(errs)
		return uuid.Nil, errFields
	}

	err := u.repo.CreateNewUser(ctx, user)
	if errors.Is(err, ErrPhoneNumberConflict) {
		return uuid.Nil, ErrPhoneNumberConflict
	}
	if err != nil {
		return uuid.Nil, err
	}

	return user.ID(), nil
}
