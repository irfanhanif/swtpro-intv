package handler

import (
	"github.com/irfanhanif/swtpro-intv/repository"
)

type Server struct {
	Repository repository.RepositoryInterface

	postV1UsersHandler IHandlePostV1Users
}

type NewServerOptions struct {
	Repository repository.RepositoryInterface
}

func NewServer(opts NewServerOptions) *Server {
	return &Server{}
}
