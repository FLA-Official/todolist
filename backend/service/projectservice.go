package service

import (
	"errors"
	"todolist/model"
	"todolist/repo"
)

type ProjectService struct {
	projectRepo repo.ProjectRepo
	memberRepo  repo.ProjectMemberRepo
}

func NewProjectService(
	projectRepo repo.ProjectRepo,
	memberRepo repo.ProjectMemberRepo,
) *ProjectService {

	return &ProjectService{
		projectRepo: projectRepo,
		memberRepo:  memberRepo,
	}
}

func (s *ProjectService) CreateProject(project *model.Project, userID int) error {

	project.OwnerID = userID // automatically assign owner

	return s.projectRepo.CreateProject(project)
}

func (s *ProjectService) GetProject(projectID, userID int) (*model.Project, error) {

	role, err := s.memberRepo.GetUserRole(projectID, userID)
	if err != nil {
		return nil, errors.New("no access to this project")
	}

	if role == "" {
		return nil, errors.New("not a member")
	}

	return s.projectRepo.GetProjectByID(projectID)
}

func (s *ProjectService) UpdateProject(project *model.Project, userID int) error {

	role, err := s.memberRepo.GetUserRole(project.ID, userID)
	if err != nil {
		return err
	}

	if role != model.RoleOwner && role != model.RoleAdmin {
		return errors.New("no permission to update project")
	}

	return s.projectRepo.UpdateProject(project)
}

func (s *ProjectService) DeleteProject(projectID, userID int) error {

	isOwner, err := s.IsOwner(projectID, userID)
	if err != nil {
		return err
	}

	if !isOwner {
		return errors.New("only owner can delete project")
	}

	return s.projectRepo.DeleteProject(projectID)
}

func (s *ProjectService) ListUserProjects(userID int) ([]model.Project, error) {

	members, err := s.memberRepo.GetMembersByProject(userID)
	if err != nil {
		return nil, err
	}

	var projects []model.Project

	for _, m := range members {

		p, err := s.projectRepo.GetProjectByID(m.ProjectID)
		if err == nil {
			projects = append(projects, *p)
		}
	}

	return projects, nil
}

func (s *ProjectService) IsOwner(projectID, userID int) (bool, error) {

	project, err := s.projectRepo.GetProjectByID(projectID)
	if err != nil {
		return false, err
	}

	return project.OwnerID == userID, nil
}
