package taskHandler

import (
	"net/http"
	"strconv"
	"todolist/utils"
)

func (h *Handler) DeleteTask(w http.ResponseWriter, r *http.Request) {

	taskID := r.PathValue("id")

	id, err := strconv.Atoi(taskID)
	if err != nil {
		http.Error(w, "Please give me a valid task id", http.StatusBadRequest)
		return
	}

	// get logged-in user
	payload := r.Context().Value("user").(utils.Payload)
	userID := payload.ID

	// pass userID to repo
	err = h.taskrepo.DeleteTask(id, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
