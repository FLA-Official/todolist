package userHandler

import (
	"net/http"
	"todolist/utils"
)

// GetUsers handles GET /users and returns a list of all users.
func (h *Handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	// creating encoder object
	allUsers, err := h.userrepo.ListUsers()

	if err != nil {
		http.Error(w, "No User Available", http.StatusBadGateway)
	}

	utils.SendData(w, allUsers, http.StatusOK)
}
