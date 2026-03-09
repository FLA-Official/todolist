package projectMemberHandler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"todolist/model"
	"todolist/utils"
)

var input struct {
	UserID int    `json:"user_id"`
	Role   string `json:"role"`
}

// CreateProjectHandler handles POST /projects and adds a new project to the database.
func (h *Handler) AddMemberHandler(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value("user").(utils.Payload)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	projectIDStr := r.PathValue("projectID")
	projectID, err := strconv.Atoi(projectIDStr)
	if err != nil {
		http.Error(w, "invalid projectId", http.StatusBadRequest)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	member := &model.ProjectMember{
		ProjectID: projectID,
		UserID:    input.UserID,
		Role:      input.Role,
	}

	if err := h.projectMemberService.AddMember(member, user.ID); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(member)
}
