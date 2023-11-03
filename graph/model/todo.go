package model

import (
	"time"
)

type NewTodo struct {
	Text   string `json:"text"`
	UserID string `json:"userId"`
}

type Todo struct {
	ID        string    `json:"id"`
	Text      string    `json:"text"`
	Done      bool      `json:"done"`
	UserId    string    `json:"user"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
