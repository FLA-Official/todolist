package projectHandler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"todolist/model"
	"todolist/utils"
)

// UpdateProject handles PUT /projects/{id} and updates an existing project.
func (h *Handler) UpdateProject(w http.ResponseWriter, r *http.Request) {

	// Get logged-in user from context
	user, ok := r.Context().Value("user").(utils.Payload)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Get project ID
	projectID := r.PathValue("id")
	id, err := strconv.Atoi(projectID)
	if err != nil {
		http.Error(w, "Invalid project id", http.StatusBadRequest)
		return
	}

	// Fetch existing project
	existingProject, err := h.projectService.GetProject(r.Context(), id, user.ID)
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
	newProject.ID = id
	newProject.OwnerID = existingProject.OwnerID     // prevent owner change
	newProject.CreatedAt = existingProject.CreatedAt // preserve original time

	// (optional but recommended)
	// prevent user from manually setting past created_at or owner_id

	// Update
	err = h.projectService.UpdateProject(&newProject, user.ID)
	if err != nil {
		http.Error(w, "Error updating project", http.StatusInternalServerError)
		return
	}

	utils.SendData(w, newProject, http.StatusOK)
}
