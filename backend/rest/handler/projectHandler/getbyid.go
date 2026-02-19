package projectHandler

import (
	"net/http"
	"strconv"
	"todolist/utils"
)

// GetProjectByID handles GET /projects/{id} and returns the requested project if found.
func (h *Handler) GetProjectByID(w http.ResponseWriter, r *http.Request) {
	// creating encoder object
	projectID := r.PathValue("id")

	id, err := strconv.Atoi(projectID)

	if err != nil {
		http.Error(w, "Invalid project id", http.StatusBadRequest)
		return
	}

	project, err := h.projectrepo.GetProjectByID(id)
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
