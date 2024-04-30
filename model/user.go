package model

import "time"

type User struct {
	ID        uint64    `json:"id"`
	Type      string    `json:"type"`
	Username  string    `json:"username"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  *string   `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
