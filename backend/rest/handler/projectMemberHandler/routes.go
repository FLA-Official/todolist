package projectMemberHandler

import (
	"net/http"
	"todolist/rest/middlewares"
)

// RegisterRoutes registers routes for ProjectMember
func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middlewares.Manager) {

	// List all members of a project
	mux.Handle("GET /projects/{projectId}/members",
		manager.With(
			http.HandlerFunc(h.GetMembersByProject),
			h.middlewares.AuthenticateJWT,
		),
	)

	// Add a member to a project
	mux.Handle("POST /projects/{projectId}/members",
		manager.With(
			http.HandlerFunc(h.AddMemberHandler),
			h.middlewares.AuthenticateJWT,
		),
	)

	// Update role of a member
	mux.Handle("PUT /projects/{projectId}/members/{userId}",
		manager.With(
			http.HandlerFunc(h.UpdateMemberRole),
			h.middlewares.AuthenticateJWT,
		),
	)

	// Remove a member
	mux.Handle("DELETE /projects/{projectId}/members/{userId}",
		manager.With(
			http.HandlerFunc(h.RemoveMember),
			h.middlewares.AuthenticateJWT,
		),
	)
}
