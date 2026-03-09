package userHandler

import (
	"todolist/rest/middlewares"
	"todolist/service"
)

// Handler handles product-related HTTP requests.
type Handler struct {
	middlewares *middlewares.Middlewares
	userService service.UserService
}

// NewHandler constructs a new product Handler with the provided middlewares.
func NewHandler(middlewares *middlewares.Middlewares, userService service.UserService) *Handler {
	return &Handler{
		middlewares: middlewares,
		userService: userService,
	}
}
