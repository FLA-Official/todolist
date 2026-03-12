package taskHandler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"todolist/model"
	"todolist/utils"
)

// UpdateTask handles PUT /tasks/{id} and updates an existing task.
func (h *Handler) UpdateTask(w http.ResponseWriter, r *http.Request) {

	taskID := r.PathValue("taskid")

	taskid, err := strconv.Atoi(taskID)
	if err != nil {
		http.Error(w, "Please give me a valid task id", http.StatusBadRequest)
		return
	}

	projectID := r.PathValue("projectid")

	projectid, err := strconv.Atoi(projectID)
	if err != nil {
		http.Error(w, "Please give me a valid project id", http.StatusBadRequest)
		return
	}

	var newTask model.Task
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err = decoder.Decode(&newTask)
	if err != nil {
		http.Error(w, "Please provide valid json", http.StatusBadRequest)
		return
	}

	payload := r.Context().Value("user").(utils.Payload)
	userID := payload.ID

	fmt.Println("Logged user ID:", userID)
	fmt.Println("Project ID:", projectid)

	newTask.ID = taskid
	newTask.ProjectID = projectid

	err = h.taskService.UpdateTask(&newTask, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	utils.SendData(w, newTask, http.StatusOK)
}
