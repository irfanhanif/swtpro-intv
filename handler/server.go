package handler

import (
	"github.com/irfanhanif/swtpro-intv/repository"
)

type Server struct {
	PostV1UsersHandler IHandlePostV1Users
}

type NewServerOptions struct {
	Repository repository.RepositoryInterface
}

func NewServer(opts NewServerOptions) *Server {
	return &Server{}
}
