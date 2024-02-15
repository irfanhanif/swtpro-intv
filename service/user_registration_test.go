package service

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/irfanhanif/swtpro-intv/entity"
	mockEntity "github.com/irfanhanif/swtpro-intv/entity/mock"
	"github.com/irfanhanif/swtpro-intv/repository"
	mockRepository "github.com/irfanhanif/swtpro-intv/repository/mock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_userRegistration_RegisterNewUser(t *testing.T) {
	type args struct {
		newUser NewUser
	}
	type expectNewUser struct {
		phoneNumber string
		password    string
		fullName    string

		returnUser entity.IUser
	}
	type expectCreateNewUser struct {
		user entity.IUser

		returnError error
	}
	type test struct {
		args args

		expectNewUser       *expectNewUser
		expectCreateNewUser *expectCreateNewUser

		wantUUID uuid.UUID
		wantErr  error
	}

	tests := map[string]func(mockCtrl *gomock.Controller) test{
		"should return an uuid " +
			"and nil error" +
			"when successfully create new user": func(mockCtrl *gomock.Controller) test {
			mockUser := mockEntity.NewMockIUser(mockCtrl)
			mockUser.EXPECT().Validate().Return(nil)
			mockUser.EXPECT().ID().Return(uuid.MustParse("bd2027f3-1a36-4c18-8e12-a9f78ddbfa84")).AnyTimes()

			return test{
				args: args{
					newUser: NewUser{
						FullName:    "John Doe",
						Password:    "ThisIsAPassword1234!",
						PhoneNumber: "+6281234567890",
					},
				},
				expectNewUser: &expectNewUser{
					phoneNumber: "+6281234567890",
					password:    "ThisIsAPassword1234!",
					fullName:    "John Doe",

					returnUser: mockUser,
				},
				expectCreateNewUser: &expectCreateNewUser{
					user:        mockUser,
					returnError: nil,
				},
				wantUUID: uuid.MustParse("bd2027f3-1a36-4c18-8e12-a9f78ddbfa84"),
				wantErr:  nil,
			}
		},
		"should return nil uuid " +
			"and errors " +
			"when user authentication or profilr data in invalid": func(mockCtrl *gomock.Controller) test {
			mockUser := mockEntity.NewMockIUser(mockCtrl)
			mockUser.EXPECT().Validate().Return([]error{
				errors.New("phone number invalid"),
				errors.New("password invalid"),
				errors.New("full name invalid"),
			})
			mockUser.EXPECT().ID().Return(uuid.MustParse("bd2027f3-1a36-4c18-8e12-a9f78ddbfa84")).AnyTimes()

			return test{
				args: args{
					newUser: NewUser{
						FullName:    "John Doe",
						Password:    "ThisIsAPassword1234!",
						PhoneNumber: "+6281234567890",
					},
				},
				expectNewUser: &expectNewUser{
					phoneNumber: "+6281234567890",
					password:    "ThisIsAPassword1234!",
					fullName:    "John Doe",

					returnUser: mockUser,
				},
				wantUUID: uuid.Nil,
				wantErr: &ErrFields{
					Errs: []string{
						"phone number invalid",
						"password invalid",
						"full name invalid",
					},
				},
			}
		},
		"should return nil " +
			"and error phone number conflict " +
			"when create new user returns error phone number conflict": func(mockCtrl *gomock.Controller) test {
			mockUser := mockEntity.NewMockIUser(mockCtrl)
			mockUser.EXPECT().Validate().Return(nil)
			mockUser.EXPECT().ID().Return(uuid.MustParse("bd2027f3-1a36-4c18-8e12-a9f78ddbfa84")).AnyTimes()

			return test{
				args: args{
					newUser: NewUser{
						FullName:    "John Doe",
						Password:    "ThisIsAPassword1234!",
						PhoneNumber: "+6281234567890",
					},
				},
				expectNewUser: &expectNewUser{
					phoneNumber: "+6281234567890",
					password:    "ThisIsAPassword1234!",
					fullName:    "John Doe",

					returnUser: mockUser,
				},
				expectCreateNewUser: &expectCreateNewUser{
					user:        mockUser,
					returnError: repository.ErrPhoneNumberConflict,
				},
				wantUUID: uuid.Nil,
				wantErr:  ErrPhoneNumberConflict,
			}
		},
		"should return nil " +
			"and error " +
			"when create new user returns error": func(mockCtrl *gomock.Controller) test {
			mockUser := mockEntity.NewMockIUser(mockCtrl)
			mockUser.EXPECT().Validate().Return(nil)
			mockUser.EXPECT().ID().Return(uuid.MustParse("bd2027f3-1a36-4c18-8e12-a9f78ddbfa84")).AnyTimes()

			return test{
				args: args{
					newUser: NewUser{
						FullName:    "John Doe",
						Password:    "ThisIsAPassword1234!",
						PhoneNumber: "+6281234567890",
					},
				},
				expectNewUser: &expectNewUser{
					phoneNumber: "+6281234567890",
					password:    "ThisIsAPassword1234!",
					fullName:    "John Doe",

					returnUser: mockUser,
				},
				expectCreateNewUser: &expectCreateNewUser{
					user:        mockUser,
					returnError: errors.New("some error"),
				},
				wantUUID: uuid.Nil,
				wantErr:  errors.New("some error"),
			}
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			mockCreateNewUser := mockRepository.NewMockICreateNewUser(mockCtrl)
			mockNewUser := mockEntity.NewMockINewUser(mockCtrl)

			tt := test(mockCtrl)

			if tt.expectNewUser != nil {
				e := tt.expectNewUser
				mockNewUser.EXPECT().NewUser(e.phoneNumber, e.password, e.fullName).Return(e.returnUser)
			}

			if tt.expectCreateNewUser != nil {
				e := tt.expectCreateNewUser
				mockCreateNewUser.EXPECT().CreateNewUser(gomock.Any(), e.user).Return(e.returnError)
			}

			s := NewUserRegistration(mockNewUser, mockCreateNewUser)
			userID, actualErr := s.RegisterNewUser(context.Background(), tt.args.newUser)

			assert.Equal(t, tt.wantUUID, userID)
			assert.Equal(t, tt.wantErr, actualErr)
		})
	}
}
