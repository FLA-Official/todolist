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
			h.middlewares.AuthenticateJWT,
		),
	) // declaring Route

	mux.Handle("POST /projects",
		manager.With(
			http.HandlerFunc(h.CreateProjectHandler),
			h.middlewares.AuthenticateJWT,
		),
	) // declaring Route

	mux.Handle("GET /projects/{key}",
		manager.With(
			http.HandlerFunc(h.GetProjectByKey),
			h.middlewares.AuthenticateJWT,
		),
	) // declaring Route

	mux.Handle("PUT /projects/{key}",
		manager.With(
			http.HandlerFunc(h.UpdateProject),
			h.middlewares.AuthenticateJWT,
		),
	) // declaring Route

	mux.Handle("DELETE /projects/{key}",
		manager.With(
			http.HandlerFunc(h.DeleteProject),
			h.middlewares.AuthenticateJWT,
		),
	) // declaring Route

}
