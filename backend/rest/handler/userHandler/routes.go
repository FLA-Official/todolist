package userHandler

import (
	"net/http"
	"todolist/rest/middlewares"
)

// RegisterRoutes registers user-related routes on the provided mux.
func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middlewares.Manager) {
	// declaring Route for users
	mux.Handle("POST /users",
		manager.With(
			http.HandlerFunc(h.CreateUserHandler),
		),
	)

	mux.Handle("POST /login",
		manager.With(
			http.HandlerFunc(h.Login),
		),
	)

	mux.Handle("PUT /users/{id}",
		manager.With(
			http.HandlerFunc(h.UpdateUser),
			h.middlewares.AuthenticateJWT,
		),
	)

}
