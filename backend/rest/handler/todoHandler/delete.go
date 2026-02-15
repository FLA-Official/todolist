package todoHandler

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

	h.taskrepo.Delete(id)

	utils.SendData(w, "Successfully deleted the task", http.StatusCreated)
}
