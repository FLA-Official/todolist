package taskHandler

import (
	"todolist/repo"
	"todolist/rest/middlewares"
)

// Handler handles product-related HTTP requests.
type Handler struct {
	middlewares *middlewares.Middlewares
	taskrepo    repo.TaskRepo
}

// NewHandler constructs a new product Handler with the provided middlewares.
func NewHandler(middlewares *middlewares.Middlewares, taskrepo repo.TaskRepo) *Handler {
	return &Handler{
		middlewares: middlewares,
		taskrepo:    taskrepo,
	}
}
