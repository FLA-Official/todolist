package todoHandler

import (
	middleware "todolist/rest/middlewares"
)

type Handler struct {
	middlewares *middleware.Middlewares
}

func Newhandler(middlewares *middleware.Middlewares) *Handler {
	return &Handler{
		middlewares: middlewares,
	}
}
