package users

import (
	"errors"
	"github.com/google/uuid"
	"github.com/irfanhanif/swtpro-intv/generated"
	handlerCtx "github.com/irfanhanif/swtpro-intv/handler/context"
	"github.com/irfanhanif/swtpro-intv/service"
	"net/http"
)

type handler struct {
	svc service.IGetUserByID
}

func NewHandler(svc service.IGetUserByID) *handler {
	return &handler{svc: svc}
}

func (h *handler) Handle(ctx handlerCtx.IContext) error {
	request := ctx.Request()

	userID, err := uuid.Parse(ctx.Param("userID"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, &generated.Error{Error: "invalid user id"})
	}

	user, err := h.svc.GetUserByID(request.Context(), userID)
	if errors.Is(err, service.ErrNotFound) {
		return ctx.JSON(http.StatusNotFound, &generated.Error{Error: "data not found"})
	}
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, &generated.Error{Error: err.Error()})
	}

	return ctx.JSON(http.StatusOK, &generated.GetV1Users{
		FullName:    user.FullName(),
		PhoneNumber: user.PhoneNumber(),
	})
}
