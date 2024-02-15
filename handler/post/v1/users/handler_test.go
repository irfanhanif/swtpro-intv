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
		"should return error " +
			"with status 400 bad request " +
			"and message body invalid" +
			"when body request unmarshalable": {
			args: args{
				request: httptest.NewRequest("post", "/v1/users", bytes.NewReader([]byte(`blabla`))),
			},
			expectContextJSON: &expectContextJSON{
				code: 400,
				body: &generated.PostV1UsersResponse400{
					Message: []string{
						"failed to unmarshal",
					},
				},
				returnError: nil,
			},
			wantErr: nil,
		},
		"should return nil " +
			"with status 400 bad request " +
			"when required fields are empty": {
			args: args{
				request: func() *http.Request {
					req := &generated.PostV1UsersRequest{
						FullName:    "",
						Password:    "",
						PhoneNumber: "",
					}
					b, _ := json.Marshal(req)
					return httptest.NewRequest("post", "/v1/users", bytes.NewReader(b))
				}(),
			},
			expectContextJSON: &expectContextJSON{
				code: 400,
				body: &generated.PostV1UsersResponse400{
					Message: []string{
						"fullName cannot empty",
						"password cannot empty",
						"phoneNumber cannot empty",
					},
				},
				returnError: nil,
			},
			wantErr: nil,
		},
		"should return nil " +
			"with status 400 bad request " +
			"when fields doen't meed the required conditions": {
			args: args{
				request: func() *http.Request {
					req := &generated.PostV1UsersRequest{
						PhoneNumber: "+6281234567890",
						Password:    "ThisIsAPassword1234!",
						FullName:    "John Doe",
					}
					b, _ := json.Marshal(req)
					return httptest.NewRequest("post", "/v1/users", bytes.NewReader(b))
				}(),
			},
			expectContextJSON: &expectContextJSON{
				code: 400,
				body: &generated.PostV1UsersResponse400{
					Message: []string{
						"invalid fullName",
						"invalid password",
						"invalid phoneNumber",
					},
				},
				returnError: nil,
			},
			expectRegisterNewUser: &expectRegisterNewUser{
				newUser: service.NewUser{
					PhoneNumber: "+6281234567890",
					Password:    "ThisIsAPassword1234!",
					FullName:    "John Doe",
				},
				returnUUID: uuid.Nil,
				returnError: &service.ErrFields{
					Errs: []string{
						"invalid fullName",
						"invalid password",
						"invalid phoneNumber",
					},
				},
			},
			wantErr: nil,
		},
		"should return nil " +
			"with status 409 conflict " +
			"when phone number already exists": {
			args: args{
				request: func() *http.Request {
					req := &generated.PostV1UsersRequest{
						PhoneNumber: "+6281234567890",
						Password:    "ThisIsAPassword1234!",
						FullName:    "John Doe",
					}
					b, _ := json.Marshal(req)
					return httptest.NewRequest("post", "/v1/users", bytes.NewReader(b))
				}(),
			},
			expectContextJSON: &expectContextJSON{
				code:        409,
				body:        &generated.Error{Error: "given phone number already exists"},
				returnError: nil,
			},
			expectRegisterNewUser: &expectRegisterNewUser{
				newUser: service.NewUser{
					PhoneNumber: "+6281234567890",
					Password:    "ThisIsAPassword1234!",
					FullName:    "John Doe",
				},
				returnUUID:  uuid.Nil,
				returnError: service.ErrPhoneNumberConflict,
			},
			wantErr: nil,
		},
		"should return nil " +
			"with status 500 internal " +
			"when register new user returns error": {
			args: args{
				request: func() *http.Request {
					req := &generated.PostV1UsersRequest{
						PhoneNumber: "+6281234567890",
						Password:    "ThisIsAPassword1234!",
						FullName:    "John Doe",
					}
					b, _ := json.Marshal(req)
					return httptest.NewRequest("post", "/v1/users", bytes.NewReader(b))
				}(),
			},
			expectContextJSON: &expectContextJSON{
				code:        500,
				body:        &generated.Error{Error: "an error happened"},
				returnError: nil,
			},
			expectRegisterNewUser: &expectRegisterNewUser{
				newUser: service.NewUser{
					PhoneNumber: "+6281234567890",
					Password:    "ThisIsAPassword1234!",
					FullName:    "John Doe",
				},
				returnUUID:  uuid.Nil,
				returnError: errors.New("an error happened"),
			},
			wantErr: nil,
		},
		"should return error from contest json response " +
			"contest JSON() returns error": {
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
				returnError: errors.New("something bad happened"),
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
			wantErr: errors.New("something bad happened"),
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
