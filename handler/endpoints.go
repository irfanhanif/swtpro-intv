package handler

import (
	"github.com/labstack/echo/v4"
)

// This is just a test endpoint to get you started. Please delete this endpoint.
// (GET /hello)
//func (s *Server) Hello(ctx echo.Context, params generated.HelloParams) error {
//
//	var resp generated.HelloResponse
//	resp.Message = fmt.Sprintf("Hello User %d", params.Id)
//	return ctx.JSON(http.StatusOK, resp)
//}

func (s *Server) PostV1Users(ctx echo.Context) error {
	return s.PostV1UsersHandler.Handle(ctx)
}

func (s *Server) PostV1Token(ctx echo.Context) error {
	return s.PostV1TokenHandler.Handle(ctx)
}

func (s *Server) GetV1UsersUserID(ctx echo.Context, userID string) error {
	return s.GetV1UsersHandler.Handle(ctx)
}

func (s *Server) PatchV1UsersUserID(ctx echo.Context, userID string) error {
	return nil
}
