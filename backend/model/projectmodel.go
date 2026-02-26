package model

import (
	"errors"
	"strings"
	"time"
)

type Project struct {
	ID          int        `db:"id" json:"id"`
	Name        string     `db:"name" json:"name"`
	Key         string     `db:"key" json:"key"`
	Description *string    `db:"description" json:"description,omitempty"`
	OwnerID     int        `db:"owner_id" json:"owner_id"`
	Partner     *int       `db:"partner" json:"partner,omitempty"`
	CreatedAt   time.Time  `db:"created_at" json:"created_at"`
	EndAt       *time.Time `db:"end_at" json:"end_at,omitempty"`
}

func (p *Project) Validate() error {
	if strings.TrimSpace(p.Name) == "" {
		return errors.New("project name is required")
	}

	if strings.TrimSpace(p.Key) == "" {
		return errors.New("project key is required")
	}

	// Key format (like Jira: uppercase short code)
	if len(p.Key) < 2 || len(p.Key) > 10 {
		return errors.New("project key must be 2-10 characters")
	}

	if p.EndAt != nil {
		if p.EndAt.Before(time.Now()) {
			return errors.New("end time cannot be in the past")
		}
	}

	p.Key = strings.ToUpper(p.Key)

	return nil
}
