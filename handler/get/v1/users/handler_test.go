package users

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/irfanhanif/swtpro-intv/entity"
	ctxMock "github.com/irfanhanif/swtpro-intv/handler/context/mock"
	"github.com/irfanhanif/swtpro-intv/service"
	"github.com/stretchr/testify/assert"

	"github.com/irfanhanif/swtpro-intv/generated"
	"net/http"
	"net/http/httptest"
	"testing"

	mockEntity "github.com/irfanhanif/swtpro-intv/entity/mock"
	mockService "github.com/irfanhanif/swtpro-intv/service/mock"
)

func Test_handler_Handle(t *testing.T) {
	type args struct {
		request *http.Request
	}
	type expectContextJSON struct {
		code int
		body interface{}

		returnError error
	}
	type expectContextParam struct {
		name        string
		returnValue string
	}
	type expectGetUserByID struct {
		id uuid.UUID

		returnUser  entity.IUser
		returnError error
	}
	type test struct {
		args args

		expectContextJSON  *expectContextJSON
		expectGetUserByID  *expectGetUserByID
		expectContextParam *expectContextParam

		wantErr error
	}

	tests := map[string]func(*gomock.Controller) test{
		"should response with 200 ok " +
			"when successfully get user data": func(mockCtrl *gomock.Controller) test {
			mockUser := mockEntity.NewMockIUser(mockCtrl)
			mockUser.EXPECT().FullName().Return("John Doe").AnyTimes()
			mockUser.EXPECT().PhoneNumber().Return("+628123456789").AnyTimes()

			return test{
				args: args{
					request: httptest.NewRequest("post", "/v1/users/{userID}", nil),
				},
				expectGetUserByID: &expectGetUserByID{
					id:          uuid.MustParse("d87a947c-6023-4e5b-8b95-1e21aafad630"),
					returnUser:  mockUser,
					returnError: nil,
				},
				expectContextJSON: &expectContextJSON{
					code: 200,
					body: &generated.GetV1Users{
						FullName:    "John Doe",
						PhoneNumber: "+628123456789",
					},
				},
				expectContextParam: &expectContextParam{
					name:        "userID",
					returnValue: "d87a947c-6023-4e5b-8b95-1e21aafad630",
				},
				wantErr: nil,
			}
		},
		"should response with 400 bad request " +
			"when user id is invalid": func(mockCtrl *gomock.Controller) test {
			mockUser := mockEntity.NewMockIUser(mockCtrl)
			mockUser.EXPECT().FullName().Return("John Doe").AnyTimes()
			mockUser.EXPECT().PhoneNumber().Return("+628123456789").AnyTimes()

			return test{
				args: args{
					request: httptest.NewRequest("post", "/v1/users/{userID}", nil),
				},
				expectContextJSON: &expectContextJSON{
					code: 400,
					body: &generated.Error{Error: "invalid user id"},
				},
				expectContextParam: &expectContextParam{
					name:        "userID",
					returnValue: "",
				},
				wantErr: nil,
			}
		},
		"should response with 404 not found " +
			"when user is not found": func(mockCtrl *gomock.Controller) test {
			mockUser := mockEntity.NewMockIUser(mockCtrl)
			mockUser.EXPECT().FullName().Return("John Doe").AnyTimes()
			mockUser.EXPECT().PhoneNumber().Return("+628123456789").AnyTimes()

			return test{
				args: args{
					request: httptest.NewRequest("post", "/v1/users/{userID}", nil),
				},
				expectGetUserByID: &expectGetUserByID{
					id:          uuid.MustParse("d87a947c-6023-4e5b-8b95-1e21aafad630"),
					returnUser:  mockUser,
					returnError: service.ErrNotFound,
				},
				expectContextJSON: &expectContextJSON{
					code: 404,
					body: &generated.Error{Error: "data not found"},
				},
				expectContextParam: &expectContextParam{
					name:        "userID",
					returnValue: "d87a947c-6023-4e5b-8b95-1e21aafad630",
				},
				wantErr: nil,
			}
		},
		"should response with 500 internal " +
			"when handling any errors": func(mockCtrl *gomock.Controller) test {
			mockUser := mockEntity.NewMockIUser(mockCtrl)
			mockUser.EXPECT().FullName().Return("John Doe").AnyTimes()
			mockUser.EXPECT().PhoneNumber().Return("+628123456789").AnyTimes()

			return test{
				args: args{
					request: httptest.NewRequest("post", "/v1/users/{userID}", nil),
				},
				expectGetUserByID: &expectGetUserByID{
					id:          uuid.MustParse("d87a947c-6023-4e5b-8b95-1e21aafad630"),
					returnUser:  mockUser,
					returnError: errors.New("oops"),
				},
				expectContextJSON: &expectContextJSON{
					code: 500,
					body: &generated.Error{Error: "oops"},
				},
				expectContextParam: &expectContextParam{
					name:        "userID",
					returnValue: "d87a947c-6023-4e5b-8b95-1e21aafad630",
				},
				wantErr: nil,
			}
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			mockGetUserByID := mockService.NewMockIGetUserByID(mockCtrl)
			mockCtx := ctxMock.NewMockIContext(mockCtrl)

			tt := test(mockCtrl)

			if tt.args.request != nil {
				mockCtx.EXPECT().Request().Return(tt.args.request)
			}

			if tt.expectContextJSON != nil {
				e := tt.expectContextJSON
				mockCtx.EXPECT().JSON(e.code, e.body).Return(e.returnError)
			}

			if tt.expectContextParam != nil {
				e := tt.expectContextParam
				mockCtx.EXPECT().Param(e.name).Return(e.returnValue)
			}

			if tt.expectGetUserByID != nil {
				e := tt.expectGetUserByID
				mockGetUserByID.EXPECT().GetUserByID(gomock.Any(), e.id).Return(e.returnUser, e.returnError)
			}

			h := NewHandler(mockGetUserByID)
			actualErr := h.Handle(mockCtx)

			assert.Equal(t, tt.wantErr, actualErr)
		})
	}
}
