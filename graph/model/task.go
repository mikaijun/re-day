package model

import (
	"time"
)

type NewTask struct {
	Content string `json:"content"`
}

type UpdateTask struct {
	TaskId  string `json:"taskId"`
	Content string `json:"content"`
}

type Task struct {
	ID        string    `json:"id"`
	Content   string    `json:"content"`
	UserId    string    `json:"user"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
