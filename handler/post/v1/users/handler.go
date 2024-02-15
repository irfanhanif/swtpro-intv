package users

import (
	"encoding/json"
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

func (h *handler) HandlePostV1Users(ctx handlerCtx.IContext) error {
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

	userID, _ := h.svc.RegisterNewUser(service.NewUser{
		PhoneNumber: req.PhoneNumber,
		Password:    req.Password,
		FullName:    req.FullName,
	})

	return ctx.JSON(http.StatusCreated, &generated.PostV1UsersResponse201{
		UserID: userID.String(),
	})
}
