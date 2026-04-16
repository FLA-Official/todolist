package projectMemberHandler

import (
	"todolist/rest/middlewares"
	"todolist/service"
)

// Handler handles project-related HTTP requests.
type Handler struct {
	middlewares          *middlewares.Middlewares
	projectService       service.ProjectService
	projectMemberService service.ProjectMemberService
}

// NewHandler constructs a new project Handler with the provided middlewares.
func NewHandler(middlewares *middlewares.Middlewares, projectService service.ProjectService, projectMemberService service.ProjectMemberService) *Handler {
	return &Handler{
		middlewares:          middlewares,
		projectService:       projectService,
		projectMemberService: projectMemberService,
	}
}
