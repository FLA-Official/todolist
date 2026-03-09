package taskHandler

import (
	"todolist/rest/middlewares"
	"todolist/service"
)

// Handler handles product-related HTTP requests.
type Handler struct {
	middlewares          *middlewares.Middlewares
	taskService          service.TaskService
	projectService       service.ProjectService
	projectMemberService service.ProjectMemberService
}

// NewHandler constructs a new product Handler with the provided middlewares.
func NewHandler(middlewares *middlewares.Middlewares, taskService service.TaskService, projectService service.ProjectService, projectMemberService service.ProjectMemberService) *Handler {
	return &Handler{
		middlewares:          middlewares,
		taskService:          taskService,
		projectService:       projectService,
		projectMemberService: projectMemberService,
	}
}
