package projectMemberHandler

import (
	"net/http"
	"strconv"
)

func (h *Handler) RemoveMember(w http.ResponseWriter, r *http.Request) {
	projectID, _ := strconv.Atoi(r.PathValue("projectId"))
	userID, _ := strconv.Atoi(r.PathValue("userId"))

	if err := h.projectmemberrepo.RemoveMember(projectID, userID); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
