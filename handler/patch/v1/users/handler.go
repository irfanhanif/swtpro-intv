package users

import (
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"github.com/irfanhanif/swtpro-intv/generated"
	handlerCtx "github.com/irfanhanif/swtpro-intv/handler/context"
	"github.com/irfanhanif/swtpro-intv/service"
	"github.com/irfanhanif/swtpro-intv/valueobj"
	"io"
	"net/http"
)

type handler struct {
	svc service.IUpdateUserByID
}

func NewHandler(svc service.IUpdateUserByID) *handler {
	return &handler{svc: svc}
}

func (h *handler) Handle(ctx handlerCtx.IContext) error {
	request := ctx.Request()

	userID, err := uuid.Parse(ctx.Param("userID"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, &generated.PatchV1UsersResponse400{Message: []string{"invalid user id"}})
	}

	bodyBytes, err := io.ReadAll(request.Body)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, &generated.PatchV1UsersResponse400{Message: []string{"failed to read body"}})
	}
	defer request.Body.Close()

	req := &generated.PatchV1UsersRequest{}
	if err := json.Unmarshal(bodyBytes, req); err != nil {
		return ctx.JSON(http.StatusBadRequest, &generated.PatchV1UsersResponse400{Message: []string{"failed to unmarshal"}})
	}

	err = h.svc.UpdateUserByID(request.Context(), userID, valueobj.UserUpdateData{
		PhoneNumber: req.PhoneNumber,
		FullName:    req.FullName,
	})
	if errors.Is(err, service.ErrNotFound) {
		return ctx.JSON(http.StatusNotFound, &generated.Error{Error: "user not found"})
	}
	if errors.Is(err, service.ErrPhoneNumberConflict) {
		return ctx.JSON(http.StatusConflict, &generated.Error{Error: "phone number already exist"})
	}
	if fieldErrors, ok := err.(*service.ErrFields); ok {
		return ctx.JSON(http.StatusBadRequest, &generated.PatchV1UsersResponse400{
			Message: fieldErrors.Errs,
		})
	}
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, &generated.Error{Error: err.Error()})
	}

	return ctx.NoContent(http.StatusNoContent)
}
