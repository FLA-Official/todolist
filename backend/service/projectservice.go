package service

import (
	"context"
	"errors"
	"todolist/model"
	"todolist/repo"
	"todolist/utils"
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

func (s *ProjectService) CreateProject(ctx context.Context, project *model.Project, userID int) error {

	logger := utils.LoggerFromContext(ctx)

	project.OwnerID = userID

	err := s.projectRepo.CreateProject(project)
	if err != nil {
		logger.Error("failed to create project", "user_id", userID)
		return err
	}

	logger.Info("project created", "project_id", project.ID, "user_id", userID)

	member := &model.ProjectMember{
		ProjectID: project.ID,
		UserID:    userID,
		Role:      model.RoleOwner,
	}

	return s.memberRepo.AddMember(member)
}
func (s *ProjectService) GetProject(ctx context.Context, projectID, userID int) (*model.Project, error) {

	logger := utils.LoggerFromContext(ctx)

	role, err := s.memberRepo.GetUserRole(projectID, userID)
	if err != nil {
		logger.Error("access denied", "project_id", projectID, "user_id", userID)
		return nil, errors.New("no access to this project")
	}

	if role == "" {
		logger.Error("not a member", "project_id", projectID, "user_id", userID)
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

	return s.projectRepo.ListProjectsByOwner(userID)
}

func (s *ProjectService) ListUserProjectsAsAdmin(userID int) ([]model.Project, error) {
	return s.projectRepo.ListProjectsWhereUserIsAdmin(userID)
}

func (s *ProjectService) ListUserProjectAsMember(userID int) ([]model.Project, error) {
	return s.projectRepo.ListProjectsWhereUserIsMember(userID)
}

func (s *ProjectService) IsOwner(projectID, userID int) (bool, error) {

	project, err := s.projectRepo.GetProjectByID(projectID)
	if err != nil {
		return false, err
	}

	return project.OwnerID == userID, nil
}
