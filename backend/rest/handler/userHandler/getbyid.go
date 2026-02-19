package userHandler

import (
	"net/http"
	"strconv"
	"todolist/utils"
)

// GetUserByID handles GET /users/{id} and returns the requested user if found.
func (h *Handler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	// creating encoder object
	userID := r.PathValue("id")

	id, err := strconv.Atoi(userID)

	if err != nil {
		http.Error(w, "Invalid user id", http.StatusBadRequest)
		return
	}

	user, err := h.userrepo.GetUserByID(id)
	if err != nil {
		http.Error(w, "Error retrieving user", http.StatusInternalServerError)
		return
	}

	if user == nil {
		utils.SendData(w, "User not found", http.StatusNotFound)
		return
	}
	// Return the found user with 200 OK
	utils.SendData(w, user, http.StatusOK)
}
