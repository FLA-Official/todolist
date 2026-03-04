package repo

import (
	"errors"
	"todolist/model"

	"github.com/jmoiron/sqlx"
)

type TaskRepo interface {
	CreateTask(task *model.Task) (*model.Task, error)
	GetTaskByID(taskID int, userID int) (*model.Task, error)
	UpdateTask(task *model.Task, userID int) error
	DeleteTask(taskID int, userID int) error
	ListTasks(userID int) ([]model.Task, error)
	ListTasksByProjectAndUser(projectID int, userID int) ([]model.Task, error)
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

// 	ID          int       `json:"id" gorm:"primaryKey"`
// 	ProjectID   int       `json:"project_id"`
// 	Project     Project   `json:"project" gorm:"foreignKey:ProjectID"`
// 	Title       string    `json:"title"`
// 	Description string    `json:"description,omitempty"`
// 	Status      string    `json:"status"`
// 	Priority    string    `json:"priority"`
// 	AssigneeID  *int      `json:"assignee_id,omitempty"`
// 	CreatedAt   time.Time `json:"created_at"`
// 	EndAt   time.Time `json:"updated_at"`

func (t *taskRepo) CreateTask(task *model.Task) (*model.Task, error) {
	if err := task.Validate(); err != nil {
		return nil, err
	}

	query := `
		INSERT INTO tasks 
		(project_id, title, description, status, priority, assignee_id, end_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id, created_at
	`

	err := t.dbCon.QueryRow(
		query,
		task.ProjectID,
		task.Title,
		task.Description,
		task.Status,
		task.Priority,
		task.AssigneeID,
		task.EndAt,
	).Scan(&task.ID, &task.CreatedAt)

	if err != nil {
		return nil, err
	}

	return task, nil
}

func (r *taskRepo) GetTaskByID(taskID int, userID int) (*model.Task, error) {

	var task model.Task

	query := `
	SELECT t.*
	FROM tasks t
	JOIN projects p ON t.project_id = p.id
	WHERE t.id = $1 AND p.owner_id = $2
	`

	err := r.dbCon.Get(&task, query, taskID, userID)
	if err != nil {
		return nil, nil // not found or not allowed
	}

	return &task, nil
}

func (r *taskRepo) UpdateTask(task *model.Task, userID int) error {

	query := `
	UPDATE tasks t
	SET title = $1,
	    description = $2,
	    status = $3
	FROM projects p
	WHERE t.project_id = p.id
	AND t.id = $4
	AND p.owner_id = $5
	`

	result, err := r.dbCon.Exec(query,
		task.Title,
		task.Description,
		task.Status,
		task.ID,
		userID,
	)

	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return errors.New("not allowed or task not found")
	}

	return nil
}
func (t *taskRepo) DeleteTask(taskID int, userID int) error {

	query := `
	DELETE FROM tasks
	USING projects
	WHERE tasks.project_id = projects.id
	AND tasks.id = $1
	AND projects.owner_id = $2
	`

	result, err := t.dbCon.Exec(query, taskID, userID)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("task not found or unauthorized")
	}

	return nil
}

func (r *taskRepo) ListTasks(userID int) ([]model.Task, error) {

	var tasks []model.Task

	query := `
	SELECT t.*
	FROM tasks t
	JOIN projects p ON t.project_id = p.id
	WHERE p.owner_id = $1
	`

	err := r.dbCon.Select(&tasks, query, userID)
	return tasks, err
}

func (r *taskRepo) ListTasksByProjectAndUser(projectID int, userID int) ([]model.Task, error) {
	var tasks []model.Task

	query := `
	SELECT t.*
	FROM tasks t
	JOIN projects p ON t.project_id = p.id
	WHERE t.project_id = $1 AND p.owner_id = $2
	`

	err := r.dbCon.Select(&tasks, query, projectID, userID)
	if err != nil {
		return nil, err
	}

	if len(tasks) == 0 {
		return nil, errors.New("no tasks found or unauthorized")
	}

	return tasks, nil
}
func (t *taskRepo) ListTasksByAssignee(assigneeID int) ([]model.Task, error) {
	var tasks []model.Task

	err := t.dbCon.Select(&tasks, `SELECT * FROM tasks WHERE assignee_id=$1`, assigneeID)
	if len(tasks) == 0 {
		return nil, errors.New("no tasks found")
	}

	return tasks, err
}
