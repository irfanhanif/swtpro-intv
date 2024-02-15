package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/irfanhanif/swtpro-intv/entity"
	"github.com/irfanhanif/swtpro-intv/repository"
)

type userRegistration struct {
	userAuthenticationFactory entity.INewUserAuthentication
	userProfileFactory        entity.INewUserProfile
	repo                      repository.ICreateNewUser
}

func NewUserRegistration(
	userAuthenticationFactory entity.INewUserAuthentication,
	userProfileFactory entity.INewUserProfile,
	repo repository.ICreateNewUser,
) *userRegistration {
	return &userRegistration{
		userAuthenticationFactory: userAuthenticationFactory,
		userProfileFactory:        userProfileFactory,
		repo:                      repo,
	}
}

func (u *userRegistration) RegisterNewUser(ctx context.Context, newUser NewUser) (uuid.UUID, error) {
	var errs []error

	userAuthentication := u.userAuthenticationFactory.NewUserAuthentication(newUser.PhoneNumber, newUser.Password)
	errs = append(errs, userAuthentication.Validate()...)

	userProfile := u.userProfileFactory.NewUserProfile(newUser.FullName)
	errs = append(errs, userProfile.Validate()...)

	errorsToStrings := func(errs []error) []string {
		result := []string{}
		for _, err := range errs {
			result = append(result, err.Error())
		}
		return result
	}

	errFields := &ErrFields{Errs: []string{}}
	if len(errs) > 0 {
		errFields.Errs = errorsToStrings(errs)
		return uuid.Nil, errFields
	}

	u.repo.CreateNewUser(ctx, userAuthentication, userProfile)

	return userProfile.ID(), nil
}
