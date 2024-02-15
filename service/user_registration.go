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
	userAuthentication := u.userAuthenticationFactory.NewUserAuthentication(newUser.PhoneNumber, newUser.Password)
	userProfile := u.userProfileFactory.NewUserProfile(newUser.FullName)

	u.repo.CreateNewUser(ctx, userAuthentication, userProfile)

	return userProfile.ID(), nil
}
