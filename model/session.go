package model

import (
	"time"
)

type Session struct {
	ID        uint64    `json:"id"`
	UserID    uint64    `json:"user_id"`
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	ExpiresAt time.Time `json:"expires_at"`
	Status    string    `json:"status"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
