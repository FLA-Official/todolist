package projectMemberHandler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"todolist/utils"
)

// UpdateMemberRole updates the role of a member
func (h *Handler) UpdateMemberRole(w http.ResponseWriter, r *http.Request) {

	user, ok := r.Context().Value("user").(utils.Payload)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	userID := user.ID

	projectID, err := strconv.Atoi(r.PathValue("projectId"))
	if err != nil {
		http.Error(w, "invalid project id", http.StatusBadRequest)
		return
	}

	targetUserID, err := strconv.Atoi(r.PathValue("userId"))
	if err != nil {
		http.Error(w, "invalid user id", http.StatusBadRequest)
		return
	}

	var input struct {
		Role string `json:"role"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	err = h.projectMemberService.UpdateMemberRole(
		targetUserID,
		projectID,
		userID,
		input.Role,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(map[string]string{
		"status": "role updated",
	})
}
