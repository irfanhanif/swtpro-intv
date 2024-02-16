package handler

import (
	"github.com/irfanhanif/swtpro-intv/generated"
	"github.com/irfanhanif/swtpro-intv/handler/context"
	"github.com/irfanhanif/swtpro-intv/utils"
	"net/http"
	"strings"
)

type authenticatorMiddleware struct {
	jwt     utils.IValidateJWT
	handler IHandle
}

func NewAuthenticatorMiddleware(jwt utils.IValidateJWT, handler IHandle) *authenticatorMiddleware {
	return &authenticatorMiddleware{jwt: jwt, handler: handler}
}

func (am *authenticatorMiddleware) Handle(ctx context.IContext) error {
	request := ctx.Request()
	token := strings.TrimPrefix(request.Header.Get("Authorization"), "Bearer ")

	if token == "" {
		return ctx.JSON(http.StatusForbidden, &generated.Error{Error: "forbidden"})
	}

	userID, err := am.jwt.ValidateJWT(token)
	if err != nil {
		return ctx.JSON(http.StatusForbidden, &generated.Error{Error: "forbidden"})
	}

	ctx.Set("userID", userID)

	return am.handler.Handle(ctx)
}
