package users

import (
	"bytes"
	"encoding/json"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/irfanhanif/swtpro-intv/generated"
	ctxMock "github.com/irfanhanif/swtpro-intv/handler/context/mock"
	"github.com/irfanhanif/swtpro-intv/service"
	serviceMock "github.com/irfanhanif/swtpro-intv/service/mock"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestHandlePostV1Users(t *testing.T) {
	type args struct {
		request *generated.PostV1UsersRequest
	}
	type expectContextJSON struct {
		code int
		body interface{}

		returnError error
	}
	type expectRegisterNewUser struct {
		newUser service.NewUser

		returnUUID  uuid.UUID
		returnError error
	}

	type test struct {
		args args

		expectContextJSON     *expectContextJSON
		expectRegisterNewUser *expectRegisterNewUser

		wantErr error
	}

	tests := map[string]test{
		"should return nil " +
			"with status 201 created " +
			"and user id" +
			"when new user successfully registered": {
			args: args{
				request: &generated.PostV1UsersRequest{
					FullName:    "John Doe",
					Password:    "ThisIsAPassword1234!",
					PhoneNumber: "+6281234567890",
				},
			},
			expectContextJSON: &expectContextJSON{
				code: 201,
				body: &generated.PostV1UsersResponse201{
					UserID: "ce29af20-aa68-4ca5-8ef4-f1199704507b",
				},
				returnError: nil,
			},
			expectRegisterNewUser: &expectRegisterNewUser{
				newUser: service.NewUser{
					PhoneNumber: "+6281234567890",
					Password:    "ThisIsAPassword1234!",
					FullName:    "John Doe",
				},
				returnUUID:  uuid.MustParse("ce29af20-aa68-4ca5-8ef4-f1199704507b"),
				returnError: nil,
			},
			wantErr: nil,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			mockCtx := ctxMock.NewMockIContext(mockCtrl)
			mockRegisterNewUser := serviceMock.NewMockIRegisterNewUser(mockCtrl)

			if test.args.request != nil {
				req := test.args.request
				b, _ := json.Marshal(req)
				mockCtx.EXPECT().Request().Return(httptest.NewRequest("post", "/v1/users", bytes.NewReader(b)))
			}

			if test.expectContextJSON != nil {
				e := test.expectContextJSON
				mockCtx.EXPECT().JSON(e.code, e.body).Return(e.returnError)
			}

			if test.expectRegisterNewUser != nil {
				e := test.expectRegisterNewUser
				mockRegisterNewUser.EXPECT().RegisterNewUser(e.newUser).Return(e.returnUUID, e.returnError)
			}

			h := NewHandler(mockRegisterNewUser)
			actualErr := h.HandlePostV1Users(mockCtx)

			assert.Equal(t, test.wantErr, actualErr)
		})
	}
}
