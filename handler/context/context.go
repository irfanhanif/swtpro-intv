package context

import "net/http"

//go:generate mockgen -source=context.go -destination=mock/context.go -package=mock

type IContext interface {
	Request() *http.Request
	JSON(code int, i interface{}) error
}
