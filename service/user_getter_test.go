package service

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/irfanhanif/swtpro-intv/entity"
	mockEntity "github.com/irfanhanif/swtpro-intv/entity/mock"
	"github.com/irfanhanif/swtpro-intv/repository"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserGetter_GetUserByID(t *testing.T) {
	type args struct {
		id uuid.UUID
	}
	type expectGetUserByID struct {
		id uuid.UUID

		returnUser  entity.IUser
		returnError error
	}
	type test struct {
		args args

		expectGetUserByID *expectGetUserByID

		wantUser entity.IUser
		wantErr  error
	}

	tests := map[string]func(*gomock.Controller) test{
		"should return user and nil error " +
			"when successfully get user": func(mockCtrl *gomock.Controller) test {
			mockUser := mockEntity.NewMockIUser(mockCtrl)

			return test{
				args: args{
					id: uuid.MustParse("a1468355-86f0-468d-ae67-dbd286a03ea8"),
				},
				expectGetUserByID: &expectGetUserByID{
					id:          uuid.MustParse("a1468355-86f0-468d-ae67-dbd286a03ea8"),
					returnUser:  mockUser,
					returnError: nil,
				},
				wantUser: mockUser,
				wantErr:  nil,
			}
		},
		"should return nil user and no data error " +
			"when data not found": func(mockCtrl *gomock.Controller) test {
			mockUser := mockEntity.NewMockIUser(mockCtrl)

			return test{
				args: args{
					id: uuid.MustParse("a1468355-86f0-468d-ae67-dbd286a03ea8"),
				},
				expectGetUserByID: &expectGetUserByID{
					id:          uuid.MustParse("a1468355-86f0-468d-ae67-dbd286a03ea8"),
					returnUser:  mockUser,
					returnError: repository.ErrNoRows,
				},
				wantUser: nil,
				wantErr:  ErrNotFound,
			}
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			mockGetUserByID := repository.NewMockIGetUserByID(mockCtrl)

			tt := test(mockCtrl)

			if tt.expectGetUserByID != nil {
				e := tt.expectGetUserByID
				mockGetUserByID.EXPECT().GetUserByID(gomock.Any(), e.id).Return(e.returnUser, e.returnError)
			}

			s := NewUserGetter(mockGetUserByID)
			user, err := s.GetUserByID(context.Background(), tt.args.id)

			assert.Equal(t, tt.wantUser, user)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
