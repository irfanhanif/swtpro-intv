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

	bodyBytes, _ := io.ReadAll(request.Body)
	defer request.Body.Close()

	req := &generated.PostV1UsersJSONRequestBody{}
	json.Unmarshal(bodyBytes, req)

	userID, _ := h.svc.RegisterNewUser(service.NewUser{
		PhoneNumber: req.PhoneNumber,
		Password:    req.Password,
		FullName:    req.FullName,
	})

	ctx.JSON(http.StatusCreated, &generated.PostV1UsersResponse201{
		UserID: userID.String(),
	})

	return nil
}
