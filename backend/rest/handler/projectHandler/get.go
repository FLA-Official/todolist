package projectHandler

import (
	"net/http"
	"todolist/utils"
)

// GetProjects handles GET /projects and returns a list of all projects.
func (h *Handler) GetProjects(w http.ResponseWriter, r *http.Request) {
	allProjects, err := h.projectrepo.ListProjects()

	if err != nil && len(allProjects) == 0 {
		http.Error(w, "No Project Available", http.StatusNotFound)
		return
	}

	utils.SendData(w, allProjects, http.StatusOK)
}
