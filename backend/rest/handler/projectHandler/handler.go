package projectHandler

import (
	"todolist/rest/middlewares"
	"todolist/service"
)

// Handler handles project-related HTTP requests.
type Handler struct {
	middlewares    *middlewares.Middlewares
	projectService service.ProjectService
}

// NewHandler constructs a new project Handler with the provided middlewares.
func NewHandler(middlewares *middlewares.Middlewares, projectService service.ProjectService) *Handler {
	return &Handler{
		middlewares:    middlewares,
		projectService: projectService,
	}
}
