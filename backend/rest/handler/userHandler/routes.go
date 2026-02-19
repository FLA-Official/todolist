package userHandler

import (
	"net/http"
	"todolist/rest/middlewares"
)

// RegisterRoutes registers user-related routes on the provided mux.
func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middlewares.Manager) {

	mux.Handle("GET /users",
		manager.With(
			http.HandlerFunc(h.GetUsers),
		),
	) // declaring Route

	mux.Handle("POST /users",
		manager.With(
			http.HandlerFunc(h.CreateUserHandler),
		),
	) // declaring Route

	mux.Handle("GET /users/{id}",
		manager.With(
			http.HandlerFunc(h.GetUserByID),
		),
	) // declaring Route

	mux.Handle("PUT /users/{id}",
		manager.With(
			http.HandlerFunc(h.UpdateUser),
			h.middlewares.AuthenticateJWT,
		),
	) // declaring Route

	mux.Handle("DELETE /users/{id}",
		manager.With(
			http.HandlerFunc(h.DeleteUser),
			h.middlewares.AuthenticateJWT,
		),
	) // declaring Route

}
