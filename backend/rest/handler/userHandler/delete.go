package userHandler

import (
	"net/http"
	"strconv"
)

func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {

	userID := r.PathValue("id")

	id, err := strconv.Atoi(userID)

	if err != nil {
		http.Error(w, "Please give me a valid user id", http.StatusBadRequest)
		return
	}

	err = h.userrepo.DeleteUser(id)
	if err != nil {
		http.Error(w, "Error deleting user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
