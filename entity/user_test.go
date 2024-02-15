package entity

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	mockUtils "github.com/irfanhanif/swtpro-intv/utils/mock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserAttributes(t *testing.T) {
	phoneNumber := "+6281234567890"
	password := "ThisIsAPassword1234!"
	fullName := "John Doe"

	mockCtrl := gomock.NewController(t)
	mockUUID := mockUtils.NewMockIUUID(mockCtrl)
	mockUUID.EXPECT().New().Return(uuid.MustParse("0ff1a3c8-dffd-48ae-9b97-9a0c6c484368"))

	user := NewUserFactory(mockUUID).NewUser(phoneNumber, password, fullName)

	assert.Equal(t, phoneNumber, user.PhoneNumber())
	assert.Equal(t, password, user.Password())
	assert.Equal(t, fullName, user.FullName())
}

func Test_user_Validate(t *testing.T) {
	type fields struct {
		id          uuid.UUID
		phoneNumber string
		password    string
		fullName    string
	}
	tests := []struct {
		name   string
		fields fields
		want   []error
	}{
		{
			name: "should returns no error when all attributes are valid",
			fields: fields{
				id:          uuid.MustParse("7ef62378-e51c-4d02-ba8b-b98a8a1375a9"),
				phoneNumber: "+6281234567890",
				password:    "ThisIsAPassword1234!",
				fullName:    "John Doe",
			},
			want: nil,
		},
		{
			name: "should returns phone number error when it is not has +62 prefix",
			fields: fields{
				id:          uuid.MustParse("7ef62378-e51c-4d02-ba8b-b98a8a1375a9"),
				phoneNumber: "+6981234567890",
				password:    "ThisIsAPassword1234!",
				fullName:    "John Doe",
			},
			want: []error{
				errors.New("Phone Number must has +62 as a prefix"),
			},
		},
		{
			name: "should returns phone number error when char is less than 3",
			fields: fields{
				id:          uuid.MustParse("7ef62378-e51c-4d02-ba8b-b98a8a1375a9"),
				phoneNumber: "+6",
				password:    "ThisIsAPassword1234!",
				fullName:    "John Doe",
			},
			want: []error{
				errors.New("Phone Number must minimum has 3 characters"),
				errors.New("Phone Number must has +62 as a prefix"),
			},
		},
		{
			name: "should returns phone number error when char is greater than 60",
			fields: fields{
				id:          uuid.MustParse("7ef62378-e51c-4d02-ba8b-b98a8a1375a9"),
				phoneNumber: "+6281234567890+6981234567890+6981234567890+6981234567890+6981234567890+6981234567890",
				password:    "ThisIsAPassword1234!",
				fullName:    "John Doe",
			},
			want: []error{
				errors.New("Phone Number must maximum has 60 characters"),
			},
		},
		{
			name: "should returns password errors no uppercase letter",
			fields: fields{
				id:          uuid.MustParse("7ef62378-e51c-4d02-ba8b-b98a8a1375a9"),
				phoneNumber: "+6281234567890",
				password:    "thisisapassword1234!",
				fullName:    "John Doe",
			},
			want: []error{
				errors.New("Password must have a capital letter, a number, a special character (non alpha numberic)"),
			},
		},
		{
			name: "should returns password errors no lowercase letter",
			fields: fields{
				id:          uuid.MustParse("7ef62378-e51c-4d02-ba8b-b98a8a1375a9"),
				phoneNumber: "+6281234567890",
				password:    "THISISAPASSWORD1234!",
				fullName:    "John Doe",
			},
			want: []error{
				errors.New("Password must have a capital letter, a number, a special character (non alpha numberic)"),
			},
		},
		{
			name: "should returns password errors no numeric letter",
			fields: fields{
				id:          uuid.MustParse("7ef62378-e51c-4d02-ba8b-b98a8a1375a9"),
				phoneNumber: "+6281234567890",
				password:    "ThisIsAPassword!",
				fullName:    "John Doe",
			},
			want: []error{
				errors.New("Password must have a capital letter, a number, a special character (non alpha numberic)"),
			},
		},
		{
			name: "should returns password errors no special character",
			fields: fields{
				id:          uuid.MustParse("7ef62378-e51c-4d02-ba8b-b98a8a1375a9"),
				phoneNumber: "+6281234567890",
				password:    "ThisIsAPassword1234",
				fullName:    "John Doe",
			},
			want: []error{
				errors.New("Password must have a capital letter, a number, a special character (non alpha numberic)"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &user{
				id:          tt.fields.id,
				phoneNumber: tt.fields.phoneNumber,
				password:    tt.fields.password,
				fullName:    tt.fields.fullName,
			}
			assert.Equalf(t, tt.want, u.Validate(), "Validate()")
		})
	}
}
