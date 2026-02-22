package repo

import (
	"errors"
	"time"
	"todolist/model"
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
	tasklist []*model.Task
}

func NewTaskRepo() TaskRepo {
	repo := &taskRepo{}
	return repo
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
	//adding ID
	task.ID = len(t.tasklist) + 1
	//Adding time
	task.CreatedAt = time.Now()
	err := task.Validate()
	if err != nil {
		return nil, err
	} else {
		t.tasklist = append(t.tasklist, task)
	}

	return task, nil
}

func (t *taskRepo) GetTaskByID(id int) (*model.Task, error) {
	for _, task := range t.tasklist {
		if task.ID == id {
			return task, nil
		}
	}
	return nil, errors.New("Task not found")
}

func (t *taskRepo) UpdateTask(utask *model.Task) error {
	for idx, task := range t.tasklist {
		if task.ID == utask.ID {
			utask.CreatedAt = task.CreatedAt
			t.tasklist[idx] = utask
			return nil
		}
	}

	return errors.New("Task not found")
}

func (t *taskRepo) DeleteTask(id int) error {
	var tempList []*model.Task

	for _, task := range t.tasklist {
		if task.ID != id {
			tempList = append(tempList, task)
		}
	}
	//to maintain ID order 1,2,3,4....
	for i := 0; i < len(tempList); i++ {
		tempList[i].ID = i + 1
	}

	t.tasklist = tempList

	return nil

}

func (t *taskRepo) ListTasks() ([]model.Task, error) {
	var tasks []model.Task

	if len(t.tasklist) == 0 {
		return nil, nil
	} else {
		for _, task := range t.tasklist {
			tasks = append(tasks, *task)
		}
		return tasks, nil
	}

}

func (t *taskRepo) ListTasksByProject(projectID int) ([]model.Task, error) {
	var tempList []model.Task

	for _, task := range t.tasklist {
		if task.ProjectID == projectID {
			tempList = append(tempList, *task)
		}
	}
	if len(tempList) > 0 {
		return tempList, nil
	} else {
		return nil, errors.New("No Task Available under this Project")
	}

}

func (t *taskRepo) ListTasksByAssignee(assigneeID int) ([]model.Task, error) {
	var tempList []model.Task

	for _, task := range t.tasklist {
		if task.AssigneeID != nil && *task.AssigneeID == assigneeID {
			tempList = append(tempList, *task)
		}
	}
	if len(tempList) > 0 {
		return tempList, nil
	} else {
		return nil, errors.New("No Task is assigned by this assignee")
	}

}

// func (t *taskRepo) Completed(id int) error {
// 	for _, task := range t.tasklist {
// 		if task.Id == id {
// 			task.Complete = true
// 			return nil
// 		}
// 	}

// 	return errors.New("Task not found")
// }
