package model

import (
	"errors"
	"time"
)

const (
	RoleOwner  = "owner"
	RoleAdmin  = "admin"
	RoleMember = "member"
)

type ProjectMember struct {
	ID        int       `db:"id" json:"id"`
	ProjectID int       `db:"project_id" json:"project_id"`
	UserID    int       `db:"user_id" json:"user_id"`
	Role      string    `db:"role" json:"role"`
	JoinedAt  time.Time `db:"joined_at" json:"joined_at"`
}

func (pm *ProjectMember) Validate() error {
	if pm.ProjectID <= 0 {
		return errors.New("invalid project id")
	}

	if pm.UserID <= 0 {
		return errors.New("invalid user id")
	}

	validRoles := map[string]bool{
		RoleOwner:  true,
		RoleAdmin:  true,
		RoleMember: true,
	}

	if !validRoles[pm.Role] {
		return errors.New("invalid role")
	}

	return nil
}
