package userHandler

import (
	"todolist/repo"
	"todolist/rest/middlewares"
)

// Handler handles product-related HTTP requests.
type Handler struct {
	middlewares *middlewares.Middlewares
	userrepo    repo.UserRepo
}

// NewHandler constructs a new product Handler with the provided middlewares.
func NewHandler(middlewares *middlewares.Middlewares, userrepo repo.UserRepo) *Handler {
	return &Handler{
		middlewares: middlewares,
		userrepo:    userrepo,
	}
}
