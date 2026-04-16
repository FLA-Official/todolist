package taskHandler

import (
	"net/http"
	"todolist/rest/middlewares"
)

// RegisterRoutes registers product-related routes on the provided mux.
func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middlewares.Manager) {

	mux.Handle("GET /projects/{projectkey}/tasks",
		manager.With(
			http.HandlerFunc(h.GetProjectTasks),
			h.middlewares.AuthenticateJWT,
		),
	) // declaring Route

	mux.Handle("POST /projects/{projectkey}/tasks",
		manager.With(
			http.HandlerFunc(h.CreateTaskHandler),
			h.middlewares.AuthenticateJWT,
		),
	) // declaring Route

	mux.Handle("PUT /projects/{projectkey}/tasks/{taskid}",
		manager.With(
			http.HandlerFunc(h.UpdateTask),
			h.middlewares.AuthenticateJWT,
		),
	) // declaring Route

	mux.Handle("DELETE /projects/{projectkey}/tasks/{taskid}",
		manager.With(
			http.HandlerFunc(h.DeleteTask),
			h.middlewares.AuthenticateJWT,
		),
	) // declaring Route

}
