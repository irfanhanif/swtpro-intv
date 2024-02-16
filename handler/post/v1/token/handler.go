package token

import (
	"encoding/json"
	"errors"
	"github.com/irfanhanif/swtpro-intv/generated"
	handlerCtx "github.com/irfanhanif/swtpro-intv/handler/context"
	"github.com/irfanhanif/swtpro-intv/service"
	"io"
	"net/http"
)

type handler struct {
	svc service.IGenerateToken
}

func NewHandler(svc service.IGenerateToken) *handler {
	return &handler{svc: svc}
}

func (h *handler) Handle(ctx handlerCtx.IContext) error {
	request := ctx.Request()

	bodyBytes, err := io.ReadAll(request.Body)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, &generated.Error{Error: err.Error()})
	}
	defer request.Body.Close()

	req := &generated.PostV1TokenRequest{}
	if err := json.Unmarshal(bodyBytes, req); err != nil {
		return ctx.JSON(http.StatusBadRequest, &generated.Error{Error: "failed to unmarshal"})
	}

	if req.PhoneNumber == "" {
		return ctx.JSON(http.StatusBadRequest, &generated.Error{Error: "phone number cannot be empty"})
	}
	if req.Password == "" {
		return ctx.JSON(http.StatusBadRequest, &generated.Error{Error: "password cannot be empty"})
	}

	token, userID, err := h.svc.GenerateToken(request.Context(), req.PhoneNumber, req.Password)
	if errors.Is(err, service.ErrLoginFailed) {
		return ctx.JSON(http.StatusBadRequest, &generated.Error{Error: err.Error()})
	}
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, &generated.Error{Error: err.Error()})
	}

	return ctx.JSON(http.StatusCreated, &generated.PostV1TokenResponse201{
		Token:  token,
		UserID: userID.String(),
	})
}
