package projectHandler

import (
	"net/http"
	"strconv"
)

func (h *Handler) DeleteProject(w http.ResponseWriter, r *http.Request) {

	projectID := r.PathValue("id")

	id, err := strconv.Atoi(projectID)

	if err != nil {
		http.Error(w, "Please give me a valid project id", http.StatusBadRequest)
		return
	}

	err = h.projectrepo.DeleteProject(id)
	if err != nil {
		http.Error(w, "Error deleting project", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
