package taskHandler

import (
	"net/http"
	"todolist/rest/middlewares"
)

// RegisterRoutes registers product-related routes on the provided mux.
func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middlewares.Manager) {

	mux.Handle("GET /projects/{projectid}/tasks",
		manager.With(
			http.HandlerFunc(h.GetTasks),
			h.middlewares.AuthenticateJWT,
		),
	) // declaring Route

	mux.Handle("POST /projects/{projectid}/tasks",
		manager.With(
			http.HandlerFunc(h.CreateTaskHandler),
			h.middlewares.AuthenticateJWT,
		),
	) // declaring Route

	mux.Handle("GET /tasks/{taskid}",
		manager.With(
			http.HandlerFunc(h.GetTaskByID),
			h.middlewares.AuthenticateJWT,
		),
	) // declaring Route

	mux.Handle("PUT /tasks/{taskid}",
		manager.With(
			http.HandlerFunc(h.UpdateTask),
			h.middlewares.AuthenticateJWT,
		),
	) // declaring Route

	mux.Handle("DELETE /tasks/{taskid}",
		manager.With(
			http.HandlerFunc(h.DeleteTask),
			h.middlewares.AuthenticateJWT,
		),
	) // declaring Route

}
