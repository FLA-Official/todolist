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
	Completed(id int) error
	Delete(id int) error
	Update(utask model.Task) (*model.Task, error)
}

type taskRepo struct {
	tasklist []*model.Task
}

func NewTaskRepo() TaskRepo {
	repo := &taskRepo{
		tasklist: []*model.Task{},
	}

	generateinittask(repo)
	return repo
}

func (t *taskRepo) List() []*model.Task {
	return t.tasklist
}

func (t *taskRepo) GetTaskByID(id int) (*model.Task, error) {
	for _, task := range t.tasklist {
		if task.Id == id {
			return task, nil
		}
	}
	return nil, errors.New("Task not found")
}

func (t *taskRepo) StoreTask(task model.Task) (*model.Task, error) {
	//adding ID
	task.Id = len(t.tasklist) + 1
	//Adding time
	task.CreatedTime = time.Now()
	//task by default false at start
	task.Complete = false

	t.tasklist = append(t.tasklist, &task)

	return &task, nil
}

func (t *taskRepo) Update(utask model.Task) (*model.Task, error) {
	for idx, task := range t.tasklist {
		if task.Id == utask.Id {
			utask.CreatedTime = task.CreatedTime
			t.tasklist[idx] = &utask
			return &utask, nil
		}
	}

	return nil, errors.New("Task not found")
}

func (t *taskRepo) Delete(taskID int) error {
	var tempList []*model.Task

	for _, task := range t.tasklist {
		if task.Id != taskID {
			tempList = append(tempList, task)
		}
	}

	t.tasklist = tempList

	return nil

}

func (t *taskRepo) Completed(id int) error {
	for _, task := range t.tasklist {
		if task.Id == id {
			task.Complete = true
			return nil
		}
	}

	return errors.New("Task not found")
}

func generateinittask(r *taskRepo) {
	t1 := &model.Task{
		Id:          1,
		Title:       "abc",
		Description: "bcd",
		CreatedTime: time.Now(),
		EndDate:     time.Date(2026, time.February, 18, 6, 00, 00, 00, time.UTC),
		Complete:    false,
	}

	t2 := &model.Task{
		Id:          2,
		Title:       "abcd",
		Description: "bcdg",
		CreatedTime: time.Now(),
		EndDate:     time.Date(2026, time.February, 18, 6, 00, 00, 00, time.UTC),
		Complete:    false,
	}

	r.tasklist = append(r.tasklist, t1)
	r.tasklist = append(r.tasklist, t2)

}
