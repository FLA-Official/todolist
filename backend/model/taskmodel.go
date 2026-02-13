package model

import (
	"errors"
	"time"
)

type TaskBox struct {
	Id          int
	Title       string
	Description string
	CreatedTime time.Time
	EndDate     time.Time
	Complete    bool
}

var task []*TaskBox

func Get() []*TaskBox {
	return task
}

func GetTaskByID(id int) *TaskBox {
	for i := range task {
		if task[i].Id == id {
			return task[i]
		}
	}
	return nil
}

func AddTask(title string, description string, endDate time.Time) error {
	//adding ID
	id := len(task) + 1
	task[len(task)].Id = id

	//Adding Title is mandetory
	if title == "" {
		return errors.New("Please provdie valid title")
	} else {
		task[len(task)].Title = title
	}

	//Adding description is optional
	if description == "" {
		task[len(task)].Description = ""
	} else {
		task[len(task)].Description = description
	}

	//Adding time and validation
	currentTime := time.Now()
	task[len(task)].CreatedTime = currentTime

	if endDate.IsZero() {
		return errors.New("Please provdie a valid time to complete the task")
	} else if endDate.Equal(currentTime) {
		return errors.New("Please provide a valid time to complete the task")
	} else if endDate.Before(currentTime) {
		return errors.New("due date cannot be in the past")
	} else {
		task[len(task)].EndDate = endDate
	}

	//task by default false at start
	task[len(task)].Complete = false

	return nil
}

func (t *TaskBox) Completed() {
	t.Complete = true
}
