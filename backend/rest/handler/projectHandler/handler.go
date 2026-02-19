package projectHandler

import (
	"todolist/repo"
	"todolist/rest/middlewares"
)

// Handler handles project-related HTTP requests.
type Handler struct {
	middlewares *middlewares.Middlewares
	projectrepo repo.ProjectRepo
}

// NewHandler constructs a new project Handler with the provided middlewares.
func NewHandler(middlewares *middlewares.Middlewares, projectrepo repo.ProjectRepo) *Handler {
	return &Handler{
		middlewares: middlewares,
		projectrepo: projectrepo,
	}
}
