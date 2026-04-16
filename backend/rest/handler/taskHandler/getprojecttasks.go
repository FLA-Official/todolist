package taskHandler

import (
	"net/http"
	"todolist/utils"
)

// GetTasks handles GET /tasks and returns a list of all tasks.
func (h *Handler) GetProjectTasks(w http.ResponseWriter, r *http.Request) {

	// extract user
	payload := r.Context().Value("user").(utils.Payload)
	userID := payload.ID

	projectKey := r.PathValue("projectkey")
	if projectKey == "" {
		http.Error(w, "Invalid project key", http.StatusBadRequest)
		return
	}

	// Get project by key to get the ID for authorization
	project, err := h.projectService.GetProjectByKey(r.Context(), projectKey, userID)
	if err != nil {
		http.Error(w, "Project not found or access denied", http.StatusNotFound)
		return
	}

	// check ownership + fetch tasks
	allTask, err := h.taskService.GetProjectTasks(project.ID, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	utils.SendData(w, allTask, http.StatusOK)
}
