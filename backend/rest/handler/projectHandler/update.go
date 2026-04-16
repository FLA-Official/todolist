package projectHandler

import (
	"encoding/json"
	"net/http"

	"todolist/model"
	"todolist/utils"
)

// UpdateProject handles PUT /projects/{key} and updates an existing project.
func (h *Handler) UpdateProject(w http.ResponseWriter, r *http.Request) {

	// Get logged-in user from context
	user, ok := r.Context().Value("user").(utils.Payload)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Get project key from path
	projectKey := r.PathValue("key")
	if projectKey == "" {
		http.Error(w, "Invalid project key", http.StatusBadRequest)
		return
	}

	// Fetch existing project
	existingProject, err := h.projectService.GetProjectByKey(r.Context(), projectKey, user.ID)
	if err != nil {
		http.Error(w, "Project not found", http.StatusNotFound)
		return
	}

	// Authorization check
	if existingProject.OwnerID != user.ID {
		http.Error(w, "Forbidden: You are not the owner", http.StatusForbidden)
		return
	}

	// Decode new data
	var newProject model.Project
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&newProject); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Preserve immutable fields
	newProject.ID = existingProject.ID
	newProject.Key = existingProject.Key             // preserve key
	newProject.OwnerID = existingProject.OwnerID     // prevent owner change
	newProject.CreatedAt = existingProject.CreatedAt // preserve original time

	// Update
	err = h.projectService.UpdateProject(&newProject, user.ID)
	if err != nil {
		http.Error(w, "Error updating project", http.StatusInternalServerError)
		return
	}

	utils.SendData(w, newProject, http.StatusOK)
}
