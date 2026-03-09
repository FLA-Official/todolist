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

	projectID := r.PathValue("projectid")

	projectid, err := strconv.Atoi(projectID)
	if err != nil {
		http.Error(w, "Please give me a valid task id", http.StatusBadRequest)
		return
	}

	// get logged-in user
	payload := r.Context().Value("user").(utils.Payload)

	// pass userID to repo
	err = h.taskService.DeleteTask(id, projectid, payload.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
