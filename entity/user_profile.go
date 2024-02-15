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

func (u *userProfileFactory) NewUserProfile(fullName string) IUserProfile {
	return nil
}

type IUserProfile interface {
	ID() uuid.UUID
	FullName() string
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
