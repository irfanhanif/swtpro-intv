package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/irfanhanif/swtpro-intv/repository"
	"github.com/irfanhanif/swtpro-intv/utils"
	"golang.org/x/crypto/bcrypt"
)

type tokenGenerator struct {
	getUser  repository.IGetUserByPhoneNumber
	incLogin repository.IIncrementLoginCount
	jwt      utils.IGenerateJWT
}

func NewTokenGenerator(getUser repository.IGetUserByPhoneNumber, incLogin repository.IIncrementLoginCount, jwt utils.IGenerateJWT) *tokenGenerator {
	return &tokenGenerator{
		getUser:  getUser,
		incLogin: incLogin,
		jwt:      jwt,
	}
}

func (tk *tokenGenerator) GenerateToken(ctx context.Context, phoneNumber, password string) (string, uuid.UUID, error) {
	user, err := tk.getUser.GetUserByPhoneNumber(ctx, phoneNumber)
	if errors.Is(err, repository.ErrNoRows) {
		fmt.Println("tidak ketemu")
		return "", uuid.Nil, ErrLoginFailed
	}
	if err != nil {
		return "", uuid.Nil, err
	}

	fmt.Println(user.HashedPassword(), user.Password())

	if err := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword()), []byte(password)); err != nil {
		fmt.Println("pas compare", err)
		return "", uuid.Nil, ErrLoginFailed
	}

	token, err := tk.jwt.GenerateJWT(user.ID())
	if err != nil {
		return "", uuid.Nil, err
	}

	if err := tk.incLogin.IncrementLoginCount(ctx, user.ID()); err != nil {
		return "", uuid.Nil, err
	}

	return token, user.ID(), nil
}
