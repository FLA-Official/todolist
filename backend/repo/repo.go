package repo

import (
	"errors"
	"time"
	"todolist/model"
)

type TaskRepo interface {
	List() []*model.Task
	GetTaskByID(id int) (*model.Task, error)
	StoreTask(task model.Task) (*model.Task, error)
	Completed(id int)
	Delete(id int) error
	Update(utask model.Task) (*model.Task, error)
}

type taskRepo struct {
	tasklist []*model.Task
}

func NewTaskRepo() TaskRepo {
	repo := &taskRepo{}
	return repo
}

func (t *taskRepo) List() []*model.Task {
	return t.tasklist
}

func (t *taskRepo) GetTaskByID(id int) (*model.Task, error) {
	for _, task := range t.tasklist {
		if task.Id == id {
			return task, nil
		} else {
			return nil, errors.New("The task is not found")
		}
	}
	return nil, nil
}

func (t *taskRepo) StoreTask(task model.Task) (*model.Task, error) {
	//adding ID
	task.Id = len(t.tasklist) + 1
	//Adding time
	task.CreatedTime = time.Now()
	//task by default false at start
	t.tasklist[len(t.tasklist)].Complete = false

	t.tasklist = append(t.tasklist, &task)

	return &task, nil
}

func (t *taskRepo) Update(utask model.Task) (*model.Task, error) {
	for idx, task := range t.tasklist {
		if task.Id == utask.Id {
			t.tasklist[idx] = &utask
		}
	}

	return &utask, nil
}

func (t *taskRepo) Delete(taskID int) error {
	var tempList []*model.Task

	for _, t := range t.tasklist {
		if t.Id != taskID {
			tempList = append(tempList, t)
		}
	}

	t.tasklist = tempList

	return nil

}

func (t *taskRepo) Completed(id int) {
	t.tasklist[id].Complete = true
}
