package model

import (
	"time"
)

type LoginInput struct {
	ID string `json:"id"`
}

type Token struct {
	ID           string    `json:"id"`
	UserId       string    `json:"user"`
	SignedString string    `json:"content"`
	CreatedAt    time.Time `json:"created_at"`
	ExpiresAt    time.Time `json:"expires_at"`
}
