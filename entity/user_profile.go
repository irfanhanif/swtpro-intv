package entity

import (
	"github.com/google/uuid"
	"github.com/irfanhanif/swtpro-intv/utils"
)

//go:generate mockgen -source=user_profile.go -destination=mock/user_profile.go -package=mock

type INewUserProfile interface {
	NewUserProfile(fullName string) IUserProfile
}

type userProfileFactory struct {
	uuidGen utils.IUUID
}

func (u *userProfileFactory) NewUserProfile(
	uuidGen utils.IUUID,
	fullName string,
) IUserProfile {
	return &userProfile{
		id:       uuidGen.New(),
		fullName: fullName,
	}
}

type IUserProfile interface {
	ID() uuid.UUID
	FullName() string

	Validate() []error
}

type userProfile struct {
	id       uuid.UUID
	fullName string
}

func (u *userProfile) ID() uuid.UUID {
	return u.id
}

func (u *userProfile) FullName() string {
	return u.fullName
}

func (u *userProfile) Validate() []error {
	return nil
}
