package service

import (
	"context"
	"errors"
	"todolist/model"
	"todolist/repo"
	"todolist/utils"
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

func (s *TaskService) CreateTask(ctx context.Context, task *model.Task, userID int) (*model.Task, error) {

	logger := utils.LoggerFromContext(ctx)

	role, err := s.memberRepo.GetUserRole(task.ProjectKey, userID)
	if err != nil {
		logger.Error("role fetch failed", "project_Key", task.ProjectKey, "user_id", userID)
		return nil, err
	}

	if role == "" {
		logger.Error("not a project member", "project_Key", task.ProjectKey, "user_id", userID)
		return nil, errors.New("not a project member")
	}

	logger.Info("task created", "project_id", task.ProjectKey, "user_id", userID)

	return s.taskRepo.CreateTask(task)
}

func (s *TaskService) UpdateTask(task *model.Task, userID int) error {

	role, err := s.memberRepo.GetUserRole(task.ProjectKey, userID)
	if err != nil {
		return err
	}

	if role == model.RoleMember {
		return errors.New("members cannot update tasks")
	}

	return s.taskRepo.UpdateTask(task)
}

func (s *TaskService) DeleteTask(projectKey string, taskID, userID int) error {

	role, err := s.memberRepo.GetUserRole(projectKey, userID)
	if err != nil {
		return err
	}

	if role != model.RoleOwner && role != model.RoleAdmin {
		return errors.New("no permission to delete task")
	}

	return s.taskRepo.DeleteTask(taskID)
}

func (s *TaskService) GetProjectTasks(projectID, userID int) ([]model.Task, error) {

	project, err := s.projectRepo.GetProjectByID(projectID)
	if err != nil {
		return nil, err
	}

	role, err := s.memberRepo.GetUserRole(project.Key, userID)
	if err != nil {
		return nil, err
	}

	if role == "" {
		return nil, errors.New("not a project member")
	}

	return s.taskRepo.ListTasksByProject(projectID)
}
