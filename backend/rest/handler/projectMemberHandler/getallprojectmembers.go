package projectMemberHandler

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// GetMembersByProject fetches all members of a project
func (h *Handler) GetMembersByProject(w http.ResponseWriter, r *http.Request) {
	projectIDStr := r.PathValue("projectId")
	projectID, err := strconv.Atoi(projectIDStr)
	if err != nil {
		http.Error(w, "invalid projectId", http.StatusBadRequest)
		return
	}

	members, err := h.projectmemberrepo.GetMembersByProject(projectID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(members)
}
