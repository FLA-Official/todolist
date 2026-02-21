package projectHandler

import (
	"net/http"
	"todolist/rest/middlewares"
)

// RegisterRoutes registers project-related routes on the provided mux.
func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middlewares.Manager) {

	mux.Handle("GET /projects",
		manager.With(
			http.HandlerFunc(h.GetProjects),
		),
	) // declaring Route

	mux.Handle("POST /projects",
		manager.With(
			http.HandlerFunc(h.CreateProjectHandler),
		),
	) // declaring Route

	mux.Handle("GET /projects/{id}",
		manager.With(
			http.HandlerFunc(h.GetProjectByID),
		),
	) // declaring Route

	mux.Handle("PUT /projects/{id}",
		manager.With(
			http.HandlerFunc(h.UpdateProject),
			h.middlewares.AuthenticateJWT,
		),
	) // declaring Route

	mux.Handle("DELETE /projects/{id}",
		manager.With(
			http.HandlerFunc(h.DeleteProject),
			h.middlewares.AuthenticateJWT,
		),
	) // declaring Route

}
