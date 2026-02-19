package projectHandler

import (
	"net/http"
	"todolist/utils"
)

// GetProjects handles GET /projects and returns a list of all projects.
func (h *Handler) GetProjects(w http.ResponseWriter, r *http.Request) {
	// creating encoder object
	allProjects, err := h.projectrepo.ListProjects()

	if err != nil {
		http.Error(w, "No Project Available", http.StatusBadGateway)
	}

	utils.SendData(w, allProjects, http.StatusOK)
}
