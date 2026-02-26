package repo

import (
	"errors"
	"todolist/model"

	"github.com/jmoiron/sqlx"
)

type ProjectMemberRepo interface {
	AddMember(member *model.ProjectMember) error
	GetMember(projectID, userID int) (*model.ProjectMember, error)
	UpdateMemberRole(projectID, userID int, role string) error
	RemoveMember(projectID, userID int) error
	ListMembers(projectID int) ([]model.ProjectMember, error)
}

type projectMemberRepo struct {
	dbCon *sqlx.DB
}

func NewProjectMemberRepo(dbCon *sqlx.DB) ProjectMemberRepo {
	return &projectMemberRepo{
		dbCon: dbCon,
	}
}

func (r *projectMemberRepo) AddMember(member *model.ProjectMember) error {
	if err := member.Validate(); err != nil {
		return err
	}

	query := `
		INSERT INTO project_members (project_id, user_id, role)
		VALUES ($1, $2, $3)
		RETURNING id, joined_at
	`

	return r.dbCon.QueryRow(
		query,
		member.ProjectID,
		member.UserID,
		member.Role,
	).Scan(&member.ID, &member.JoinedAt)
}

func (r *projectMemberRepo) GetMember(projectID, userID int) (*model.ProjectMember, error) {
	var member model.ProjectMember

	query := `
		SELECT * FROM project_members
		WHERE project_id=$1 AND user_id=$2
	`

	err := r.dbCon.Get(&member, query, projectID, userID)
	if err != nil {
		return nil, errors.New("member not found")
	}

	return &member, nil
}

func (r *projectMemberRepo) UpdateMemberRole(projectID, userID int, role string) error {
	validRoles := map[string]bool{
		model.RoleOwner:  true,
		model.RoleAdmin:  true,
		model.RoleMember: true,
	}

	if !validRoles[role] {
		return errors.New("invalid role")
	}

	query := `
		UPDATE project_members
		SET role=$1
		WHERE project_id=$2 AND user_id=$3
	`

	result, err := r.dbCon.Exec(query, role, projectID, userID)
	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return errors.New("member not found")
	}

	return nil
}

func (r *projectMemberRepo) RemoveMember(projectID, userID int) error {
	query := `
		DELETE FROM project_members
		WHERE project_id=$1 AND user_id=$2
	`

	result, err := r.dbCon.Exec(query, projectID, userID)
	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return errors.New("member not found")
	}

	return nil
}

func (r *projectMemberRepo) ListMembers(projectID int) ([]model.ProjectMember, error) {
	var members []model.ProjectMember

	query := `
		SELECT * FROM project_members
		WHERE project_id=$1
	`

	err := r.dbCon.Select(&members, query, projectID)
	if len(members) == 0 {
		return nil, errors.New("no members found")
	}

	return members, err
}
