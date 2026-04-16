package projectHandler

import (
	"net/http"
	"todolist/utils"
)

func (h *Handler) DeleteProject(w http.ResponseWriter, r *http.Request) {

	// Get logged-in user
	user, ok := r.Context().Value("user").(utils.Payload)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Get project KEY instead of ID
	projectKey := r.PathValue("key")
	if projectKey == "" {
		http.Error(w, "Invalid project key", http.StatusBadRequest)
		return
	}

	// Get project using KEY + user
	project, err := h.projectService.GetProjectByKey(r.Context(), projectKey, user.ID)
	if err != nil {
		http.Error(w, "Project not found", http.StatusNotFound)
		return
	}

	// Authorization check
	if project.OwnerID != user.ID {
		http.Error(w, "Forbidden: You are not the owner of this project", http.StatusForbidden)
		return
	}

	// Delete by key
	err = h.projectService.DeleteProjectByKey(r.Context(), projectKey, user.ID)
	if err != nil {
		http.Error(w, "Error deleting project", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
