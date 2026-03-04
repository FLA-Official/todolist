package projectHandler

import (
	"net/http"
	"todolist/utils"
)

// GetProjects handles GET /projects and returns a list of all projects.
func (h *Handler) GetProjects(w http.ResponseWriter, r *http.Request) {

	user, ok := r.Context().Value("user").(utils.Payload)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	projects, err := h.projectrepo.ListProjectsByOwner(user.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	utils.SendData(w, projects, http.StatusOK)
}
