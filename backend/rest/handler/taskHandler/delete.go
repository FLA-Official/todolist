package taskHandler

import (
	"net/http"
	"strconv"
	"todolist/utils"
)

func (h *Handler) DeleteTask(w http.ResponseWriter, r *http.Request) {

	taskID := r.PathValue("taskid")

	id, err := strconv.Atoi(taskID)
	if err != nil {
		http.Error(w, "Please give me a valid task id", http.StatusBadRequest)
		return
	}

	projectKey := r.PathValue("projectkey")
	if projectKey == "" {
		http.Error(w, "Invalid project key", http.StatusBadRequest)
		return
	}

	// get logged-in user
	payload := r.Context().Value("user").(utils.Payload)

	// Verify project access by key
	_, err = h.projectService.GetProjectByKey(r.Context(), projectKey, payload.ID)
	if err != nil {
		http.Error(w, "Project not found or access denied", http.StatusNotFound)
		return
	}

	// pass project key and user ID to service
	err = h.taskService.DeleteTask(projectKey, id, payload.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
