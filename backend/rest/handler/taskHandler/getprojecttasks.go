package taskHandler

import (
	"net/http"
	"strconv"
	"todolist/utils"
)

// GetTasks handles GET /tasks and returns a list of all tasks.
func (h *Handler) GetProjectTasks(w http.ResponseWriter, r *http.Request) {

	// extract user
	payload := r.Context().Value("user").(utils.Payload)
	userID := payload.ID

	projectIDStr := r.PathValue("projectid")
	projectID, err := strconv.Atoi(projectIDStr)
	if err != nil {
		http.Error(w, "Invalid project id", http.StatusBadRequest)
		return
	}

	// check ownership + fetch tasks
	allTask, err := h.taskService.GetProjectTasks(projectID, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	utils.SendData(w, allTask, http.StatusOK)
}
