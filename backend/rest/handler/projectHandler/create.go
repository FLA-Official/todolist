package projectHandler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"todolist/model"
	"todolist/utils"
)

// CreateProjectHandler handles POST /projects and adds a new project to the database.
func (h *Handler) CreateProjectHandler(w http.ResponseWriter, r *http.Request) {

	var newProject model.Project
	// creating decoder object
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&newProject)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Please provide valid json", http.StatusBadRequest)
		return
	}
	user := r.Context().Value("user").(utils.Payload) // JWT payload
	newProject.OwnerID = user.ID                      // automatically assign owner

	err = h.projectrepo.CreateProject(&newProject)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// creating encoder object
	utils.SendData(w, newProject, http.StatusCreated)

}
