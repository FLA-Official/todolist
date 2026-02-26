package model

import (
	"errors"
	"time"
)

const (
	StatusTodo       = "todo"
	StatusInProgress = "in_progress"
	StatusDone       = "done"
	PriorityLow      = "low"
	PriorityMedium   = "medium"
	PriorityHigh     = "high"
)

type Task struct {
	ID          int        `db:"id" json:"id"`
	ProjectID   int        `db:"project_id" json:"project_id"`
	Title       string     `db:"title" json:"title"`
	Description *string    `db:"description" json:"description,omitempty"`
	Status      string     `db:"status" json:"status"`
	Priority    string     `db:"priority" json:"priority"`
	AssigneeID  *int       `db:"assignee_id" json:"assignee_id,omitempty"`
	CreatedAt   time.Time  `db:"created_at" json:"created_at"`
	EndAt       *time.Time `db:"end_at" json:"end_at,omitempty"`
}

func (t *Task) Validate() error {
	if t.Title == "" {
		return errors.New("title is required")
	}

	validStatus := map[string]bool{
		StatusTodo:       true,
		StatusInProgress: true,
		StatusDone:       true,
	}
	if !validStatus[t.Status] {
		return errors.New("invalid status")
	}

	validPriority := map[string]bool{
		PriorityLow:    true,
		PriorityMedium: true,
		PriorityHigh:   true,
	}
	if !validPriority[t.Priority] {
		return errors.New("invalid priority")
	}

	if t.Status != StatusDone && t.EndAt != nil {
		return errors.New("end_at only allowed when task is done")
	}

	if t.Status == StatusDone && t.EndAt == nil {
		now := time.Now()
		t.EndAt = &now
	}

	return nil
}

func isValidTransition(old, new string) bool {
	switch old {
	case StatusTodo:
		return new == StatusInProgress || new == StatusDone
	case StatusInProgress:
		return new == StatusDone
	case StatusDone:
		return false
	}
	return false
}
