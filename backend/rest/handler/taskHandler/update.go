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

	projectKey := r.PathValue("projectkey")
	if projectKey == "" {
		http.Error(w, "Invalid project key", http.StatusBadRequest)
		return
	}

	payload := r.Context().Value("user").(utils.Payload)
	userID := payload.ID

	// Get project by key to verify access and get ID
	project, err := h.projectService.GetProjectByKey(r.Context(), projectKey, userID)
	if err != nil {
		http.Error(w, "Project not found or access denied", http.StatusNotFound)
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

	fmt.Println("Logged user ID:", userID)
	fmt.Println("Project Key:", projectKey)

	newTask.ID = taskid
	newTask.ProjectKey = project.Key

	err = h.taskService.UpdateTask(&newTask, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	utils.SendData(w, newTask, http.StatusOK)
}
