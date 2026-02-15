package todoHandler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"todolist/model"
	"todolist/utils"
)

// CreateTask handles POST /tasks and adds a new task to the database.
func (h *Handler) CreateTask(w http.ResponseWriter, r *http.Request) {

	var newTask model.Task
	// creating decoder object
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&newTask)
	if err != nil {
		fmt.Println(err)

		http.Error(w, "Please provide valid json", http.StatusBadRequest)
		return
	}
	createdTask, err := h.taskrepo.StoreTask(newTask)

	// creating encoder object
	utils.SendData(w, createdTask, http.StatusCreated)

}
