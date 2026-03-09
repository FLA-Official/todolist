package repo

import (
	"errors"
	"todolist/model"

	"github.com/jmoiron/sqlx"
)

type TaskRepo interface {
	CreateTask(task *model.Task) (*model.Task, error)
	GetTaskByID(taskID int) (*model.Task, error)
	UpdateTask(task *model.Task) error
	DeleteTask(taskID int) error
	ListTasks() ([]model.Task, error)
	ListTasksByProject(projectID int) ([]model.Task, error)
	ListTasksByAssignee(assigneeID int) ([]model.Task, error)
}

type taskRepo struct {
	dbCon *sqlx.DB
}

func NewTaskRepo(dbCon *sqlx.DB) TaskRepo {
	return &taskRepo{
		dbCon: dbCon,
	}
}

func (r *taskRepo) CreateTask(task *model.Task) (*model.Task, error) {

	query := `
	INSERT INTO tasks (project_id, title, description, status, priority, assignee_id)
	VALUES ($1,$2,$3,$4,$5,$6)
	RETURNING id, created_at
	`

	err := r.dbCon.QueryRow(
		query,
		task.ProjectID,
		task.Title,
		task.Description,
		task.Status,
		task.Priority,
		task.AssigneeID,
	).Scan(&task.ID, &task.CreatedAt)

	if err != nil {
		return nil, err
	}

	return task, nil
}

func (r *taskRepo) GetTaskByID(taskID int) (*model.Task, error) {

	var task model.Task

	query := `SELECT * FROM tasks WHERE id = $1`

	err := r.dbCon.Get(&task, query, taskID)
	if err != nil {
		return nil, nil
	}

	return &task, nil
}

func (r *taskRepo) UpdateTask(task *model.Task) error {

	query := `
	UPDATE tasks
	SET title = $1,
	    description = $2,
	    status = $3,
	    priority = $4,
	    assignee_id = $5
	WHERE id = $6
	`

	result, err := r.dbCon.Exec(
		query,
		task.Title,
		task.Description,
		task.Status,
		task.Priority,
		task.AssigneeID,
		task.ID,
	)

	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return errors.New("task not found")
	}

	return nil
}

func (r *taskRepo) DeleteTask(taskID int) error {

	query := `DELETE FROM tasks WHERE id = $1`

	result, err := r.dbCon.Exec(query, taskID)
	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return errors.New("task not found")
	}

	return nil
}

func (r *taskRepo) ListTasks() ([]model.Task, error) {

	var tasks []model.Task

	query := `SELECT * FROM tasks ORDER BY created_at DESC`

	err := r.dbCon.Select(&tasks, query)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *taskRepo) ListTasksByProject(projectID int) ([]model.Task, error) {

	var tasks []model.Task

	query := `
	SELECT *
	FROM tasks
	WHERE project_id = $1
	ORDER BY created_at DESC
	`

	err := r.dbCon.Select(&tasks, query, projectID)
	if err != nil {
		return nil, err
	}

	if len(tasks) == 0 {
		return nil, errors.New("no tasks found")
	}

	return tasks, nil
}

func (r *taskRepo) ListTasksByAssignee(assigneeID int) ([]model.Task, error) {

	var tasks []model.Task

	query := `
	SELECT *
	FROM tasks
	WHERE assignee_id = $1
	ORDER BY created_at DESC
	`

	err := r.dbCon.Select(&tasks, query, assigneeID)
	if err != nil {
		return nil, err
	}

	if len(tasks) == 0 {
		return nil, errors.New("no tasks found")
	}

	return tasks, nil
}
