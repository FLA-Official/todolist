package taskHandler

import (
	"net/http"
	"strconv"
	"todolist/utils"
)

// GetTaskByID handles GET /tasks/{id} and returns the requested task if found.
func (h *Handler) GetTaskByID(w http.ResponseWriter, r *http.Request) {
	// creating encoder object
	TaskID := r.PathValue("id")

	id, err := strconv.Atoi(TaskID)

	if err != nil {
		http.Error(w, "Invalid task id", http.StatusBadRequest)
		return
	}

	task, err := h.taskrepo.GetTaskByID(id)
	if err != nil {
		http.Error(w, "Error retrieving task", http.StatusInternalServerError)
		return
	}

	if task == nil {
		utils.SendData(w, "Task not found", http.StatusNotFound)
		return
	}
	// Return the found product with 200 OK
	utils.SendData(w, task, http.StatusOK)
}
