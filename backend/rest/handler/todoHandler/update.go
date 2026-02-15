package todoHandler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"todolist/model"
	"todolist/utils"
)

// creating GET http Method to get data from database.
func (h *Handler) Updatetasks(w http.ResponseWriter, r *http.Request) {

	taskID := r.PathValue("id")

	id, err := strconv.Atoi(taskID)

	if err != nil {
		http.Error(w, "Please give me a valid task id", http.StatusBadRequest)
		return
	}

	var newTask model.Task
	// creating decoder object
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&newTask)
	if err != nil {
		fmt.Println(err)
		// http.Error(w, "Please provide valid json", 400)
		http.Error(w, "Please provide valid json", http.StatusBadRequest)
		return
	}

	newTask.Id = id
	// creating encoder object
	h.taskrepo.Update(newTask)

	utils.SendData(w, "Successfully updated product", http.StatusCreated)
}
