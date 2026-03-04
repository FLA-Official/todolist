package taskHandler

import (
	"net/http"
	"strconv"
	"todolist/utils"
)

// GetTaskByID handles GET /tasks/{id} and returns the requested task if found.
func (h *Handler) GetTaskByID(w http.ResponseWriter, r *http.Request) {

	taskID := r.PathValue("id")

	id, err := strconv.Atoi(taskID)
	if err != nil {
		http.Error(w, "Invalid task id", http.StatusBadRequest)
		return
	}

	payload := r.Context().Value("user").(utils.Payload)
	userID := payload.ID

	task, err := h.taskrepo.GetTaskByID(id, userID)
	if err != nil {
		http.Error(w, "Error retrieving task", http.StatusInternalServerError)
		return
	}

	if task == nil {
		utils.SendData(w, "Task not found", http.StatusNotFound)
		return
	}

	utils.SendData(w, task, http.StatusOK)
}
