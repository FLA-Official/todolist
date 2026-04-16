package projectHandler

import (
	"net/http"
	"todolist/utils"
)

// GetProjectByKey handles GET /projects/{key} and returns the requested project if found.
func (h *Handler) GetProjectByKey(w http.ResponseWriter, r *http.Request) {

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

	project, err := h.projectService.GetProjectByKey(r.Context(), projectKey, user.ID)
	if err != nil {
		http.Error(w, "Error retrieving project", http.StatusInternalServerError)
		return
	}

	if project == nil {
		utils.SendData(w, "Project not found", http.StatusNotFound)
		return
	}

	// Return the found project with 200 OK
	utils.SendData(w, project, http.StatusOK)
}
