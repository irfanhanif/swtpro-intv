package handler

import (
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/irfanhanif/swtpro-intv/generated"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"

	ctxMock "github.com/irfanhanif/swtpro-intv/handler/context/mock"
	handlerMock "github.com/irfanhanif/swtpro-intv/handler/mock"
	utilsMock "github.com/irfanhanif/swtpro-intv/utils/mock"
)

const TOKEN_EXAMPLE = "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDgxMzYyMTEsInVzZXJJRCI6ImZlNDEyMTg3LTMxYzAtNGNmOS1iYTYyLTVjOGM1YWQzYWZlMSJ9.UZeEoV9_K4NNulwhUOBsW9jB3c4mfO33eTPliTphl3O5mZ7YhnmU_gxshxlbFW18Aj28L5t1vfrsTFa13WBMH_vFud5ceb1yFEYy5TfGVSR68Z-kF26zz2Nnu18RVEkf3zeHf-yskFHqzEzptIZCrHiRJCc3CuxJLXVpAPlpzX79vkqeQ7lw9hARKtFaFci3epiaBmX_BPaYIMuTeVMB1EYHSGT3BTNDix6qdF4xzQlggZZmlqIv369RF-QEEcJPTe6AF8uw3E_buMNT6uHa65wHE60TE_rvdi42Ur84RwjeB-v7X8YXDSL6I9EnrRB-Rfbd-2vQyl86V8ltdrTABz-yb7pPQvWz-wYZ1IWvua3KtSHlV4-MSZLVQE5S1oh4Ujfve3QYi4W_2PdsFseOq9Uu4S-TNiyJr7tXig-dx-licrfFzCfHeB2qqoblMimEX2aMOeQaLghm4CtY-zdbI8sBgCJdNQPT3es-tL5Pojpx5CMyaxa2LWAlXCzuM5_EkPqwZ4E9S0iQ405pe-LrJKVuCiY-dM3g5GidYnyQJ9SvuqDExtkBHMctCk0edZGEmgHsBu6LpqavwPpVMM77JEvlqtVAofECr2s9YUqLH-nDcALmhKcalhdIUmgq1L7Pgc6Jyf4Xh2CLeXkDDDwnXVNvhmAIbBynChSkwCVw8hA"

func TestAuthenticatorMiddleware_Handle(t *testing.T) {
	type args struct {
		request *http.Request
	}
	type expectContextJSON struct {
		code int
		body interface{}

		returnError error
	}
	type expectValidateJWT struct {
		token string

		returnUserID uuid.UUID
		returnError  error
	}
	type expectHandle struct {
		isCalled bool
	}
	type expectContextSet struct {
		key string
		val interface{}
	}
	type test struct {
		args args

		expectValidateJWT *expectValidateJWT
		expectContextSet  *expectContextSet
		expectContextJSON *expectContextJSON
		expectHandle      *expectHandle

		wantError error
	}

	tests := map[string]test{
		"should forward to handler.go " +
			"when token is valid": {
			args: args{
				request: func() *http.Request {
					req := httptest.NewRequest("post", "/v1/users", nil)
					req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", TOKEN_EXAMPLE))
					return req
				}(),
			},
			expectValidateJWT: &expectValidateJWT{
				token: TOKEN_EXAMPLE,

				returnUserID: uuid.MustParse("c608ec88-1422-4574-9cda-d631ce2469b7"),
				returnError:  nil,
			},
			expectContextSet: &expectContextSet{
				key: "userID",
				val: uuid.MustParse("c608ec88-1422-4574-9cda-d631ce2469b7"),
			},
			expectHandle: &expectHandle{
				isCalled: true,
			},
			wantError: nil,
		},
		"should return 403 " +
			"when token is empty": {
			args: args{
				request: func() *http.Request {
					req := httptest.NewRequest("post", "/v1/users", nil)
					req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", ""))
					return req
				}(),
			},
			expectContextJSON: &expectContextJSON{
				code:        403,
				body:        &generated.Error{Error: "forbidden"},
				returnError: nil,
			},
			wantError: nil,
		},
		"should return forbidden " +
			"when validate returns error": {
			args: args{
				request: func() *http.Request {
					req := httptest.NewRequest("post", "/v1/users", nil)
					req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", TOKEN_EXAMPLE))
					return req
				}(),
			},
			expectValidateJWT: &expectValidateJWT{
				token: TOKEN_EXAMPLE,

				returnUserID: uuid.MustParse("c608ec88-1422-4574-9cda-d631ce2469b7"),
				returnError:  errors.New("validate error"),
			},
			expectContextJSON: &expectContextJSON{
				code:        403,
				body:        &generated.Error{Error: "forbidden"},
				returnError: nil,
			},
			wantError: nil,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			mockCtx := ctxMock.NewMockIContext(mockCtrl)
			mockJWT := utilsMock.NewMockIValidateJWT(mockCtrl)
			mockHandler := handlerMock.NewMockIHandle(mockCtrl)

			if test.args.request != nil {
				mockCtx.EXPECT().Request().Return(test.args.request).AnyTimes()
			}

			if test.expectContextJSON != nil {
				e := test.expectContextJSON
				mockCtx.EXPECT().JSON(e.code, e.body).Return(e.returnError)
			}

			if test.expectHandle != nil {
				mockHandler.EXPECT().Handle(mockCtx).Return(nil)
			}

			if test.expectContextSet != nil {
				e := test.expectContextSet
				mockCtx.EXPECT().Set(e.key, e.val)
			}

			if test.expectValidateJWT != nil {
				e := test.expectValidateJWT
				mockJWT.EXPECT().ValidateJWT(e.token).Return(e.returnUserID, e.returnError)
			}

			h := NewAuthenticatorMiddleware(mockJWT, mockHandler)
			actualErr := h.Handle(mockCtx)

			assert.Equal(t, test.wantError, actualErr)
		})
	}
}
