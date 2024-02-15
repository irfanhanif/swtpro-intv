package token

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

func TestHandler_HandlePostV1Token(t *testing.T) {
	type args struct {
		request *http.Request
	}
	type expectContextJSON struct {
		code int
		body interface{}

		returnError error
	}
	type expectGenerateToken struct {
		phoneNumber string
		password    string

		returnToken  string
		returnUserID uuid.UUID
		returnError  error
	}
	type test struct {
		args args

		expectContextJSON   *expectContextJSON
		expectGenerateToken *expectGenerateToken

		wantErr error
	}

	tests := map[string]test{
		"should return nil " +
			"with status 201 created " +
			"and with user id and token " +
			"when successfully generate token": {
			args: args{
				request: func() *http.Request {
					req := &generated.PostV1TokenRequest{
						Password:    "ThisIsAPassword1234!",
						PhoneNumber: "+6281234567890",
					}
					b, _ := json.Marshal(req)
					return httptest.NewRequest("post", "/v1/token", bytes.NewReader(b))
				}(),
			},
			expectContextJSON: &expectContextJSON{
				code: 201,
				body: &generated.PostV1TokenResponse201{
					Token:  "this-is-a-token",
					UserID: "8e89a19b-e5ea-4c3a-93ea-62b99e23d5a4",
				},
			},
			expectGenerateToken: &expectGenerateToken{
				phoneNumber: "+6281234567890",
				password:    "ThisIsAPassword1234!",

				returnToken:  "this-is-a-token",
				returnUserID: uuid.MustParse("8e89a19b-e5ea-4c3a-93ea-62b99e23d5a4"),
				returnError:  nil,
			},
			wantErr: nil,
		},
		"should return nil " +
			"with status 400 bad request " +
			"and message body invalid " +
			"when body request unreadable": {
			args: args{
				request: httptest.NewRequest("post", "/v1/token", &readError{}),
			},
			expectContextJSON: &expectContextJSON{
				code:        400,
				body:        &generated.Error{Error: "fail to read"},
				returnError: nil,
			},
			wantErr: nil,
		},
		"should return nil " +
			"with status 400 bad request " +
			"and message body invalid " +
			"when body request unmarshalable": {
			args: args{
				request: httptest.NewRequest("post", "/v1/token", bytes.NewReader([]byte(`blabla`))),
			},
			expectContextJSON: &expectContextJSON{
				code:        400,
				body:        &generated.Error{Error: "failed to unmarshal"},
				returnError: nil,
			},
			wantErr: nil,
		},
		"should return nil " +
			"with status 400 bad request " +
			"when phone number are emtpy": {
			args: args{
				request: func() *http.Request {
					req := &generated.PostV1TokenRequest{
						Password:    "ThisIsAPassword1234!",
						PhoneNumber: "",
					}
					b, _ := json.Marshal(req)
					return httptest.NewRequest("post", "/v1/token", bytes.NewReader(b))
				}(),
			},
			expectContextJSON: &expectContextJSON{
				code:        400,
				body:        &generated.Error{Error: "phone number cannot be empty"},
				returnError: nil,
			},
			wantErr: nil,
		},
		"should return nil " +
			"with status 400 bad request " +
			"when password are emtpy": {
			args: args{
				request: func() *http.Request {
					req := &generated.PostV1TokenRequest{
						Password:    "",
						PhoneNumber: "+628123456789",
					}
					b, _ := json.Marshal(req)
					return httptest.NewRequest("post", "/v1/token", bytes.NewReader(b))
				}(),
			},
			expectContextJSON: &expectContextJSON{
				code:        400,
				body:        &generated.Error{Error: "password cannot be empty"},
				returnError: nil,
			},
			wantErr: nil,
		},
		"should return nil " +
			"with status 400 bad request " +
			"when login is failed": {
			args: args{
				request: func() *http.Request {
					req := &generated.PostV1TokenRequest{
						Password:    "ThisIsAPassword1234!",
						PhoneNumber: "+6281234567890",
					}
					b, _ := json.Marshal(req)
					return httptest.NewRequest("post", "/v1/token", bytes.NewReader(b))
				}(),
			},
			expectContextJSON: &expectContextJSON{
				code:        400,
				body:        &generated.Error{Error: "login failed"},
				returnError: nil,
			},
			expectGenerateToken: &expectGenerateToken{
				phoneNumber: "+6281234567890",
				password:    "ThisIsAPassword1234!",

				returnToken:  "",
				returnUserID: uuid.Nil,
				returnError:  service.ErrLoginFailed,
			},
			wantErr: nil,
		},
		"should return nil " +
			"with status 500 internal server error " +
			"when generate token returns error": {
			args: args{
				request: func() *http.Request {
					req := &generated.PostV1TokenRequest{
						Password:    "ThisIsAPassword1234!",
						PhoneNumber: "+6281234567890",
					}
					b, _ := json.Marshal(req)
					return httptest.NewRequest("post", "/v1/token", bytes.NewReader(b))
				}(),
			},
			expectContextJSON: &expectContextJSON{
				code:        500,
				body:        &generated.Error{Error: "something bad happened"},
				returnError: nil,
			},
			expectGenerateToken: &expectGenerateToken{
				phoneNumber: "+6281234567890",
				password:    "ThisIsAPassword1234!",

				returnToken:  "",
				returnUserID: uuid.Nil,
				returnError:  errors.New("something bad happened"),
			},
			wantErr: nil,
		},
		"should return error " +
			"when failed to response": {
			args: args{
				request: func() *http.Request {
					req := &generated.PostV1TokenRequest{
						Password:    "ThisIsAPassword1234!",
						PhoneNumber: "+6281234567890",
					}
					b, _ := json.Marshal(req)
					return httptest.NewRequest("post", "/v1/token", bytes.NewReader(b))
				}(),
			},
			expectContextJSON: &expectContextJSON{
				code: 201,
				body: &generated.PostV1TokenResponse201{
					Token:  "this-is-a-token",
					UserID: "8e89a19b-e5ea-4c3a-93ea-62b99e23d5a4",
				},
				returnError: errors.New("oops"),
			},
			expectGenerateToken: &expectGenerateToken{
				phoneNumber: "+6281234567890",
				password:    "ThisIsAPassword1234!",

				returnToken:  "this-is-a-token",
				returnUserID: uuid.MustParse("8e89a19b-e5ea-4c3a-93ea-62b99e23d5a4"),
				returnError:  nil,
			},
			wantErr: errors.New("oops"),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			mockCtx := ctxMock.NewMockIContext(mockCtrl)
			mockGenerateToken := serviceMock.NewMockIGenerateToken(mockCtrl)

			if test.args.request != nil {
				mockCtx.EXPECT().Request().Return(test.args.request)
			}

			if test.expectContextJSON != nil {
				e := test.expectContextJSON
				mockCtx.EXPECT().JSON(e.code, e.body).Return(e.returnError)
			}

			if test.expectGenerateToken != nil {
				e := test.expectGenerateToken
				mockGenerateToken.EXPECT().GenerateToken(gomock.Any(), e.phoneNumber, e.password).Return(e.returnToken, e.returnUserID, e.returnError)
			}

			h := NewHandler(mockGenerateToken)
			actualErr := h.HandlePostV1Token(mockCtx)

			assert.Equal(t, test.wantErr, actualErr)
		})
	}
}
