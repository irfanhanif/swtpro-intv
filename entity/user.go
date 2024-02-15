package entity

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/irfanhanif/swtpro-intv/utils"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"unicode"
)

const (
	INDONESIA_PHONE_CODE   = "+62"
	MIN_PHONE_NUMBER_CHARS = 10
	MAX_PHONE_NUMBER_CHARS = 13

	MIN_PASSWORD_CHARS = 6
	MAX_PASSWORD_CHARS = 64

	MIN_FULL_NAME_CHARS = 3
	MAX_FULL_NAME_CHARS = 60
)

//go:generate mockgen -source=user.go -destination=mock/user.go -package=mock

type INewUser interface {
	NewUser(phoneNumber, password, fullName string) IUser
	NewUserWithID(id uuid.UUID, phoneNumber, password, fullName string) IUser
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

func (u *userFactory) NewUserWithID(id uuid.UUID, phoneNumber, password, fullName string) IUser {
	return &user{
		id:          id,
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
	HashedPassword() string
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

func (u *user) HashedPassword() string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(u.password), 10)
	return string(bytes)
}

func (u *user) FullName() string {
	return u.fullName
}

func (u *user) Validate() []error {
	errs := []error{}

	if len(u.phoneNumber) < MIN_PHONE_NUMBER_CHARS {
		errs = append(errs, fmt.Errorf("Phone Number must minimum has %d characters", MIN_PHONE_NUMBER_CHARS))
	}
	if len(u.phoneNumber) > MAX_PHONE_NUMBER_CHARS {
		errs = append(errs, fmt.Errorf("Phone Number must maximum has %d characters", MAX_PHONE_NUMBER_CHARS))
	}
	if !strings.HasPrefix(u.phoneNumber, INDONESIA_PHONE_CODE) {
		errs = append(errs, fmt.Errorf("Phone Number must has %s as a prefix", INDONESIA_PHONE_CODE))
	}

	if len(u.password) < MIN_PASSWORD_CHARS {
		errs = append(errs, fmt.Errorf("Password cannot less than %d characters", MIN_PASSWORD_CHARS))
	}
	if len(u.password) > MAX_PASSWORD_CHARS {
		errs = append(errs, fmt.Errorf("Password cannot more than %d characters", MAX_PASSWORD_CHARS))
	}
	if !u.checkPassword(u.password) {
		errs = append(errs, errors.New("Password must have a capital letter, a number, a special character (non alpha numberic)"))
	}

	if len(u.fullName) < MIN_FULL_NAME_CHARS {
		errs = append(errs, fmt.Errorf("Full name cannot less than %d characters", MIN_FULL_NAME_CHARS))
	}
	if len(u.fullName) > MAX_FULL_NAME_CHARS {
		errs = append(errs, fmt.Errorf("Full name cannot more than %d characters", MAX_FULL_NAME_CHARS))
	}

	if len(errs) > 0 {
		return errs
	}

	return nil
}

func (u *user) checkPassword(password string) bool {
	lowerCasePresent := false
	uppercasePresent := false
	numberPresent := false
	specialCharacterPresent := false

	for _, c := range password {
		switch {
		case unicode.IsUpper(c):
			uppercasePresent = true
		case unicode.IsLower(c):
			lowerCasePresent = true
		case unicode.IsNumber(c):
			numberPresent = true
		case unicode.IsSymbol(c) || unicode.IsPunct(c):
			specialCharacterPresent = true
		}
	}

	return lowerCasePresent && uppercasePresent && numberPresent && specialCharacterPresent
}
