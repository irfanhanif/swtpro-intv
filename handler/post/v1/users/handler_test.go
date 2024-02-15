package users

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/irfanhanif/swtpro-intv/generated"
	ctxMock "github.com/irfanhanif/swtpro-intv/handler/context/mock"
	"github.com/irfanhanif/swtpro-intv/service"
	serviceMock "github.com/irfanhanif/swtpro-intv/service/mock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

type readError struct{}

func (re *readError) Read(p []byte) (n int, err error) {
	return 0, errors.New("fail to read")
}

func TestHandlePostV1Users(t *testing.T) {
	type args struct {
		request *http.Request
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
				request: func() *http.Request {
					req := &generated.PostV1UsersRequest{
						FullName:    "John Doe",
						Password:    "ThisIsAPassword1234!",
						PhoneNumber: "+6281234567890",
					}
					b, _ := json.Marshal(req)
					return httptest.NewRequest("post", "/v1/users", bytes.NewReader(b))
				}(),
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
		"should return error " +
			"with status 400 bad request " +
			"and message body invalid" +
			"when body request unreadable": {
			args: args{
				request: httptest.NewRequest("post", "/v1/users", &readError{}),
			},
			expectContextJSON: &expectContextJSON{
				code: 400,
				body: &generated.PostV1UsersResponse400{
					Message: []string{
						"fail to read",
					},
				},
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
				mockCtx.EXPECT().Request().Return(test.args.request)
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
