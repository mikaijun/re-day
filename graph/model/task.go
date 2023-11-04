package model

import (
	"time"
)

type NewTask struct {
	Text   string `json:"text"`
	UserID string `json:"userId"`
}

type Task struct {
	ID        string    `json:"id"`
	Text      string    `json:"text"`
	Done      bool      `json:"done"`
	UserId    string    `json:"user"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
