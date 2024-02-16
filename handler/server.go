package handler

import (
	"github.com/irfanhanif/swtpro-intv/repository"
)

type Server struct {
	PostV1UsersHandler IHandle
	PostV1TokenHandler IHandle
	GetV1UsersHandler  IHandle
}

type NewServerOptions struct {
	Repository repository.RepositoryInterface
}

func NewServer(opts NewServerOptions) *Server {
	return &Server{}
}
