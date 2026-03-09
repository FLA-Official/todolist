package service

import (
	"errors"
	"todolist/model"
	"todolist/repo"
)

type TaskService struct {
	taskRepo    repo.TaskRepo
	projectRepo repo.ProjectRepo
	memberRepo  repo.ProjectMemberRepo
}

func NewTaskService(
	taskRepo repo.TaskRepo,
	projectRepo repo.ProjectRepo,
	memberRepo repo.ProjectMemberRepo,
) *TaskService {

	return &TaskService{
		taskRepo:    taskRepo,
		projectRepo: projectRepo,
		memberRepo:  memberRepo,
	}
}

func (s *TaskService) CreateTask(task *model.Task, userID int) (*model.Task, error) {

	role, err := s.memberRepo.GetUserRole(task.ProjectID, userID)
	if err != nil {
		return nil, err
	}

	if role == "" {
		return nil, errors.New("not a project member")
	}

	return s.taskRepo.CreateTask(task)
}

func (s *TaskService) UpdateTask(task *model.Task, userID int) error {

	role, err := s.memberRepo.GetUserRole(task.ProjectID, userID)
	if err != nil {
		return err
	}

	if role == model.RoleMember {
		return errors.New("members cannot update tasks")
	}

	return s.taskRepo.UpdateTask(task)
}

func (s *TaskService) DeleteTask(taskID, projectID, userID int) error {

	role, err := s.memberRepo.GetUserRole(projectID, userID)
	if err != nil {
		return err
	}

	if role != model.RoleOwner && role != model.RoleAdmin {
		return errors.New("no permission to delete task")
	}

	return s.taskRepo.DeleteTask(taskID)
}

func (s *TaskService) GetProjectTasks(projectID, userID int) ([]model.Task, error) {

	role, err := s.memberRepo.GetUserRole(projectID, userID)
	if err != nil {
		return nil, err
	}

	if role == "" {
		return nil, errors.New("not a project member")
	}

	return s.taskRepo.ListTasksByProject(projectID)
}
