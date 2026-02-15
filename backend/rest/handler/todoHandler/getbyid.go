package todoHandler

import (
	"net/http"
	"strconv"
	"todolist/utils"
)

// GetProductsByID handles GET /products/{id} and returns the requested product if found.
func (h *Handler) GetTaskByID(w http.ResponseWriter, r *http.Request) {
	// creating encoder object
	TaskID := r.PathValue("id")

	id, err := strconv.Atoi(TaskID)

	if err != nil {
		http.Error(w, "Please give me a valid product id", http.StatusBadRequest)
		return
	}

	task, _ := h.taskrepo.GetTaskByID(id)

	if task == nil {
		utils.SendData(w, "Task not found", http.StatusNotFound)
		return
	}
	// Return the found product with 200 OK
	utils.SendData(w, task, http.StatusOK)
}
