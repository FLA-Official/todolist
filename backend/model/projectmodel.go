package model

import (
	"errors"
	"strings"
	"time"
)

type Project struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Key         string    `json:"key"`
	Description string    `json:"description,omitempty"`
	OwnerID     int       `json:"owner_id"`
	Partner     *int      `json:"partner"`
	CreatedAt   time.Time `json:"created_at"`
	EndAt       time.Time `json:"end_at"`
}

func (p *Project) Validate() error {
	if strings.TrimSpace(p.Name) == "" {
		return errors.New("Project Name is required")
	}

	if p.EndAt.IsZero() {
		return errors.New("End time is required")
	}

	if p.EndAt.Before(time.Now()) {
		return errors.New("End time can not be in past")
	}
	return nil
}
