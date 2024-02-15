package service

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/irfanhanif/swtpro-intv/entity"
	repository "github.com/irfanhanif/swtpro-intv/repository"
	mockUtils "github.com/irfanhanif/swtpro-intv/utils/mock"
	"github.com/stretchr/testify/assert"
	"testing"

	mockEntity "github.com/irfanhanif/swtpro-intv/entity/mock"
)

func TestTokenGenerator_GenerateToken(t *testing.T) {
	type args struct {
		phoneNumber string
		password    string
	}
	type expectGetUserByPhoneNumber struct {
		phoneNumber string

		returnUser  entity.IUser
		returnError error
	}
	type expectIncrementLoginCount struct {
		userID uuid.UUID

		returnError error
	}
	type expectGenerateJWT struct {
		userID uuid.UUID

		returnJWT   string
		returnError error
	}
	type test struct {
		args args

		expectGetUserByPhoneNumber *expectGetUserByPhoneNumber
		expectIncrementLoginCount  *expectIncrementLoginCount
		expectGenerateJWT          *expectGenerateJWT

		wantToken  string
		wantUserID uuid.UUID
		wantError  error
	}

	tests := map[string]func(controller *gomock.Controller) test{
		"should returns token with user id and nil error " +
			"when successfully generate token": func(mockCtrl *gomock.Controller) test {
			mockUser := mockEntity.NewMockIUser(mockCtrl)
			mockUser.EXPECT().ID().Return(uuid.MustParse("be0bfd72-8d44-45fe-b410-77017b73e1e6")).AnyTimes()
			mockUser.EXPECT().HashedPassword().Return("$2a$10$kXxzCBr.T7mJzHFTjJnbd.Ww9uQQq.SO/3/3dCSdfguFw91d7Rd1.").AnyTimes()

			return test{
				args: args{
					phoneNumber: "+628123456789",
					password:    "ThisIsAPassword1234!",
				},
				expectGetUserByPhoneNumber: &expectGetUserByPhoneNumber{
					phoneNumber: "+628123456789",
					returnUser:  mockUser,
					returnError: nil,
				},
				expectIncrementLoginCount: &expectIncrementLoginCount{
					userID:      uuid.MustParse("be0bfd72-8d44-45fe-b410-77017b73e1e6"),
					returnError: nil,
				},
				expectGenerateJWT: &expectGenerateJWT{
					userID:      uuid.MustParse("be0bfd72-8d44-45fe-b410-77017b73e1e6"),
					returnJWT:   "this-is-a-token",
					returnError: nil,
				},
				wantToken:  "this-is-a-token",
				wantUserID: uuid.MustParse("be0bfd72-8d44-45fe-b410-77017b73e1e6"),
				wantError:  nil,
			}
		},
		"should returns login failed " +
			"when password is incorrect": func(mockCtrl *gomock.Controller) test {
			mockUser := mockEntity.NewMockIUser(mockCtrl)
			mockUser.EXPECT().ID().Return(uuid.MustParse("be0bfd72-8d44-45fe-b410-77017b73e1e6")).AnyTimes()
			mockUser.EXPECT().HashedPassword().Return("$2a$10$kXxzCBr.T7mJzHFTjJnbd.Ww9uQQq.SO/3/3dCSdfguFw91d7Rd1.").AnyTimes()

			return test{
				args: args{
					phoneNumber: "+628123456789",
					password:    "blabla",
				},
				expectGetUserByPhoneNumber: &expectGetUserByPhoneNumber{
					phoneNumber: "+628123456789",
					returnUser:  mockUser,
					returnError: nil,
				},
				wantToken:  "",
				wantUserID: uuid.Nil,
				wantError:  ErrLoginFailed,
			}
		},
		"should returns login failed " +
			"when get user by phone number returns no row": func(mockCtrl *gomock.Controller) test {
			mockUser := mockEntity.NewMockIUser(mockCtrl)
			mockUser.EXPECT().ID().Return(uuid.MustParse("be0bfd72-8d44-45fe-b410-77017b73e1e6")).AnyTimes()
			mockUser.EXPECT().HashedPassword().Return("$2a$10$kXxzCBr.T7mJzHFTjJnbd.Ww9uQQq.SO/3/3dCSdfguFw91d7Rd1.").AnyTimes()

			return test{
				args: args{
					phoneNumber: "+628123456789",
					password:    "ThisIsAPassword1234!",
				},
				expectGetUserByPhoneNumber: &expectGetUserByPhoneNumber{
					phoneNumber: "+628123456789",
					returnUser:  mockUser,
					returnError: repository.ErrNoRows,
				},
				wantToken:  "",
				wantUserID: uuid.Nil,
				wantError:  ErrLoginFailed,
			}
		},
		"should returns error " +
			"when get user by phone number returns error": func(mockCtrl *gomock.Controller) test {
			mockUser := mockEntity.NewMockIUser(mockCtrl)
			mockUser.EXPECT().ID().Return(uuid.MustParse("be0bfd72-8d44-45fe-b410-77017b73e1e6")).AnyTimes()
			mockUser.EXPECT().HashedPassword().Return("$2a$10$kXxzCBr.T7mJzHFTjJnbd.Ww9uQQq.SO/3/3dCSdfguFw91d7Rd1.").AnyTimes()

			return test{
				args: args{
					phoneNumber: "+628123456789",
					password:    "ThisIsAPassword1234!",
				},
				expectGetUserByPhoneNumber: &expectGetUserByPhoneNumber{
					phoneNumber: "+628123456789",
					returnUser:  mockUser,
					returnError: errors.New("oops"),
				},
				wantToken:  "",
				wantUserID: uuid.Nil,
				wantError:  errors.New("oops"),
			}
		},
		"should returns error " +
			"when failed generate jwt": func(mockCtrl *gomock.Controller) test {
			mockUser := mockEntity.NewMockIUser(mockCtrl)
			mockUser.EXPECT().ID().Return(uuid.MustParse("be0bfd72-8d44-45fe-b410-77017b73e1e6")).AnyTimes()
			mockUser.EXPECT().HashedPassword().Return("$2a$10$kXxzCBr.T7mJzHFTjJnbd.Ww9uQQq.SO/3/3dCSdfguFw91d7Rd1.").AnyTimes()

			return test{
				args: args{
					phoneNumber: "+628123456789",
					password:    "ThisIsAPassword1234!",
				},
				expectGetUserByPhoneNumber: &expectGetUserByPhoneNumber{
					phoneNumber: "+628123456789",
					returnUser:  mockUser,
					returnError: nil,
				},
				expectGenerateJWT: &expectGenerateJWT{
					userID:      uuid.MustParse("be0bfd72-8d44-45fe-b410-77017b73e1e6"),
					returnJWT:   "",
					returnError: errors.New("failed generate token"),
				},
				wantToken:  "",
				wantUserID: uuid.Nil,
				wantError:  errors.New("failed generate token"),
			}
		},
		"should returns error " +
			"when failed to increment login count": func(mockCtrl *gomock.Controller) test {
			mockUser := mockEntity.NewMockIUser(mockCtrl)
			mockUser.EXPECT().ID().Return(uuid.MustParse("be0bfd72-8d44-45fe-b410-77017b73e1e6")).AnyTimes()
			mockUser.EXPECT().HashedPassword().Return("$2a$10$kXxzCBr.T7mJzHFTjJnbd.Ww9uQQq.SO/3/3dCSdfguFw91d7Rd1.").AnyTimes()

			return test{
				args: args{
					phoneNumber: "+628123456789",
					password:    "ThisIsAPassword1234!",
				},
				expectGetUserByPhoneNumber: &expectGetUserByPhoneNumber{
					phoneNumber: "+628123456789",
					returnUser:  mockUser,
					returnError: nil,
				},
				expectIncrementLoginCount: &expectIncrementLoginCount{
					userID:      uuid.MustParse("be0bfd72-8d44-45fe-b410-77017b73e1e6"),
					returnError: errors.New("oops"),
				},
				expectGenerateJWT: &expectGenerateJWT{
					userID:      uuid.MustParse("be0bfd72-8d44-45fe-b410-77017b73e1e6"),
					returnJWT:   "this-is-a-token",
					returnError: nil,
				},
				wantToken:  "",
				wantUserID: uuid.Nil,
				wantError:  errors.New("oops"),
			}
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			mockGetUserByPhoneNumber := repository.NewMockIGetUserByPhoneNumber(mockCtrl)
			mockIncrementLoginCount := repository.NewMockIIncrementLoginCount(mockCtrl)
			mockGenerateJWT := mockUtils.NewMockIGenerateJWT(mockCtrl)

			tt := test(mockCtrl)

			if tt.expectGetUserByPhoneNumber != nil {
				e := tt.expectGetUserByPhoneNumber
				mockGetUserByPhoneNumber.EXPECT().GetUserByPhoneNumber(gomock.Any(), e.phoneNumber).Return(e.returnUser, e.returnError)
			}

			if tt.expectIncrementLoginCount != nil {
				e := tt.expectIncrementLoginCount
				mockIncrementLoginCount.EXPECT().IncrementLoginCount(gomock.Any(), e.userID).Return(e.returnError)
			}

			if tt.expectGenerateJWT != nil {
				e := tt.expectGenerateJWT
				mockGenerateJWT.EXPECT().GenerateJWT(e.userID).Return(e.returnJWT, e.returnError)
			}

			s := NewTokenGenerator(mockGetUserByPhoneNumber, mockIncrementLoginCount, mockGenerateJWT)
			actualToken, actualUserID, actualErr := s.GenerateToken(context.Background(), tt.args.phoneNumber, tt.args.password)

			assert.Equal(t, tt.wantToken, actualToken)
			assert.Equal(t, tt.wantUserID, actualUserID)
			assert.Equal(t, tt.wantError, actualErr)
		})
	}
}
