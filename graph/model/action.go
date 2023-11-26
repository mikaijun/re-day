package model

import (
	"time"
)

type NewAction struct {
	TaskId  string `json:"taskId"`
	Score   int    `json:"score"`
	Comment string `json:"comment"`
}

type UpdateAction struct {
	ActionId string `json:"actionId"`
	Score    int    `json:"score"`
	Comment  string `json:"comment"`
}

type Action struct {
	ID        string    `json:"id"`
	TaskId    string    `json:"task"`
	Score     int       `json:"score"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
