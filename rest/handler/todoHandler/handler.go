package todohandler

import (
	"todolist/rest/middleware"
)

type Handler struct {
	middlewares *middleware.Middleware
}

func Newhandler(middlewares *middleware.Middleware) *Handler {
	return &Handler{
		middlewares: middlewares,
	}
}
