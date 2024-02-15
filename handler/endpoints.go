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
	return s.postV1UsersHandler.HandlePostV1Users(ctx)
}

func (s *Server) PostV1Token(ctx echo.Context) error {
	return nil
}

func (s *Server) GetV1UsersUserIDProfile(ctx echo.Context, userID string) error {
	return nil
}

func (s *Server) PatchV1UsersUserIDProfile(ctx echo.Context, userID string) error {
	return nil
}
