package taskHandler

import (
	"net/http"
	"todolist/utils"
)

// GetTasks handles GET /tasks and returns a list of all tasks.
func (h *Handler) GetTasks(w http.ResponseWriter, r *http.Request) {
	// creating encoder object
	allTask, err := h.taskrepo.ListTasks()

	if err != nil {
		http.Error(w, "No Task Available", http.StatusBadGateway)
	}

	utils.SendData(w, allTask, http.StatusOK)
}
