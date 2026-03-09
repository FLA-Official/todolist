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
	projectID, _ := strconv.Atoi(r.PathValue("projectId"))
	targetUserID, _ := strconv.Atoi(r.PathValue("userId"))

	if err := h.projectMemberService.RemoveMember(projectID, targetUserID); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
