package projectMemberHandler

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// UpdateMemberRole updates the role of a member
func (h *Handler) UpdateMemberRole(w http.ResponseWriter, r *http.Request) {
	projectID, _ := strconv.Atoi(r.PathValue("projectId"))
	userID, _ := strconv.Atoi(r.PathValue("userId"))

	var input struct {
		Role string `json:"role"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.projectmemberrepo.UpdateMemberRole(projectID, userID, input.Role); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "role updated"})
}
