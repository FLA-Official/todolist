package model

import (
	"errors"
	"time"
)

type TaskBox struct {
	Id          int
	Title       string
	Description string
	DueDate     string
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

func AddTask(title string, description string, duedate string) error {
	id := len(task) + 1
	task[len(task)].Id = id
	if title == "" {
		return errors.New("Please provdie valid title")
	} else {
		task[len(task)].Title = title
	}

	if description == "" {
		task[len(task)].Description = ""
	} else {
		task[len(task)].Description = description
	}

	currentTime := time.Now()

	if duedate == "" {
		return errors.New("Please provdie a valid time to complete the task")
	} else if duedate == currentTime {
		return errors.New("Please provide a ")
	}

	task[len(task)].DueDate = duedate
	task[len(task)].Complete = false

	return nil
}
