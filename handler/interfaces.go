package handler

import "github.com/irfanhanif/swtpro-intv/handler/context"

type IHandlePostV1Users interface {
	HandlePostV1Users(ctx context.IContext) error
}
