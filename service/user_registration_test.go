package service

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/irfanhanif/swtpro-intv/entity"
	mockEntity "github.com/irfanhanif/swtpro-intv/entity/mock"
	mockRepository "github.com/irfanhanif/swtpro-intv/repository/mock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_userRegistration_RegisterNewUser(t *testing.T) {
	type args struct {
		newUser NewUser
	}
	type expectNewUserAuthentication struct {
		phoneNumber string
		password    string

		returnUserAuthentication entity.IUserAuthentication
	}
	type expectNewUserProfile struct {
		fullName string

		returnUserProfile entity.IUserProfile
	}
	type expectCreateNewUser struct {
		userAuthentication entity.IUserAuthentication
		userProfile        entity.IUserProfile

		returnError error
	}
	type test struct {
		args args

		expectNewUserAuthentication *expectNewUserAuthentication
		expectNewUserProfile        *expectNewUserProfile
		expectCreateNewUser         *expectCreateNewUser

		wantUUID uuid.UUID
		wantErr  error
	}

	tests := map[string]func(mockCtrl *gomock.Controller) test{
		"should return an uuid " +
			"and nil error" +
			"when successfully create new user": func(mockCtrl *gomock.Controller) test {
			mockUserAuthentication := mockEntity.NewMockIUserAuthentication(mockCtrl)
			mockUserAuthentication.EXPECT().Validate().Return([]error{})

			mockUserProfile := mockEntity.NewMockIUserProfile(mockCtrl)
			mockUserProfile.EXPECT().Validate().Return([]error{})
			mockUserProfile.EXPECT().ID().Return(uuid.MustParse("bd2027f3-1a36-4c18-8e12-a9f78ddbfa84")).AnyTimes()

			return test{
				args: args{
					newUser: NewUser{
						FullName:    "John Doe",
						Password:    "ThisIsAPassword1234!",
						PhoneNumber: "+6281234567890",
					},
				},
				expectNewUserAuthentication: &expectNewUserAuthentication{
					phoneNumber: "+6281234567890",
					password:    "ThisIsAPassword1234!",

					returnUserAuthentication: mockUserAuthentication,
				},
				expectNewUserProfile: &expectNewUserProfile{
					fullName: "John Doe",

					returnUserProfile: mockUserProfile,
				},
				expectCreateNewUser: &expectCreateNewUser{
					userAuthentication: mockUserAuthentication,
					userProfile:        mockUserProfile,
					returnError:        nil,
				},
				wantUUID: uuid.MustParse("bd2027f3-1a36-4c18-8e12-a9f78ddbfa84"),
				wantErr:  nil,
			}
		},
		"should return nil uuid " +
			"and errors " +
			"when user authentication or profilr data in invalid": func(mockCtrl *gomock.Controller) test {
			mockUserAuthentication := mockEntity.NewMockIUserAuthentication(mockCtrl)
			mockUserAuthentication.EXPECT().Validate().Return([]error{
				errors.New("phone number invalid"),
				errors.New("password invalid"),
			})

			mockUserProfile := mockEntity.NewMockIUserProfile(mockCtrl)
			mockUserProfile.EXPECT().Validate().Return([]error{
				errors.New("full name invalid"),
			})
			mockUserProfile.EXPECT().ID().Return(uuid.MustParse("bd2027f3-1a36-4c18-8e12-a9f78ddbfa84")).AnyTimes()

			return test{
				args: args{
					newUser: NewUser{
						FullName:    "John Doe",
						Password:    "ThisIsAPassword1234!",
						PhoneNumber: "+6281234567890",
					},
				},
				expectNewUserAuthentication: &expectNewUserAuthentication{
					phoneNumber: "+6281234567890",
					password:    "ThisIsAPassword1234!",

					returnUserAuthentication: mockUserAuthentication,
				},
				expectNewUserProfile: &expectNewUserProfile{
					fullName: "John Doe",

					returnUserProfile: mockUserProfile,
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
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			mockNewUserAuthentication := mockEntity.NewMockINewUserAuthentication(mockCtrl)
			mockNewUserProfile := mockEntity.NewMockINewUserProfile(mockCtrl)
			mockCreateNewUser := mockRepository.NewMockICreateNewUser(mockCtrl)

			tt := test(mockCtrl)

			if tt.expectNewUserAuthentication != nil {
				e := tt.expectNewUserAuthentication
				mockNewUserAuthentication.EXPECT().NewUserAuthentication(e.phoneNumber, e.password).Return(e.returnUserAuthentication)
			}

			if tt.expectNewUserProfile != nil {
				e := tt.expectNewUserProfile
				mockNewUserProfile.EXPECT().NewUserProfile(e.fullName).Return(e.returnUserProfile)
			}

			if tt.expectCreateNewUser != nil {
				e := tt.expectCreateNewUser
				mockCreateNewUser.EXPECT().CreateNewUser(gomock.Any(), e.userAuthentication, e.userProfile).Return(e.returnError)
			}

			s := NewUserRegistration(mockNewUserAuthentication, mockNewUserProfile, mockCreateNewUser)
			userID, actualErr := s.RegisterNewUser(context.Background(), tt.args.newUser)

			assert.Equal(t, tt.wantUUID, userID)
			assert.Equal(t, tt.wantErr, actualErr)
		})
	}
}
