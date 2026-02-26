package projectMemberHandler

import (
	"todolist/repo"
	"todolist/rest/middlewares"
)

// Handler handles project-related HTTP requests.
type Handler struct {
	middlewares       *middlewares.Middlewares
	projectmemberrepo repo.ProjectMemberRepo
}

// NewHandler constructs a new project Handler with the provided middlewares.
func NewHandler(middlewares *middlewares.Middlewares, projectmemberrepo repo.ProjectMemberRepo) *Handler {
	return &Handler{
		middlewares:       middlewares,
		projectmemberrepo: projectmemberrepo,
	}
}
