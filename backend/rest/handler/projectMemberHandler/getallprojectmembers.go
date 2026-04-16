package projectMemberHandler

import (
	"encoding/json"
	"net/http"
)

// GetMembersByProject fetches all members of a project
func (h *Handler) GetMembersByProject(w http.ResponseWriter, r *http.Request) {
	projectKey := r.PathValue("projectkey")
	if projectKey == "" {
		http.Error(w, "Invalid project key", http.StatusBadRequest)
		return
	}

	// Get user from context for authorization
	user, ok := r.Context().Value("user").(map[string]interface{})
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	userID := int(user["id"].(float64))

	// Get project by key to verify access and get ID
	project, err := h.projectService.GetProjectByKey(r.Context(), projectKey, userID)
	if err != nil {
		http.Error(w, "Project not found or access denied", http.StatusNotFound)
		return
	}

	members, err := h.projectMemberService.GetProjectMembers(project.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(members)
}
