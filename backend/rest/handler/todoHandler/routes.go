package todoHandler

import (
	"net/http"
	"todolist/rest/middlewares"
)

// RegisterRoutes registers product-related routes on the provided mux.
func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middlewares.Manager) {

	mux.Handle("GET /tasks",
		manager.With(
			http.HandlerFunc(h.GetTasks),
		),
	) // declaring Route

	mux.Handle("POST /tasks",
		manager.With(
			http.HandlerFunc(h.CreateTask),
		),
	) // declaring Route

	mux.Handle("GET /tasks/{id}",
		manager.With(
			http.HandlerFunc(h.GetTaskByID),
		),
	) // declaring Route

	mux.Handle("PUT /tasks/{id}",
		manager.With(
			http.HandlerFunc(h.UpdateTask),
			// h.middlewares.AuthenticateJWT,
		),
	) // declaring Route

	mux.Handle("DELETE /tasks/{id}",
		manager.With(
			http.HandlerFunc(h.DeleteTask),
			// h.middlewares.AuthenticateJWT,
		),
	) // declaring Route

}
