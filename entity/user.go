package entity

import (
	"github.com/google/uuid"
	"github.com/irfanhanif/swtpro-intv/utils"
)

//go:generate mockgen -source=user.go -destination=mock/user.go -package=mock

type INewUser interface {
	NewUser(phoneNumber, password, fullName string) IUser
}

type userFactory struct {
	uuid utils.IUUID
}

func (u *userFactory) NewUser(phoneNumber, password, fullName string) IUser {
	return &user{
		id:          u.uuid.New(),
		phoneNumber: phoneNumber,
		password:    password,
		fullName:    fullName,
	}
}

func NewUserFactory(uuid utils.IUUID) INewUser {
	return &userFactory{uuid: uuid}
}

type IUser interface {
	ID() uuid.UUID
	PhoneNumber() string
	Password() string
	FullName() string

	Validate() []error
}

type user struct {
	id          uuid.UUID
	phoneNumber string
	password    string
	fullName    string
}

func (u *user) ID() uuid.UUID {
	return u.id
}

func (u *user) PhoneNumber() string {
	return u.phoneNumber
}

func (u *user) Password() string {
	return u.password
}

func (u *user) FullName() string {
	return u.fullName
}

func (u *user) Validate() []error {
	return nil
}
