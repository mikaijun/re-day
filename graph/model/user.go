package model

import "time"

type ctxKey string

const (
	AuthKey = ctxKey("auth")
)

type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type NewUser struct {
	Name string `json:"name"`
}
