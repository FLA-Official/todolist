package userHandler

import (
	"net/http"
	"todolist/utils"
)

// GetUsers handles GET /users and returns a list of all users.
func (h *Handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	allUsers, err := h.userrepo.ListUsers()

	if err != nil && len(allUsers) == 0 {
		http.Error(w, "No User Available", http.StatusNotFound)
		return
	}

	utils.SendData(w, allUsers, http.StatusOK)
}
