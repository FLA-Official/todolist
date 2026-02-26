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
	DeleteProject(id int) error
	ListProjects() ([]model.Project, error)
	ListProjectsByOwner(ownerID int) ([]model.Project, error)
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
	if err := project.Validate(); err != nil {
		return err
	}

	query := `
		INSERT INTO projects 
		(name, key, description, owner_id, partner, end_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, created_at
	`

	return p.dbCon.QueryRow(
		query,
		project.Name,
		project.Key,
		project.Description,
		project.OwnerID,
		project.Partner,
		project.EndAt,
	).Scan(&project.ID, &project.CreatedAt)
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

func (p *projectRepo) DeleteProject(id int) error {
	result, err := p.dbCon.Exec(`DELETE FROM projects WHERE id=$1`, id)
	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return errors.New("project not found")
	}

	return nil
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
