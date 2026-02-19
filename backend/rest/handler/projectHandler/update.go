package projectHandler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"todolist/model"
	"todolist/utils"
)

// UpdateProject handles PUT /projects/{id} and updates an existing project.
func (h *Handler) UpdateProject(w http.ResponseWriter, r *http.Request) {

	projectID := r.PathValue("id")

	id, err := strconv.Atoi(projectID)

	if err != nil {
		http.Error(w, "Please give me a valid project id", http.StatusBadRequest)
		return
	}

	var newProject model.Project
	// creating decoder object
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&newProject)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Please provide valid json", http.StatusBadRequest)
		return
	}

	newProject.ID = id
	err = h.projectrepo.UpdateProject(&newProject)
	if err != nil {
		http.Error(w, "Error updating project", http.StatusInternalServerError)
		return
	}

	utils.SendData(w, newProject, http.StatusOK)

}
