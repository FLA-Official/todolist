package projectHandler

import (
	"net/http"
	"strconv"
	"todolist/utils"
)

func (h *Handler) DeleteProject(w http.ResponseWriter, r *http.Request) {

	// Get logged-in user from context
	user, ok := r.Context().Value("user").(utils.Payload)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Get project ID from URL
	projectID := r.PathValue("id")
	id, err := strconv.Atoi(projectID)
	if err != nil {
		http.Error(w, "Invalid project id", http.StatusBadRequest)
		return
	}

	// Fetch project from DB
	project, err := h.projectrepo.GetProjectByID(id)
	if err != nil {
		http.Error(w, "Project not found", http.StatusNotFound)
		return
	}

	// Authorization check
	if project.OwnerID != user.ID {
		http.Error(w, "Forbidden: You are not the owner of this project", http.StatusForbidden)
		return
	}

	// Delete project
	err = h.projectrepo.DeleteProject(id)
	if err != nil {
		http.Error(w, "Error deleting project", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
