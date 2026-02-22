package repo

import (
	"errors"
	"fmt"
	"time"
	"todolist/model"
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
	projectList []*model.Project
}

func NewProjectRepo() ProjectRepo {
	repo := &projectRepo{}
	return repo
}

// ID          int       `json:"id"`
// 	Name        string    `json:"name"`
// 	Description string    `json:"description,omitempty"`
// 	OwnerID     int       `json:"owner_id"`
// 	Partner     *int      `json:"partner"`
// 	CreatedAt   time.Time `json:"created_at"`

func (p *projectRepo) CreateProject(project *model.Project) error {
	//adding ID
	project.ID = len(p.projectList) + 1
	//Adding time
	project.CreatedAt = time.Now()
	err := project.Validate()
	if err != nil {
		return err
	} else {
		p.projectList = append(p.projectList, project)
	}

	return nil
}

func (p *projectRepo) GetProjectByID(id int) (*model.Project, error) {
	for _, project := range p.projectList {
		if project.ID == id {
			return project, nil
		}
	}
	return nil, errors.New("Task not found")
}

func (p *projectRepo) UpdateProject(uproject *model.Project) error {
	for idx, project := range p.projectList {
		if project.ID == uproject.ID {
			uproject.CreatedAt = project.CreatedAt
			p.projectList[idx] = uproject
			return nil
		}
	}

	return errors.New("Project not found")
}

func (p *projectRepo) DeleteProject(id int) error {
	var tempList []*model.Project

	for _, project := range p.projectList {
		if project.ID != id {
			tempList = append(tempList, project)
		}
	}
	//to maintain ID order 1,2,3,4....
	for i := 0; i < len(tempList); i++ {
		tempList[i].ID = i + 1
	}

	p.projectList = tempList

	return nil

}

func (p *projectRepo) ListProjects() ([]model.Project, error) {
	var projects []model.Project
	for _, project := range p.projectList {
		projects = append(projects, *project)
	}
	return projects, nil
}

func (p *projectRepo) ListProjectsByOwner(ownerID int) ([]model.Project, error) {
	var tempList []model.Project

	for _, project := range p.projectList {
		if project.OwnerID == ownerID {
			tempList = append(tempList, *project)
		}
	}
	if len(tempList) > 0 {
		return tempList, nil
	} else {
		return nil, fmt.Errorf("Owener with the ID:%d Does not have any Project", ownerID)
	}

}
