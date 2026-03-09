package taskHandler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"todolist/model"
	"todolist/utils"
)

// CreateTask handles POST /tasks and adds a new task to the database.
func (h *Handler) CreateTaskHandler(w http.ResponseWriter, r *http.Request) {

	// Get logged-in user
	user, ok := r.Context().Value("user").(utils.Payload)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Get project ID from URL
	projectIDStr := r.PathValue("projectid")
	projectID, err := strconv.Atoi(projectIDStr)
	if err != nil {
		http.Error(w, "Invalid project id", http.StatusBadRequest)
		return
	}

	// Check project exists
	project, err := h.projectService.GetProject(projectID, user.ID)
	if err != nil {
		http.Error(w, "Project not found", http.StatusNotFound)
		return
	}

	// Authorization check
	isOwner := project.OwnerID == user.ID

	isMember := false
	member, err := h.projectMemberService.GetProjectMemberbyID(projectID, user.ID)
	if err == nil && member != nil {
		isMember = true
	}

	if !isOwner && !isMember {
		http.Error(w, "Forbidden: Not part of this project", http.StatusForbidden)
		return
	}

	// Decode request body
	var newTask model.Task
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&newTask); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Enforce project from URL (NOT from body)
	newTask.ProjectID = projectID

	// enforce assignee belongs to project
	if newTask.AssigneeID != nil {
		if *newTask.AssigneeID != user.ID {
			_, err := h.projectMemberService.GetProjectMemberbyID(projectID, *newTask.AssigneeID)
			if err != nil {
				http.Error(w, "Assignee is not part of this project", http.StatusBadRequest)
				return
			}
		}
	}

	// Create task
	createdTask, err := h.taskService.CreateTask(&newTask, user.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	utils.SendData(w, createdTask, http.StatusCreated)
}
