package taskHandler

import (
	"net/http"
	"todolist/utils"
)

// GetTasks handles GET /tasks and returns a list of all tasks.
func (h *Handler) GetTasks(w http.ResponseWriter, r *http.Request) {
	allTask, err := h.taskrepo.ListTasks()

	if err != nil && len(allTask) == 0 {
		http.Error(w, "No Task Available", http.StatusNotFound)
		return
	}

	utils.SendData(w, allTask, http.StatusOK)
}
