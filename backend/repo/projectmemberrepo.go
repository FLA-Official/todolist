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
	GetMembersByProject(projectID int) ([]model.ProjectMember, error)
	GetUserRole(projectKey string, userID int) (string, error)
	GetProjectsByUser(userID int) ([]model.ProjectMember, error)
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
		ON CONFLICT (project_id, user_id) DO NOTHING
	`

	_, err := r.dbCon.Exec(
		query,
		member.ProjectID,
		member.UserID,
		member.Role,
	)
	return err
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

func (r *projectMemberRepo) GetMembersByProject(projectID int) ([]model.ProjectMember, error) {
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

func (r *projectMemberRepo) GetUserRole(projectKey string, userID int) (string, error) {

	var role string

	err := r.dbCon.Get(&role,
		`SELECT pm.role
		 FROM project_members pm
		 JOIN projects p ON p.id = pm.project_id
		 WHERE p.key = $1 AND pm.user_id = $2`,
		projectKey, userID,
	)

	if err != nil {
		return "", errors.New("user is not a member of this project")
	}

	return role, nil
}

func (r *projectMemberRepo) GetProjectsByUser(userID int) ([]model.ProjectMember, error) {

	var members []model.ProjectMember

	query := `
	SELECT * FROM project_members
	WHERE user_id=$1
	`

	err := r.dbCon.Select(&members, query, userID)
	if err != nil {
		return nil, err
	}

	if len(members) == 0 {
		return nil, errors.New("user is not part of any project")
	}

	return members, nil
}
