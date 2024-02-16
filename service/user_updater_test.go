package service

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/irfanhanif/swtpro-intv/repository"
	"github.com/irfanhanif/swtpro-intv/valueobj"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserUpdater_UpdateUserByID(t *testing.T) {
	type args struct {
		id         uuid.UUID
		updateData valueobj.UserUpdateData
	}
	type expectUpdateUserID struct {
		id         uuid.UUID
		updateData valueobj.UserUpdateData

		returnError error
	}
	type test struct {
		args args

		expectUpdateUserID *expectUpdateUserID

		wantError error
	}

	tests := map[string]test{
		"should return nil error " +
			"when successfully update data": {
			args: args{
				id: uuid.MustParse("4c872457-93c2-405c-9881-65f46364d6d0"),
				updateData: valueobj.UserUpdateData{
					PhoneNumber: func() *string {
						phoneNumber := "+628123456789"
						return &phoneNumber
					}(),
					FullName: func() *string {
						fullName := "John Doe"
						return &fullName
					}(),
				},
			},
			expectUpdateUserID: &expectUpdateUserID{
				id: uuid.MustParse("4c872457-93c2-405c-9881-65f46364d6d0"),
				updateData: valueobj.UserUpdateData{
					PhoneNumber: func() *string {
						phoneNumber := "+628123456789"
						return &phoneNumber
					}(),
					FullName: func() *string {
						fullName := "John Doe"
						return &fullName
					}(),
				},
				returnError: nil,
			},
			wantError: nil,
		},
		"should return error " +
			"when full name is error": {
			args: args{
				id: uuid.MustParse("4c872457-93c2-405c-9881-65f46364d6d0"),
				updateData: valueobj.UserUpdateData{
					PhoneNumber: func() *string {
						phoneNumber := "+628123456789"
						return &phoneNumber
					}(),
					FullName: func() *string {
						fullName := "Jo"
						return &fullName
					}(),
				},
			},
			wantError: &ErrFields{Errs: []string{"Full name cannot less than 3 characters"}},
		},
		"should return error " +
			"when phone number is error": {
			args: args{
				id: uuid.MustParse("4c872457-93c2-405c-9881-65f46364d6d0"),
				updateData: valueobj.UserUpdateData{
					PhoneNumber: func() *string {
						phoneNumber := "+698123456789"
						return &phoneNumber
					}(),
					FullName: func() *string {
						fullName := "John Doe"
						return &fullName
					}(),
				},
			},
			wantError: &ErrFields{Errs: []string{"Phone Number must has +62 as a prefix"}},
		},
		"should return err not found " +
			"when repo returns err not found": {
			args: args{
				id: uuid.MustParse("4c872457-93c2-405c-9881-65f46364d6d0"),
				updateData: valueobj.UserUpdateData{
					PhoneNumber: func() *string {
						phoneNumber := "+628123456789"
						return &phoneNumber
					}(),
					FullName: func() *string {
						fullName := "John Doe"
						return &fullName
					}(),
				},
			},
			expectUpdateUserID: &expectUpdateUserID{
				id: uuid.MustParse("4c872457-93c2-405c-9881-65f46364d6d0"),
				updateData: valueobj.UserUpdateData{
					PhoneNumber: func() *string {
						phoneNumber := "+628123456789"
						return &phoneNumber
					}(),
					FullName: func() *string {
						fullName := "John Doe"
						return &fullName
					}(),
				},
				returnError: repository.ErrNoRows,
			},
			wantError: ErrNotFound,
		},
		"should return err phone number conflict " +
			"when repo returns phone number already exist": {
			args: args{
				id: uuid.MustParse("4c872457-93c2-405c-9881-65f46364d6d0"),
				updateData: valueobj.UserUpdateData{
					PhoneNumber: func() *string {
						phoneNumber := "+628123456789"
						return &phoneNumber
					}(),
					FullName: func() *string {
						fullName := "John Doe"
						return &fullName
					}(),
				},
			},
			expectUpdateUserID: &expectUpdateUserID{
				id: uuid.MustParse("4c872457-93c2-405c-9881-65f46364d6d0"),
				updateData: valueobj.UserUpdateData{
					PhoneNumber: func() *string {
						phoneNumber := "+628123456789"
						return &phoneNumber
					}(),
					FullName: func() *string {
						fullName := "John Doe"
						return &fullName
					}(),
				},
				returnError: repository.ErrPhoneNumberConflict,
			},
			wantError: ErrPhoneNumberConflict,
		},
		"should return error " +
			"when repo returns error": {
			args: args{
				id: uuid.MustParse("4c872457-93c2-405c-9881-65f46364d6d0"),
				updateData: valueobj.UserUpdateData{
					PhoneNumber: func() *string {
						phoneNumber := "+628123456789"
						return &phoneNumber
					}(),
					FullName: func() *string {
						fullName := "John Doe"
						return &fullName
					}(),
				},
			},
			expectUpdateUserID: &expectUpdateUserID{
				id: uuid.MustParse("4c872457-93c2-405c-9881-65f46364d6d0"),
				updateData: valueobj.UserUpdateData{
					PhoneNumber: func() *string {
						phoneNumber := "+628123456789"
						return &phoneNumber
					}(),
					FullName: func() *string {
						fullName := "John Doe"
						return &fullName
					}(),
				},
				returnError: errors.New("oops"),
			},
			wantError: errors.New("oops"),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			mockUpdateUserByID := repository.NewMockIUpdateUserByID(mockCtrl)

			if test.expectUpdateUserID != nil {
				e := test.expectUpdateUserID
				mockUpdateUserByID.EXPECT().UpdateUserByID(gomock.Any(), e.id, e.updateData).Return(e.returnError)
			}

			s := NewUserUpdater(mockUpdateUserByID)
			err := s.UpdateUserByID(context.Background(), test.args.id, test.args.updateData)

			assert.Equal(t, test.wantError, err)
		})
	}
}
