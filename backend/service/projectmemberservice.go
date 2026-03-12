package service

import (
	"errors"
	"todolist/model"
	"todolist/repo"
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

	role, err := s.memberRepo.GetUserRole(member.ProjectID, userID)
	if err != nil {
		return err
	}

	if role != model.RoleOwner && role != model.RoleAdmin {
		return errors.New("no permission to add member")
	}

	return s.memberRepo.AddMember(member)
}

func (s *ProjectMemberService) RemoveMember(projectID, targetUserID, userID int) error {
	//Authentication of user
	role, err := s.memberRepo.GetUserRole(projectID, userID)
	if err != nil {
		return err
	}

	if role != model.RoleOwner && role != model.RoleAdmin {
		return errors.New("no permission to remove member")
	}

	targetMemberRole, err := s.memberRepo.GetUserRole(projectID, targetUserID)

	//Authorization of who can remove who

	if targetMemberRole == model.RoleOwner {
		return errors.New("cannot remove project owner")
	}

	if targetMemberRole == model.RoleAdmin && role == model.RoleAdmin {
		return errors.New("Only Owner can remove an admin")
	}

	return s.memberRepo.RemoveMember(projectID, targetUserID)
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

	role, err := s.memberRepo.GetUserRole(targetMember.ProjectID, userID)
	if err != nil {
		return err
	}

	if role != model.RoleOwner {
		return errors.New("no permission to update member's role")
	}

	return s.memberRepo.UpdateMemberRole(projectID, targetMember.UserID, newrole)
}
