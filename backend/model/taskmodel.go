package model

import (
	"errors"
	"strings"
	"time"
)

type Task struct {
	ID          int       `json:"id" gorm:"primaryKey"`
	ProjectID   int       `json:"project_id"`
	Project     Project   `json:"project" gorm:"foreignKey:ProjectID"`
	Title       string    `json:"title"`
	Description *string   `json:"description,omitempty"`
	Status      string    `json:"status"`
	Priority    string    `json:"priority"`
	AssigneeID  *int      `json:"assignee_id,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	EndAt       time.Time `json:"end_at"`
}

func (t *Task) Validate() error {
	if strings.TrimSpace(t.Title) == "" {
		return errors.New("title is required")
	}

	if t.ProjectID == 0 {
		return errors.New("The task must be under a project")
	}

	if t.EndAt.IsZero() {
		return errors.New("End time is required")
	}

	if t.EndAt.Before(time.Now()) {
		return errors.New("End time can not be in past")
	}

	if t.EndAt.After(t.Project.EndAt) {
		return errors.New("The task ")
	}

	return nil
}
