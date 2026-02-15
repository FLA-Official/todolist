package todoHandler

import (
	"net/http"
	"todolist/utils"
)

// GetTasks handles GET /tasks and returns a list of all tasks.
func (h *Handler) GetTasks(w http.ResponseWriter, r *http.Request) {
	// creating encoder object
	utils.SendData(w, h.taskrepo.List(), http.StatusOK)
}
