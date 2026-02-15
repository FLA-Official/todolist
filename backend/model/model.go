package model

import (
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
