package repo

import (
	"errors"
	"todolist/model"

	"github.com/jmoiron/sqlx"
)

type TaskRepo interface {
	CreateTask(task *model.Task) (*model.Task, error)
	GetTaskByID(id int) (*model.Task, error)
	UpdateTask(task *model.Task) error
	DeleteTask(id int) error
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

func (t *taskRepo) GetTaskByID(id int) (*model.Task, error) {
	var task model.Task

	query := `SELECT * FROM tasks WHERE id = $1`

	err := t.dbCon.Get(&task, query, id)
	if err != nil {
		return nil, errors.New("task not found")
	}

	return &task, nil
}

func (t *taskRepo) UpdateTask(task *model.Task) error {
	var existing model.Task

	err := t.dbCon.Get(&existing, `SELECT * FROM tasks WHERE id=$1`, task.ID)
	if err != nil {
		return errors.New("task not found")
	}

	// preserve created_at
	task.CreatedAt = existing.CreatedAt

	if err := task.Validate(); err != nil {
		return err
	}

	query := `
		UPDATE tasks SET
			project_id=$1,
			title=$2,
			description=$3,
			status=$4,
			priority=$5,
			assignee_id=$6,
			end_at=$7
		WHERE id=$8
	`

	_, err = t.dbCon.Exec(
		query,
		task.ProjectID,
		task.Title,
		task.Description,
		task.Status,
		task.Priority,
		task.AssigneeID,
		task.EndAt,
		task.ID,
	)

	return err
}
func (t *taskRepo) DeleteTask(id int) error {
	result, err := t.dbCon.Exec(`DELETE FROM tasks WHERE id=$1`, id)
	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return errors.New("task not found")
	}

	return nil
}

func (t *taskRepo) ListTasks() ([]model.Task, error) {
	var tasks []model.Task

	err := t.dbCon.Select(&tasks, `SELECT * FROM tasks`)
	return tasks, err
}

func (t *taskRepo) ListTasksByProject(projectID int) ([]model.Task, error) {
	var tasks []model.Task

	err := t.dbCon.Select(&tasks, `SELECT * FROM tasks WHERE project_id=$1`, projectID)
	if len(tasks) == 0 {
		return nil, errors.New("no tasks found")
	}

	return tasks, err
}
func (t *taskRepo) ListTasksByAssignee(assigneeID int) ([]model.Task, error) {
	var tasks []model.Task

	err := t.dbCon.Select(&tasks, `SELECT * FROM tasks WHERE assignee_id=$1`, assigneeID)
	if len(tasks) == 0 {
		return nil, errors.New("no tasks found")
	}

	return tasks, err
}
