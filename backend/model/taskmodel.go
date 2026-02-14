package model

import (
	"errors"
	"time"
)

type Task struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedTime time.Time `json:"createdtime"`
	EndDate     time.Time `json:"endtime"`
	Complete    bool      `json:"complete"`
}

type TaskRepo interface {
	List() []*Task
	GetTaskByID(id int) (*Task, error)
	CreateTask(title string, description string, endDate time.Time) error
	Completed()
	Delete(id int) error
}

type taskRepo struct {
	tasklist []*Task
}

func NewTaskRepo() TaskRepo {
	repo := &taskRepo{}
	return repo
}

func (t *taskRepo) List() []*Task {
	return t.tasklist
}

func (t *taskRepo) GetTaskByID(id int) (*Task, error) {
	for _, task := range t.tasklist {
		if task.Id == id {
			return task, nil
		}
	}
	return nil, nil
}

func (t *taskRepo) CreateTask(title string, description string, endDate time.Time) error {
	//adding ID
	id := len(t.tasklist) + 1
	t.tasklist[len(t.tasklist)].Id = id

	//Adding Title is mandetory
	if title == "" {
		return errors.New("Please provdie valid title")
	} else {
		t.tasklist[len(t.tasklist)].Title = title
	}

	//Adding description is optional
	if description == "" {
		t.tasklist[len(t.tasklist)].Description = ""
	} else {
		t.tasklist[len(t.tasklist)].Description = description
	}

	//Adding time and validation
	currentTime := time.Now()
	t.tasklist[len(t.tasklist)].CreatedTime = currentTime

	if endDate.IsZero() {
		return errors.New("Please provdie a valid time to complete the task")
	} else if endDate.Equal(currentTime) {
		return errors.New("Please provide a valid time to complete the task")
	} else if endDate.Before(currentTime) {
		return errors.New("due date cannot be in the past")
	} else {
		t.tasklist[len(t.tasklist)].EndDate = endDate
	}

	//task by default false at start
	t.tasklist[len(t.tasklist)].Complete = false

	return nil
}

func (t *taskRepo) Update(utask Task) (*Task, error) {
	for idx, task := range t.tasklist {
		if task.Id == utask.Id {
			t.tasklist[idx] = &utask
		}
	}

	return &utask, nil
}

func (t *taskRepo) Delete(taskID int) error {
	var tempList []*Task

	for _, t := range t.tasklist {
		if t.Id != taskID {
			tempList = append(tempList, t)
		}
	}

	t.tasklist = tempList

	return nil

}

func (t *taskRepo) Completed() {

}
