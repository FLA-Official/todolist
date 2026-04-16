package taskHandler

import (
	"encoding/json"
	"net/http"
	"todolist/model"
	"todolist/utils"
)

// CreateTask handles POST /tasks and adds a new task to the database.
func (h *Handler) CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	logger := utils.LoggerFromContext(r.Context())

	// Get logged-in user
	user, ok := r.Context().Value("user").(utils.Payload)
	if !ok {
		logger.Error("Wrong or corrupted JWT Token")
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Get project key from URL
	projectKey := r.PathValue("projectkey")
	if projectKey == "" {
		logger.Error("Invalid project key")
		http.Error(w, "Invalid project key", http.StatusBadRequest)
		return
	}

	// Check project exists and user has access
	project, err := h.projectService.GetProjectByKey(r.Context(), projectKey, user.ID)
	if err != nil {
		logger.Error("Project not found for this user", "user_ID", user.ID)
		http.Error(w, "Project not found", http.StatusNotFound)
		return
	}

	// Authorization check
	isOwner := project.OwnerID == user.ID

	isMember := false
	member, err := h.projectMemberService.GetProjectMemberbyID(project.ID, user.ID)
	if err == nil && member != nil {
		isMember = true
	}

	if !isOwner && !isMember {
		logger.Error("This user is not part of the project", "user_ID", user.ID)
		http.Error(w, "Forbidden: Not part of this project", http.StatusForbidden)
		return
	}

	// Decode request body
	var newTask model.Task
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&newTask); err != nil {
		logger.Error("Invalid Request Body")
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Enforce project from URL (NOT from body)
	newTask.ProjectKey = project.Key

	// enforce assignee belongs to project
	if newTask.AssigneeID != nil {
		if *newTask.AssigneeID != user.ID {
			_, err := h.projectMemberService.GetProjectMemberbyID(project.ID, *newTask.AssigneeID)
			if err != nil {
				logger.Error(
					"Assignee is not part of the project",
					"assignee_id", newTask.AssigneeID,
					"user_id", user.ID,
				)
				http.Error(w, "Assignee is not part of this project", http.StatusBadRequest)
				return
			}
		}
	}

	// Create task
	createdTask, err := h.taskService.CreateTask(r.Context(), &newTask, user.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	utils.SendData(w, createdTask, http.StatusCreated)
}
