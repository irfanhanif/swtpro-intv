package users

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
	svc service.IRegisterNewUser
}

func NewHandler(svc service.IRegisterNewUser) *handler {
	return &handler{
		svc: svc,
	}
}

func (h *handler) Handle(ctx handlerCtx.IContext) error {
	request := ctx.Request()

	bodyBytes, err := io.ReadAll(request.Body)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, &generated.PostV1UsersResponse400{
			Message: []string{
				err.Error(),
			},
		})
	}
	defer request.Body.Close()

	req := &generated.PostV1UsersJSONRequestBody{}
	if err := json.Unmarshal(bodyBytes, req); err != nil {
		return ctx.JSON(http.StatusBadRequest, &generated.PostV1UsersResponse400{
			Message: []string{
				"failed to unmarshal",
			},
		})
	}

	var errs []string
	if req.FullName == "" {
		errs = append(errs, "fullName cannot empty")
	}
	if req.Password == "" {
		errs = append(errs, "password cannot empty")
	}
	if req.PhoneNumber == "" {
		errs = append(errs, "phoneNumber cannot empty")
	}

	if len(errs) > 0 {
		return ctx.JSON(http.StatusBadRequest, &generated.PostV1UsersResponse400{
			Message: errs,
		})
	}

	userID, err := h.svc.RegisterNewUser(request.Context(), service.NewUser{
		PhoneNumber: req.PhoneNumber,
		Password:    req.Password,
		FullName:    req.FullName,
	})
	if fieldErrors, ok := err.(*service.ErrFields); ok {
		return ctx.JSON(http.StatusBadRequest, &generated.PostV1UsersResponse400{
			Message: fieldErrors.Errs,
		})
	}
	if errors.Is(err, service.ErrPhoneNumberConflict) {
		return ctx.JSON(http.StatusConflict, &generated.Error{Error: err.Error()})
	}
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, &generated.Error{Error: err.Error()})
	}

	return ctx.JSON(http.StatusCreated, &generated.PostV1UsersResponse201{
		UserID: userID.String(),
	})
}
