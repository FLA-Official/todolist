package repo

import (
	"errors"
	"fmt"
	"todolist/model"

	"github.com/jmoiron/sqlx"
)

type ProjectRepo interface {
	CreateProject(project *model.Project) error
	GetProjectByID(id int) (*model.Project, error)
	UpdateProject(project *model.Project) error
	DeleteProject(projectID int) error
	ListProjects() ([]model.Project, error)
	ListProjectsByOwner(ownerID int) ([]model.Project, error)
	ListProjectsWhereUserIsAdmin(userID int) ([]model.Project, error)
	ListProjectsWhereUserIsMember(userID int) ([]model.Project, error)
}

type projectRepo struct {
	dbCon *sqlx.DB
}

func NewProjectRepo(dbCon *sqlx.DB) ProjectRepo {
	return &projectRepo{dbCon: dbCon}
}

// ID          int       `json:"id"`
// 	Name        string    `json:"name"`
// 	Description string    `json:"description,omitempty"`
// 	OwnerID     int       `json:"owner_id"`
// 	Partner     *int      `json:"partner"`
// 	CreatedAt   time.Time `json:"created_at"`

func (p *projectRepo) CreateProject(project *model.Project) error {
	tx, err := p.dbCon.Beginx()
	if err != nil {
		return err
	}
	err = project.Validate()
	if err != nil {
		return err
	}

	query := `
	INSERT INTO projects (name, key, description, owner_id)
	VALUES ($1, $2, $3, $4)
	RETURNING id, created_at
	`

	err = tx.QueryRow(
		query,
		project.Name,
		project.Key,
		project.Description,
		project.OwnerID,
	).Scan(&project.ID, &project.CreatedAt)

	if err != nil {
		tx.Rollback()
		return err
	}

	// ADD OWNER TO MEMBERS TABLE
	_, err = tx.Exec(`
	INSERT INTO project_members (project_id, user_id, role)
	VALUES ($1, $2, $3)
	`, project.ID, project.OwnerID, model.RoleOwner)

	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (p *projectRepo) GetProjectByID(id int) (*model.Project, error) {
	var project model.Project

	err := p.dbCon.Get(&project, `SELECT * FROM projects WHERE id=$1`, id)
	if err != nil {
		return nil, errors.New("project not found")
	}

	return &project, nil
}

func (p *projectRepo) UpdateProject(project *model.Project) error {
	var existing model.Project

	err := p.dbCon.Get(&existing, `SELECT * FROM projects WHERE id=$1`, project.ID)
	if err != nil {
		return errors.New("project not found")
	}

	// preserve created_at
	project.CreatedAt = existing.CreatedAt

	if err := project.Validate(); err != nil {
		return err
	}

	query := `
		UPDATE projects SET
			name=$1,
			key=$2,
			description=$3,
			owner_id=$4,
			partner=$5,
			end_at=$6
		WHERE id=$7
	`

	_, err = p.dbCon.Exec(
		query,
		project.Name,
		project.Key,
		project.Description,
		project.OwnerID,
		project.Partner,
		project.EndAt,
		project.ID,
	)

	return err
}

func (p *projectRepo) DeleteProject(projectID int) error {

	_, err := p.dbCon.Exec(
		`DELETE FROM projects WHERE id=$1`,
		projectID,
	)

	return err
}

func (p *projectRepo) ListProjects() ([]model.Project, error) {
	var projects []model.Project

	err := p.dbCon.Select(&projects, `SELECT * FROM projects`)
	return projects, err
}

func (p *projectRepo) ListProjectsByOwner(ownerID int) ([]model.Project, error) {
	var projects []model.Project

	err := p.dbCon.Select(&projects, `SELECT * FROM projects WHERE owner_id=$1`, ownerID)
	if len(projects) == 0 {
		return nil, fmt.Errorf("owner %d has no projects", ownerID)
	}

	return projects, err
}

// Check owner
func (p *projectRepo) IsOwner(projectID, userID int) (bool, error) {
	var exists bool

	query := `
	SELECT EXISTS (
		SELECT 1 FROM projects
		WHERE id=$1 AND owner_id=$2
	)
	`

	err := p.dbCon.Get(&exists, query, projectID, userID)
	return exists, err
}

func (p *projectRepo) ListProjectsWhereUserIsAdmin(userID int) ([]model.Project, error) {

	var projects []model.Project

	query := `
	SELECT p.*
	FROM projects p
	JOIN project_members pm ON pm.project_id = p.id
	WHERE pm.user_id = $1 AND pm.role = 'admin'
	`

	err := p.dbCon.Select(&projects, query, userID)
	return projects, err
}

func (p *projectRepo) ListProjectsWhereUserIsMember(userID int) ([]model.Project, error) {

	var projects []model.Project

	query := `
	SELECT p.*
	FROM projects p
	JOIN project_members pm ON pm.project_id = p.id
	WHERE pm.user_id = $1 AND pm.role = 'member'
	`

	err := p.dbCon.Select(&projects, query, userID)
	return projects, err
}
