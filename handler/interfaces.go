package handler

import "github.com/irfanhanif/swtpro-intv/handler/context"

//go:generate mockgen -source=interfaces.go -destination=mock/interfaces.go -package=mock

type IHandle interface {
	Handle(ctx context.IContext) error
}
