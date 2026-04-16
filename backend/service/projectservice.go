package service

import (
	"context"
	"errors"
	"fmt"
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

	// Start transaction
	tx, err := s.projectRepo.BeginTx()
	if err != nil {
		logger.Error("failed to start transaction", "error", err)
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	// 1. Generate prefix from project name in service layer
	prefix := utils.GeneratePrefix(project.Name)
	project.Prefix = &prefix

	// 2. Get next sequence (atomic)
	seq, err := s.projectRepo.GetNextProjectSequenceTx(tx, prefix)
	if err != nil {
		logger.Error("failed to get project sequence", "error", err)
		tx.Rollback()
		return err
	}
	project.Sequence = &seq

	// 3. Build final key
	finalKey := fmt.Sprintf("%s-%d", prefix, seq)
	project.Key = finalKey

	// 4. Store project
	err = s.projectRepo.CreateProjectTx(tx, project)
	if err != nil {
		logger.Error("failed to create project", "user_id", userID, "error", err)
		tx.Rollback()
		return err
	}

	// 5. Commit
	if err = tx.Commit(); err != nil {
		logger.Error("failed to commit transaction", "error", err)
		return err
	}

	logger.Info("project created",
		"project_id", project.ID,
		"key", project.Key,
		"user_id", userID,
	)

	return nil
}
func (s *ProjectService) GetProjectByKey(ctx context.Context, key string, userID int) (*model.Project, error) {

	logger := utils.LoggerFromContext(ctx)

	// Get project by key
	project, err := s.projectRepo.GetProjectByKey(key)
	if err != nil {
		logger.Error("project not found", "key", key)
		return nil, errors.New("project not found")
	}

	// Check if user has access to this projects
	role, err := s.memberRepo.GetUserRole(project.Key, userID)
	if err != nil {
		logger.Error("access denied", "project_key", key, "user_id", userID)
		return nil, errors.New("no access to this project")
	}

	if role == "" {
		logger.Error("not a member", "project_key", key, "user_id", userID)
		return nil, errors.New("not a member")
	}

	return project, nil
}

func (s *ProjectService) UpdateProject(project *model.Project, userID int) error {

	role, err := s.memberRepo.GetUserRole(project.Key, userID)
	if err != nil {
		return err
	}

	if role != model.RoleOwner && role != model.RoleAdmin {
		return errors.New("no permission to update project")
	}

	return s.projectRepo.UpdateProject(project)
}

func (s *ProjectService) DeleteProjectByKey(ctx context.Context, key string, userID int) error {

	project, err := s.projectRepo.GetProjectByKey(key)
	if err != nil {
		return err
	}

	if project.OwnerID != userID {
		return errors.New("only owner can delete project")
	}

	return s.projectRepo.DeleteProjectByKey(key)
}

func (s *ProjectService) ListUserProjects(userID int) ([]model.Project, error) {
	// Get all projects where user is owner, admin, or member
	var allProjects []model.Project
	// Get owned projects
	owned, err := s.projectRepo.ListProjectsByOwner(userID)
	if err != nil {
		// If no owned projects, continue
		owned = []model.Project{}
	}

	// Get projects where user is admin
	admin, err := s.projectRepo.ListProjectsWhereUserIsAdmin(userID)
	if err != nil {
		admin = []model.Project{}
	}

	// Get projects where user is member
	member, err := s.projectRepo.ListProjectsWhereUserIsMember(userID)
	if err != nil {
		member = []model.Project{}
	}

	// Combine all projects and remove duplicates
	projectMap := make(map[int]model.Project)
	for _, p := range owned {
		projectMap[p.ID] = p
	}
	for _, p := range admin {
		projectMap[p.ID] = p
	}
	for _, p := range member {
		projectMap[p.ID] = p
	}

	for _, p := range projectMap {
		allProjects = append(allProjects, p)
	}

	if len(allProjects) == 0 {
		return nil, fmt.Errorf("user %d has no projects", userID)
	}

	return allProjects, nil
}

func (s *ProjectService) IsOwner(projectID, userID int) (bool, error) {

	project, err := s.projectRepo.GetProjectByID(projectID)
	if err != nil {
		return false, err
	}

	return project.OwnerID == userID, nil
}
