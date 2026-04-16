package projectMemberHandler

import (
	"net/http"
	"strconv"
	"todolist/utils"
)

func (h *Handler) RemoveMember(w http.ResponseWriter, r *http.Request) {

	user, ok := r.Context().Value("user").(utils.Payload)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	projectKey := r.PathValue("projectkey")
	if projectKey == "" {
		http.Error(w, "Invalid project key", http.StatusBadRequest)
		return
	}

	// Verify project access by key
	_, err := h.projectService.GetProjectByKey(r.Context(), projectKey, user.ID)
	if err != nil {
		http.Error(w, "Project not found or access denied", http.StatusNotFound)
		return
	}

	targetMemberID, _ := strconv.Atoi(r.PathValue("userid"))

	if err := h.projectMemberService.RemoveMember(r.Context(), projectKey, targetMemberID, user.ID); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
