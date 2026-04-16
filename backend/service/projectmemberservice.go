package service

import (
	"context"
	"errors"
	"todolist/model"
	"todolist/repo"
	"todolist/utils"
)

type ProjectMemberService struct {
	memberRepo  repo.ProjectMemberRepo
	projectRepo repo.ProjectRepo
}

func NewProjectMemberService(
	memberRepo repo.ProjectMemberRepo,
	projectRepo repo.ProjectRepo,
) *ProjectMemberService {

	return &ProjectMemberService{
		memberRepo:  memberRepo,
		projectRepo: projectRepo,
	}
}

func (s *ProjectMemberService) AddMember(member *model.ProjectMember, userID int) error {

	project, err := s.projectRepo.GetProjectByID(member.ProjectID)
	if err != nil {
		return err
	}

	role, err := s.memberRepo.GetUserRole(project.Key, userID)
	if err != nil {
		return err
	}

	if role != model.RoleOwner && role != model.RoleAdmin {
		return errors.New("no permission to add member")
	}

	return s.memberRepo.AddMember(member)
}

func (s *ProjectMemberService) RemoveMember(ctx context.Context, projectKey string, targetUserID, userID int) error {

	logger := utils.LoggerFromContext(ctx)

	role, err := s.memberRepo.GetUserRole(projectKey, userID)
	if err != nil {
		logger.Error("failed to get role", "project_key", projectKey, "user_id", userID)
		return err
	}

	if role != model.RoleOwner && role != model.RoleAdmin {
		logger.Error("unauthorized remove attempt",
			"project_key", projectKey,
			"user_id", userID,
			"target_user_id", targetUserID,
		)
		return errors.New("no permission to remove member")
	}

	project, err := s.projectRepo.GetProjectByKey(projectKey)
	if err != nil {
		logger.Error("failed to resolve project key", "project_key", projectKey)
		return err
	}

	targetMemberRole, err := s.memberRepo.GetUserRole(projectKey, targetUserID)
	if err != nil {
		logger.Error("failed to get target role", "target_user_id", targetUserID)
		return err
	}

	if targetMemberRole == model.RoleOwner {
		logger.Error("attempt to remove owner", "target_user_id", targetUserID)
		return errors.New("cannot remove project owner")
	}

	if targetMemberRole == model.RoleAdmin && role == model.RoleAdmin {
		logger.Error("admin tried removing admin", "user_id", userID)
		return errors.New("Only Owner can remove an admin")
	}

	logger.Info("member removed",
		"project_key", projectKey,
		"removed_user", targetUserID,
		"removed_by", userID,
	)

	return s.memberRepo.RemoveMember(project.ID, targetUserID)
}

func (s *ProjectMemberService) GetProjectMembers(projectID int) ([]model.ProjectMember, error) {

	return s.memberRepo.GetMembersByProject(projectID)
}

func (s *ProjectMemberService) GetProjectMemberbyID(projectID int, userID int) (*model.ProjectMember, error) {

	return s.memberRepo.GetMember(projectID, userID)
}

func (s *ProjectMemberService) UpdateMemberRole(targetMemberID, projectID, userID int, newrole string) error {
	targetMember, err := s.memberRepo.GetMember(projectID, targetMemberID)
	if err != nil {
		return err
	}

	project, err := s.projectRepo.GetProjectByID(targetMember.ProjectID)
	if err != nil {
		return err
	}

	role, err := s.memberRepo.GetUserRole(project.Key, userID)
	if err != nil {
		return err
	}

	if role != model.RoleOwner {
		return errors.New("no permission to update member's role")
	}

	return s.memberRepo.UpdateMemberRole(projectID, targetMember.UserID, newrole)
}
